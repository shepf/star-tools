package main

import (
	"context"
	"github.com/shepf/star-tools/lib/lotuslog"
	"github.com/shepf/star-tools/lib/tracing"
	"os"

	"github.com/urfave/cli/v2"
	"go.opencensus.io/trace"

	"github.com/shepf/star-tools/build"
	lcli "github.com/shepf/star-tools/cli"
)

func main() {
	lotuslog.SetupLogLevels()

	local := []*cli.Command{
		DaemonCmd,
	}

	jaeger := tracing.SetupJaegerTracing("lotus")
	defer func() {
		if jaeger != nil {
			jaeger.Flush()
		}
	}()

	for _, cmd := range local {
		cmd := cmd
		originBefore := cmd.Before
		cmd.Before = func(cctx *cli.Context) error {
			trace.UnregisterExporter(jaeger)
			jaeger = tracing.SetupJaegerTracing("star/" + cmd.Name)

			if originBefore != nil {
				return originBefore(cctx)
			}
			return nil
		}
	}
	ctx, span := trace.StartSpan(context.Background(), "/cli")
	defer span.End()

	app := &cli.App{
		Name:                 "star",
		Usage:                "star server",
		Version:              build.UserVersion(),
		EnableBashCompletion: true,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "repo",
				EnvVars: []string{"STAR_PATH"},
				Hidden:  true,
				Value:   "~/.star",
			},
		},

		Commands: append(local, lcli.Commands...),
	}

	app.Setup()
	app.Metadata["traceContext"] = ctx

	if err := app.Run(os.Args); err != nil {
		span.SetStatus(trace.Status{
			Code:    trace.StatusCodeFailedPrecondition,
			Message: err.Error(),
		})
		_, ok := err.(*lcli.ErrCmdFailed)
		if ok {
			log.Debugf("%+v", err)
		} else {
			log.Warnf("%+v", err)
		}
		os.Exit(1)
	}

}
