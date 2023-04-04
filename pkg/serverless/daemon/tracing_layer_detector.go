// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

package daemon

import (
	"net/http"
	"sync"
)

// tracingLayerDetector TODO
type tracingLayerDetector struct {
	sync.Mutex
	tracingLayerDetected bool
}

// markLayerDetected TODO
func (t *tracingLayerDetector) markLayerDetected() {
	t.Lock()
	defer t.Unlock()
	t.tracingLayerDetected = true
}

// detected TODO
func (t *tracingLayerDetector) detected() bool {
	t.Lock()
	defer t.Unlock()
	return t.tracingLayerDetected
}

// markLayerWrapper TODO
func (t *tracingLayerDetector) markLayerWrapper(h http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t.markLayerDetected()
		h.ServeHTTP(w, r)
	}
}
