package harness

import (
	"github.com/codegangsta/cli"
)

func MakeApp(opts *ExporterOpts) *cli.App {
	exp := &exporter{opts}

	app := cli.NewApp()
	app.Name = opts.Name
	app.Version = opts.Version
	app.Usage = "A prometheus " + opts.Name
	app.UsageText = opts.Usage
	app.Action = exp.main
	app.Flags = append(opts.Flags,
		cli.IntFlag{
			Name:  "port",
			Usage: "The port number used to expose metrics via http",
			Value: 7979,
		},
		cli.StringFlag{
			Name:  "log-level",
			Usage: "Set Logging level",
			Value: "info",
		},
	)
	if opts.Tick {
		app.Flags = append(app.Flags,
			cli.IntFlag{
				Name:  "interval",
				Usage: "Interval to fetch metrics from the endpoint in second",
				Value: 60,
			},
		)
	}

	return app
}

func Main(opts *ExporterOpts) {
	MakeApp(opts).RunAndExitOnError()
}
