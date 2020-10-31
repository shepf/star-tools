package main

import (
	"context"
	"github.com/filecoin-project/go-jsonrpc"
	"github.com/filecoin-project/go-jsonrpc/auth"
	"github.com/shepf/star-tools/api/apistruct"
	"github.com/shepf/star-tools/node"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"syscall"

	logging "github.com/ipfs/go-log/v2"
	"github.com/multiformats/go-multiaddr"
	manet "github.com/multiformats/go-multiaddr-net"
	"golang.org/x/xerrors"

	"contrib.go.opencensus.io/exporter/prometheus"

	"github.com/shepf/star-tools/api"
)

var log = logging.Logger("main")

func serveRPC(a api.FullNode, stop node.StopFunc, addr multiaddr.Multiaddr, shutdownCh <-chan struct{}) error {
	rpcServer := jsonrpc.NewServer()
	rpcServer.Register("Star", apistruct.PermissionedFullAPI(a))

	ah := &auth.Handler{
		Verify: a.AuthVerify,
		Next:   rpcServer.ServeHTTP,
	}

	http.Handle("/rpc/v0", ah)

	exporter, err := prometheus.NewExporter(prometheus.Options{
		Namespace: "star",
	})
	if err != nil {
		log.Fatalf("could not create the prometheus stats exporter: %v", err)
	}

	http.Handle("/debug/metrics", exporter)

	log.Info("rpc addr: ", addr)
	lst, err := manet.Listen(addr)
	if err != nil {
		return xerrors.Errorf("could not listen: %w", err)
	}

	srv := &http.Server{Handler: http.DefaultServeMux}

	sigCh := make(chan os.Signal, 2)
	go func() {
		select {
		case <-sigCh:
		case <-shutdownCh:
		}

		log.Warn("Shutting down...")
		if err := srv.Shutdown(context.TODO()); err != nil {
			log.Errorf("shutting down RPC server failed: %s", err)
		}
		//if err := stop(context.TODO()); err != nil {
		//	log.Errorf("graceful shutting down failed: %s", err)
		//}
		log.Warn("Graceful shutdown successful")
	}()
	signal.Notify(sigCh, syscall.SIGTERM, syscall.SIGINT)

	return srv.Serve(manet.NetListener(lst))
}
