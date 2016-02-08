package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/kawamuray/prometheus-exporter-harness/harness"
	"github.com/prometheus/client_golang/prometheus"
)

type collector struct{}

func (col *collector) Collect(reg *harness.MetricRegistry) {
	reg.Get("example_metric_A").(prometheus.Gauge).Set(128.0)
}

func exampleInit(c *cli.Context, reg *harness.MetricRegistry) (harness.Collector, error) {
	fooEnabled := c.Bool("flagfoo")
	if fooEnabled {
		fmt.Println("foo enabled")
	}

	reg.Register("example_metric_A", prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "example_metric_A",
		Help: "is example_metric_A",
	}))
	reg.Register("example_metric_B", prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "example_metric_B",
		Help: "is example_metric_B",
	}))

	return &collector{}, nil
}

func main() {
	opts := harness.NewExporterOpts("example_exporter", "0.0.1")
	opts.Init = exampleInit
	opts.Flags = []cli.Flag{ // additional flags if necessary
		cli.BoolFlag{
			Name:  "flagfoo",
			Usage: "indicates foo",
		},
	}
	opts.MetricsPath = "/metrics" // default
	opts.Tick = true              // default
	opts.ResetOnTick = true       // default
	harness.Main(opts)
}
