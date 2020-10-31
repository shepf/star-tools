package common

import (
	"context"
	"fmt"
	"os/exec"

	logging "github.com/ipfs/go-log/v2"

	"github.com/gbrlsnchs/jwt/v3"
	"go.uber.org/fx"
	"golang.org/x/xerrors"

	"github.com/filecoin-project/go-jsonrpc/auth"

	"github.com/shepf/star-tools/api"
	"github.com/shepf/star-tools/build"
	"github.com/shepf/star-tools/node/modules/dtypes"
)

type CommonAPI struct {
	fx.In

	APISecret    *dtypes.APIAlg
	ShutdownChan dtypes.ShutdownChan
}

//获取系统运行时间
func (a *CommonAPI) SysUptime(ctx context.Context) (string, error) {
	fmt.Println("SysUptime start:")
	var (
		cmd    *exec.Cmd
		output []byte
		err    error
	)

	// 生成Cmd
	cmd = exec.Command("/bin/bash", "-c", "cat /proc/uptime| awk -F. '{run_days=$1 / 86400;run_hour=($1 % 86400)/3600;run_minute=($1 % 3600)/60;run_second=$1 % 60;printf(\"系统已运行: %d天%d时%d分%d秒\",run_days,run_hour,run_minute,run_second)}'")
	// 执行了命令, 捕获了子进程的输出( pipe )
	if output, err = cmd.CombinedOutput(); err != nil {
		fmt.Println(err)
		return "", err
	}
	//打印子进程的输出
	fmt.Println(string(output))

	return string(output), nil
}

type jwtPayload struct {
	Allow []auth.Permission
}

func (a *CommonAPI) AuthVerify(ctx context.Context, token string) ([]auth.Permission, error) {
	var payload jwtPayload
	if _, err := jwt.Verify([]byte(token), (*jwt.HMACSHA)(a.APISecret), &payload); err != nil {
		return nil, xerrors.Errorf("JWT Verification failed: %w", err)
	}

	return payload.Allow, nil
}

func (a *CommonAPI) AuthNew(ctx context.Context, perms []auth.Permission) ([]byte, error) {
	p := jwtPayload{
		Allow: perms, // TODO: consider checking validity
	}

	return jwt.Sign(&p, (*jwt.HMACSHA)(a.APISecret))
}

func (a *CommonAPI) Version(context.Context) (api.Version, error) {
	return api.Version{
		Version:    build.UserVersion(),
		APIVersion: build.APIVersion,
	}, nil
}

func (a *CommonAPI) LogList(context.Context) ([]string, error) {
	return logging.GetSubsystems(), nil
}

func (a *CommonAPI) LogSetLevel(ctx context.Context, subsystem, level string) error {
	return logging.SetLogLevel(subsystem, level)
}

func (a *CommonAPI) Shutdown(ctx context.Context) error {
	a.ShutdownChan <- struct{}{}
	return nil
}

func (a *CommonAPI) Closing(ctx context.Context) (<-chan struct{}, error) {
	return make(chan struct{}), nil // relies on jsonrpc closing
}

var _ api.Common = &CommonAPI{}
