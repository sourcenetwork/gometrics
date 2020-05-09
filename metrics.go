package metrics

import (
	km "github.com/go-kit/kit/metrics"
)

// Opts provide struct based options for creating metrics
type Opts struct {
	// Namespace the metrics exists under (IE: Service or App name)
	Namespace string
	// Subsystem namespace is the scoped module name responsible for the metrics
	Subsystem string
	// Name of the metric
	Name string
	// Description of the metric, can be provided as help text for certain backends like Prometheus
	Description string
	// SampleRate is the rate at which the metric is update (Good default is 1)
	SampleRate float64
	// Tags to assign to the metric
	Tags []string
}

// CounterOpts opts for Counter
type CounterOpts Opts

// GaugeOpts opgs for Gauges
type GaugeOpts Opts

// HistogramOpts opts for Histograms
type HistogramOpts Opts

// Client is the metric provider of a given service
type Client interface {
	WithSubystem(subsystem string) Client
	NewCounter(name string, sampleRate float64) km.Counter
	NewCounterFrom(opts CounterOpts) km.Counter
	NewGauge(name string) km.Gauge
	NewGaugeFrom(opts GaugeOpts) km.Gauge
	NewHistogram(name string, sampleRate float64) km.Histogram
	NewHistogramFrom(opts HistogramOpts) km.Histogram
}
