package api

import (
	"context"
	"fmt"

	"github.com/filecoin-project/go-jsonrpc/auth"

	"github.com/shepf/star-tools/build"
)

type Common interface {

	// MethodGroup: Auth

	AuthVerify(ctx context.Context, token string) ([]auth.Permission, error)
	AuthNew(ctx context.Context, perms []auth.Permission) ([]byte, error)

	// MethodGroup: Sys powerd by Star
	// 获取系统运行时间
	SysUptime(context.Context) (string, error)
	// TODO CPUINFO MEMINFO

	// MethodGroup: Common

	// Version provides information about API provider
	Version(context.Context) (Version, error)

	LogList(context.Context) ([]string, error)
	LogSetLevel(context.Context, string, string) error

	// trigger graceful shutdown
	Shutdown(context.Context) error

	Closing(context.Context) (<-chan struct{}, error)
}

// Version provides various build-time information
type Version struct {
	Version string

	// APIVersion is a binary encoded semver version of the remote implementing
	// this api
	//
	// See APIVersion in build/version.go
	APIVersion build.Version
}

func (v Version) String() string {
	return fmt.Sprintf("%s+api%s", v.Version, v.APIVersion.String())
}
