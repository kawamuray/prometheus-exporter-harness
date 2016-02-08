prometheus-exporter-harness
===========================

A [prometheus](https://prometheus.io/) exporter framework - Make it bit easy to build own prometheus exporter.

Usage
=====
Please see `example/` directory for more detail.

Coding
------
```go
package main

import (
	"github.com/codegangsta/cli"
	"github.com/kawamuray/prometheus-exporter-harness/harness"
	"github.com/prometheus/client_golang/prometheus"
)

type collector struct{}

func (col *collector) Collect(reg *harness.MetricRegistry) {
	reg.Get("example_metric_A").(prometheus.Gauge).Set(128.0)
}

func exampleInit(c *cli.Context, reg *harness.MetricRegistry) (harness.Collector, error) {
	reg.Register("example_metric_A", prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "example_metric_A",
		Help: "is example_metric_A",
	}))
	return &collector{}, nil
}

func main() {
	opts := harness.NewExporterOpts("example_exporter", "0.0.1")
	opts.Init = exampleInit
	harness.Main(opts)
}
```

Building
--------
```sh
./gow get ./example
./gow build -o example_exporter ./example
```

Running
-------
```sh
./example_exporter --interval 10
```

See Also
========
- [prometheus-json-exporter](https://github.com/kawamuray/prometheus-json-exporter)
