package apistruct

import (
	"context"
	"github.com/libp2p/go-libp2p-core/peer"

	"github.com/filecoin-project/go-jsonrpc/auth"

	"github.com/shepf/star-tools/api"
)

// All permissions are listed in permissioned.go
var _ = AllPermissions

type CommonStruct struct {
	Internal struct {
		AuthVerify func(ctx context.Context, token string) ([]auth.Permission, error) `perm:"read"`
		AuthNew    func(ctx context.Context, perms []auth.Permission) ([]byte, error) `perm:"admin"`

		SysUptime func(context.Context) (string, error)      `perm:"read"`
		Version   func(context.Context) (api.Version, error) `perm:"read"`

		LogList     func(context.Context) ([]string, error)     `perm:"write"`
		LogSetLevel func(context.Context, string, string) error `perm:"write"`

		Shutdown func(context.Context) error                    `perm:"admin"`
		Closing  func(context.Context) (<-chan struct{}, error) `perm:"read"`
	}
}

// FullNodeStruct implements API passing calls to user-provided function values.
type FullNodeStruct struct {
	CommonStruct

	Internal struct {
		//star function
		WorkersList func(context.Context) (string, error) `perm:"read"`
		MinerInfo   func(context.Context) (string, error) `perm:"read"`
	}
}

func (f FullNodeStruct) ID(ctx context.Context) (peer.ID, error) {
	panic("implement me")
}

// CommonStruct

func (c *CommonStruct) AuthVerify(ctx context.Context, token string) ([]auth.Permission, error) {
	return c.Internal.AuthVerify(ctx, token)
}

func (c *CommonStruct) AuthNew(ctx context.Context, perms []auth.Permission) ([]byte, error) {
	return c.Internal.AuthNew(ctx, perms)
}

func (c *CommonStruct) SysUptime(ctx context.Context) (string, error) {
	return c.Internal.SysUptime(ctx)
}

// Version implements API.Version
func (c *CommonStruct) Version(ctx context.Context) (api.Version, error) {
	return c.Internal.Version(ctx)
}

func (c *CommonStruct) LogList(ctx context.Context) ([]string, error) {
	return c.Internal.LogList(ctx)
}

func (c *CommonStruct) LogSetLevel(ctx context.Context, group, level string) error {
	return c.Internal.LogSetLevel(ctx, group, level)
}

func (c *CommonStruct) Shutdown(ctx context.Context) error {
	return c.Internal.Shutdown(ctx)
}

func (c *CommonStruct) Closing(ctx context.Context) (<-chan struct{}, error) {
	return c.Internal.Closing(ctx)
}

// FullNodeStruct
func (c *FullNodeStruct) WorkersList(ctx context.Context) (string, error) {
	return c.Internal.WorkersList(ctx)
}
func (c *FullNodeStruct) MinerInfo(ctx context.Context) (string, error) {
	return c.Internal.MinerInfo(ctx)
}

var _ api.Common = &CommonStruct{}
var _ api.FullNode = &FullNodeStruct{}
