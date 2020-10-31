package full

import (
	"context"
	"fmt"
	"go.uber.org/fx"
	"os/exec"
)

type MonitorAPI struct {
	fx.In

	///	Running bool
	//	RunInfo string
}

func (m *MonitorAPI) WorkersList(ctx context.Context) (string, error) {
	fmt.Println("WorkersList start:")
	var (
		cmd    *exec.Cmd
		output []byte
		err    error
	)

	// 生成Cmd
	cmd = exec.Command("/bin/bash", "-c", "lotus-miner sealing workers list")
	// 执行了命令, 捕获了子进程的输出( pipe )
	if output, err = cmd.CombinedOutput(); err != nil {
		fmt.Println(err)
		return "", err
	}
	//打印子进程的输出
	fmt.Println(string(output))

	return string(output), nil
}

func (m *MonitorAPI) MinerInfo(ctx context.Context) (string, error) {
	fmt.Println("MinerInfo start:")
	var (
		cmd    *exec.Cmd
		output []byte
		err    error
	)

	// 生成Cmd
	cmd = exec.Command("/bin/bash", "-c", "lotus-miner info")
	// 执行了命令, 捕获了子进程的输出( pipe )
	if output, err = cmd.CombinedOutput(); err != nil {
		fmt.Println(err)
		return "", err
	}
	//打印子进程的输出
	fmt.Println(string(output))

	return string(output), nil
}
