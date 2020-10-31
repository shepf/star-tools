// +build !nodaemon

package main

import (
	"context"
	"github.com/mitchellh/go-homedir"
	"github.com/multiformats/go-multiaddr"
	"github.com/shepf/star-tools/api"
	"github.com/shepf/star-tools/build"
	"github.com/shepf/star-tools/metrics"
	"github.com/shepf/star-tools/node"
	"github.com/shepf/star-tools/node/modules/dtypes"
	"github.com/shepf/star-tools/node/repo"
	"go.opencensus.io/tag"
	"os"
	"runtime/pprof"

	"github.com/urfave/cli/v2"
	"go.opencensus.io/plugin/runmetrics"
	"golang.org/x/xerrors"

	lcli "github.com/shepf/star-tools/cli"
)

var daemonStopCmd = &cli.Command{
	Name:  "stop",
	Usage: "Stop a running star daemon",
	Flags: []cli.Flag{},
	Action: func(cctx *cli.Context) error {
		api, closer, err := lcli.GetAPI(cctx)
		if err != nil {
			return err
		}
		defer closer()

		err = api.Shutdown(lcli.ReqContext(cctx))
		if err != nil {
			return err
		}

		return nil
	},
}

// DaemonCmd is the `star daemon` command
var DaemonCmd = &cli.Command{
	Name:  "daemon",
	Usage: "Start a lotus daemon process",
	Flags: []cli.Flag{
		&cli.StringFlag{
			Name:  "api",
			Value: "3333",
		},
		&cli.BoolFlag{
			Name:  "bootstrap",
			Value: true,
		},
	},
	Action: func(cctx *cli.Context) error {
		err := runmetrics.Enable(runmetrics.RunMetricOptions{
			EnableCPU:    true,
			EnableMemory: true,
		})
		if err != nil {
			return xerrors.Errorf("enabling runtime metrics: %w", err)
		}
		if prof := cctx.String("pprof"); prof != "" {
			profile, err := os.Create(prof)
			if err != nil {
				return err
			}

			if err := pprof.StartCPUProfile(profile); err != nil {
				return err
			}
			defer pprof.StopCPUProfile()
		}

		r, err := repo.NewFS(cctx.String("repo"))
		if err != nil {
			return xerrors.Errorf("opening fs repo: %w", err)
		}
		if err := r.Init(repo.FullNode); err != nil && err != repo.ErrRepoExists {
			return xerrors.Errorf("repo init error: %w", err)
		}

		ctx, _ := tag.New(context.Background(), tag.Insert(metrics.Version, build.BuildVersion), tag.Insert(metrics.Commit, build.CurrentCommit))
		{
			dir, err := homedir.Expand(cctx.String("repo"))
			if err != nil {
				log.Warnw("could not expand repo location", "error", err)
			} else {
				log.Infof("lotus repo: %s", dir)
			}
		}

		shutdownChan := make(chan struct{})
		var api api.FullNode

		stop, err := node.New(ctx,
			node.FullAPI(&api),
			node.Override(new(dtypes.ShutdownChan), shutdownChan),
			//node.Online(),

			node.Repo(r),

			node.ApplyIf(func(s *node.Settings) bool { return cctx.IsSet("api") },
				node.Override(node.SetApiEndpointKey, func(lr repo.LockedRepo) error {
					apima, err := multiaddr.NewMultiaddr("/ip4/127.0.0.1/tcp/" +
						cctx.String("api"))
					if err != nil {
						return err
					}
					return lr.SetAPIEndpoint(apima)
				})),
		)
		if err != nil {
			return xerrors.Errorf("initializing node: %w", err)
		}

		endpoint, err := r.APIEndpoint()

		return serveRPC(api, stop, endpoint, shutdownChan)

	},
	Subcommands: []*cli.Command{
		daemonStopCmd,
	},
}
