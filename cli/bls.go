package cli

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

var blsCmd = &cli.Command{
	Name:  "bls",
	Usage: "bls function",
	Subcommands: []*cli.Command{
		blsCreatePrivateKey,
	},
}

var blsCreatePrivateKey = &cli.Command{
	Name:      "blsCreateblsCreatePrivateKey",
	Usage:     "Generate a new CreatePrivateKey",
	ArgsUsage: "",
	Action: func(cctx *cli.Context) error {

		fmt.Println("xxx")

		return nil
	},
}

//func main() {
//
//	logging.SetLogLevel("*", "info")
//
//	app := &cli.App{
//		Name:  "bls",
//		Usage: "bls function test",
//		Flags: []cli.Flag{
//			&cli.StringFlag{
//				Name: "createPrivateKey",
//			},
//		},
//		Action: func(cctx *cli.Context) error {
//
//			fmt.Println("bls test start")
//			//todo some thing
//
//			if cmd := cctx.String("createPrivateKey"); cmd != "" {
//				fmt.Println("createPrivateKey ")
//
//			}
//
//			fmt.Println("Hello world!2")
//
//			return nil
//		},
//	}
//
//	jaeger := setupJaegerTracing("lotus")
//	defer func() {
//		if jaeger != nil {
//			jaeger.Flush()
//		}
//	}()
//
//	ctx, span := trace.StartSpan(context.Background(), "/cli")
//	defer span.End()
//
//	app.Setup()
//	app.Metadata["traceContext"] = ctx
//	if err := app.Run(os.Args); err != nil {
//		span.SetStatus(trace.Status{
//			Code:    trace.StatusCodeFailedPrecondition,
//			Message: err.Error(),
//		})
//		os.Exit(1)
//	}
//
//}
//
//func setupJaegerTracing(serviceName string) *jaeger.Exporter {
//
//	//if _, ok := os.LookupEnv("LOTUS_JAEGER"); !ok {
//	//	return nil
//	//}
//	//agentEndpointURI := os.Getenv("LOTUS_JAEGER")
//	agentEndpointURI := "xxx.xxx.cn:6831"
//
//	je, err := jaeger.NewExporter(jaeger.Options{
//		AgentEndpoint: agentEndpointURI,
//		ServiceName:   serviceName,
//	})
//	if err != nil {
//		log.Errorw("Failed to create the Jaeger exporter", "error", err)
//		return nil
//	}
//
//	trace.RegisterExporter(je)
//	trace.ApplyConfig(trace.Config{
//		DefaultSampler: trace.AlwaysSample(),
//	})
//	return je
//}
