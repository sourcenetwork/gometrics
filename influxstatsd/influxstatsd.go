package influxstatsd

import (
	"strings"

	"github.com/go-kit/kit/log"
	km "github.com/go-kit/kit/metrics"
	kstatsd "github.com/go-kit/kit/metrics/influxstatsd"
	metrics "github.com/sourcenetwork/gometrics"
)

type InfluxStatsd struct {
	provider  *kstatsd.Influxstatsd
	subsystem string
}

func New(namespace string, logger log.Logger, labelValues ...string) *InfluxStatsd {
	return &InfluxStatsd{
		provider: kstatsd.New(namespace, logger, labelValues...),
	}
}

// applies any subsystem prefixs to the name
func (d *InfluxStatsd) name(names ...string) string {
	return strings.Join(filterEmpty(names), ".")
}

func (d *InfluxStatsd) WithSub(subsystem string) metrics.Client {
	subsystem = d.name(d.subsystem, subsystem)
	return &InfluxStatsd{
		provider:  d.provider,
		subsystem: subsystem,
	}
}

func (d *InfluxStatsd) NewCounter(name string, sampleRate float64) km.Counter {
	name = d.name(name)
	return d.provider.NewCounter(name, sampleRate)
}

func (d *InfluxStatsd) NewCounterFrom(opts metrics.CounterOpts) km.Counter {
	subsystem := d.subsystem
	if subsystem == "" {
		subsystem = opts.Subsystem
	}
	name := d.name(subsystem, opts.Name)
	var ctr km.Counter = d.provider.NewCounter(name, opts.SampleRate)
	if len(opts.Tags) > 0 {
		ctr = ctr.With(opts.Tags...)
	}
	return ctr
}

func (d *InfluxStatsd) NewGauge(name string) km.Gauge {
	name = d.name(name)
	return d.provider.NewGauge(name)
}

func (d *InfluxStatsd) NewGaugeFrom(opts metrics.GaugeOpts) km.Gauge {
	subsystem := d.subsystem
	if subsystem == "" {
		subsystem = opts.Subsystem
	}
	name := d.name(subsystem, opts.Name)
	var g km.Gauge = d.provider.NewGauge(name)
	if len(opts.Tags) > 0 {
		g = g.With(opts.Tags...)
	}
	return g
}

func (d *InfluxStatsd) NewHistogram(name string, sampleRate float64) km.Histogram {
	name = d.name(name)
	return d.provider.NewHistogram(name, sampleRate)
}

func (d *InfluxStatsd) NewHistogramFrom(opts metrics.HistogramOpts) km.Histogram {
	subsystem := d.subsystem
	if subsystem == "" {
		subsystem = opts.Subsystem
	}
	name := d.name(subsystem, opts.Name)
	var h km.Histogram = d.provider.NewHistogram(name, opts.SampleRate)
	if len(opts.Tags) > 0 {
		h = h.With(opts.Tags...)
	}
	return h
}

func filterEmpty(s []string) []string {
	var notEmpty []string
	for _, str := range s {
		if str != "" {
			notEmpty = append(notEmpty, str)
		}
	}
	return notEmpty
}
