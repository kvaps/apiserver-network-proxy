/*
Copyright 2020 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package metrics

import (
	"time"

	"github.com/prometheus/client_golang/prometheus"
)

type Direction string

const (
	namespace = "konnectivity_network_proxy"
	subsystem = "server"
)

var (
	// Use buckets ranging from 5 ms to 12.5 seconds.
	latencyBuckets = []float64{0.005, 0.025, 0.1, 0.5, 2.5, 12.5}

	// Metrics provides access to all dial metrics.
	Metrics = newServerMetrics()
)

// ServerMetrics includes all the metrics of the proxy server.
type ServerMetrics struct {
	latencies *prometheus.HistogramVec
}

// newServerMetrics create a new ServerMetrics, configured with default metric names.
func newServerMetrics() *ServerMetrics {
	latencies := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Namespace: namespace,
			Subsystem: subsystem,
			Name:      "dial_duration_seconds",
			Help:      "Latency of dial to the remote endpoint in seconds",
			Buckets:   latencyBuckets,
		},
		[]string{},
	)
	prometheus.MustRegister(latencies)
	return &ServerMetrics{latencies: latencies}
}

// Reset resets the metrics.
func (a *ServerMetrics) Reset() {
	a.latencies.Reset()
}

// ObserveDialLatency records the latency of dial to the remote endpoint.
func (a *ServerMetrics) ObserveDialLatency(elapsed time.Duration) {
	a.latencies.WithLabelValues().Observe(elapsed.Seconds())
}
