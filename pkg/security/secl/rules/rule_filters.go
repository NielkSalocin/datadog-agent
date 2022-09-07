// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

package rules

import (
	"fmt"

	"github.com/DataDog/datadog-agent/pkg/security/secl/validators"
	"github.com/Masterminds/semver"
)

// RuleFilter definition of a rule filter
type RuleFilter interface {
	IsRuleAccepted(*RuleDefinition) (bool, error)
}

// MacroFilter definition of a macro filter
type MacroFilter interface {
	IsMacroAccepted(*MacroDefinition) (bool, error)
}

// RuleIDFilter defines a ID based filter
type RuleIDFilter struct {
	ID string
}

// IsRuleAccepted checks whether the rule is accepted
func (r *RuleIDFilter) IsRuleAccepted(rule *RuleDefinition) (bool, error) {
	return r.ID == rule.ID, nil
}

// AgentVersionFilter defines a agent version filter
type AgentVersionFilter struct {
	version *semver.Version
}

// NewAgentVersionFilter returns a new agent version based rule filter
func NewAgentVersionFilter(version *semver.Version) (*AgentVersionFilter, error) {
	withoutPreAgentVersion, err := version.SetPrerelease("")
	if err != nil {
		return nil, err
	}

	cleanAgentVersion, err := withoutPreAgentVersion.SetMetadata("")
	if err != nil {
		return nil, err
	}

	return &AgentVersionFilter{
		version: &cleanAgentVersion,
	}, nil
}

// IsRuleAccepted checks whether the rule is accepted
func (r *AgentVersionFilter) IsRuleAccepted(rule *RuleDefinition) (bool, error) {
	constraint, err := validators.ValidateAgentVersionConstraint(rule.AgentVersionConstraint)
	if err != nil {
		return false, fmt.Errorf("failed to parse agent version constraint: %v", err)
	}

	return constraint.Check(r.version), nil
}

// IsMacroAccepted checks whether the macro is accepted
func (r *AgentVersionFilter) IsMacroAccepted(macro *MacroDefinition) (bool, error) {
	constraint, err := validators.ValidateAgentVersionConstraint(macro.AgentVersionConstraint)
	if err != nil {
		return false, fmt.Errorf("failed to parse agent version constraint: %v", err)
	}

	return constraint.Check(r.version), nil
}
