package main

import (
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
	"sync"
	"time"
)

func main() {

	var (
		cmd    *exec.Cmd
		output []byte
		err    error
	)

	filename := "star_monitor.txt"
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	var wg sync.WaitGroup

	//创建定时器，每隔600秒后，定时器就会给channel发送一个事件(当前时间)
	ticker := time.NewTicker(time.Second * 600)

	i := 0
	count := 0
	wg.Add(1)
	go func() {
		defer wg.Done()
		for { //循环
			<-ticker.C
			i++
			fmt.Println("i = ", i)
			// 生成Cmd
			cmd = exec.Command("/bin/bash", "-c", "lotus-miner sealing workers")
			// 执行了命令, 捕获了子进程的输出( pipe )
			if output, err = cmd.CombinedOutput(); err != nil {
				fmt.Println(err)
				return
			}
			//打印子进程的输出
			fmt.Println(string(output))

			//fmt.Println(strings.Contains(string(output),"] 0/32 core(s) in use"))
			if strings.Contains(string(output), "] 0/") {
				count++
				str := fmt.Sprintf("The number of call sectors pledge: %d", count)
				fmt.Println(str)

				// 生成Cmd
				cmd = exec.Command("/bin/bash", "-c", "lotus-miner sectors pledge")
				// 执行了命令, 捕获了子进程的输出( pipe )
				if output, err = cmd.CombinedOutput(); err != nil {
					fmt.Println(err)
					return
				}
				fmt.Println(string(output))

				n, err := io.WriteString(file, str)
				if err != nil {
					fmt.Println(n, err)
				}

			}

			//if i == 5 {
			//    ticker.Stop() //停止定时器
			//}
		}
	}()

	wg.Wait()

}
