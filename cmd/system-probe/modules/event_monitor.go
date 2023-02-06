// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.
//go:build linux
// +build linux

package modules

import (
	"os"

	"github.com/DataDog/datadog-go/v5/statsd"

	"github.com/DataDog/datadog-agent/cmd/system-probe/api/module"
	"github.com/DataDog/datadog-agent/cmd/system-probe/config"
	"github.com/DataDog/datadog-agent/cmd/system-probe/event_monitor"
	"github.com/DataDog/datadog-agent/pkg/ebpf"
	secconfig "github.com/DataDog/datadog-agent/pkg/security/config"
	secmodule "github.com/DataDog/datadog-agent/pkg/security/module"
	"github.com/DataDog/datadog-agent/pkg/security/probe"
	"github.com/DataDog/datadog-agent/pkg/util/log"
)

const (
	statsdPoolSize = 64
)

func getStatsdClient(seccfg *secconfig.Config) (statsd.ClientInterface, error) {
	statsdAddr := os.Getenv("STATSD_URL")
	if statsdAddr == "" {
		statsdAddr = seccfg.StatsdAddr
	}

	return statsd.New(statsdAddr, statsd.WithBufferPoolSize(statsdPoolSize))
}

// EventMonitor - Event monitor Factory
var EventMonitor = module.Factory{
	Name:             config.EventMonitorModule,
	ConfigNamespaces: []string{"runtime_security_config"},
	Fn: func(sysProbeConfig *config.Config) (module.Module, error) {
		seccfg, err := secconfig.NewConfig(sysProbeConfig)
		if err != nil {
			log.Info("Event monitoring configuration error")
			return nil, module.ErrNotEnabled
		}

		statsdClient, err := getStatsdClient(seccfg)
		if err != nil {
			log.Info("Unable to init statsd client")
			return nil, module.ErrNotEnabled
		}

		evm, err := event_monitor.NewEventMonitor(sysProbeConfig, statsdClient, probe.Opts{})
		if err == ebpf.ErrNotImplemented {
			log.Info("Datadog event monitoring is only supported on Linux")
			return nil, module.ErrNotEnabled
		}

		cws, err := secmodule.NewCWSModule(evm)
		if err != nil {
			return nil, err
		}
		evm.RegisterEventModule(cws)

		/*	process, err := secmodule.NewProcessModule(evm)
			if err != nil {
				return nil, err
			}
			evm.RegisterEventModule(process)*/

		return evm, err
	},
}
