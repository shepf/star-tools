#! /bin/bash
##########################################################
########              重启lotus               #############
##########################################################
# 注掉
#添加这句 某个命令执行失败 就不会继续往下执行，grep没有返回也算失败
#$(ps aux | grep "lotus-miner" | grep -v "grep" | awk '{print $2}')
#set -euo pipefail

echo "start update.."

echo "stop lotus-miner..."
miner_pid_list=$(ps aux | grep "lotus-miner" | grep -v "grep" | awk '{print $2}')
if [ "$miner_pid_list" = "" ]; then
        echo "lotus not running" 
else
        echo "miner_pid_list: $miner_pid_list"
        `lotus-miner stop`
        sleep 1
fi

echo "stop lotus daemon..." 
lotus_daemon_pid_list=`ps aux | grep "lotus daemon" | grep -v "grep" | awk '{print $2}'`
if [ "$lotus_daemon_pid_list" = "" ]; then
        echo "lotus not running" 
else
        echo "lotus_daemon_pid_list: $lotus_daemon_pid_list"
        `lotus daemon stop`
        sleep 1
fi


##cd /data/
nohup lotus daemon > ~/lotus.output 2>&1 &
echo ""
echo "start lotus daemon success"
sleep 1
nohup  lotus-miner run   > ~/lotus-miner.output 2>&1 &
echo "start lotus-miner run success"
