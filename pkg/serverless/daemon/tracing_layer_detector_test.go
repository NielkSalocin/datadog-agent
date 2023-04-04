// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

package daemon

import (
	"net/http"
	"testing"
)

func TestTracingLayerDetector(t *testing.T) {
	var detector tracingLayerDetector
	var funcCalled bool
	handlerFunc := http.HandlerFunc(func(http.ResponseWriter, *http.Request) {
		funcCalled = true
	})
	wrappedFunc := detector.markLayerWrapper(handlerFunc)
	wrappedFunc.ServeHTTP(nil, nil)
	if !funcCalled {
		t.Fatal("handle func was never called")
	}
	if !detector.detected() {
		t.Error("tracing layer should have been marked detected")
	}
}
