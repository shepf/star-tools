# star-tools
star-tools 包含一个jsonRPC服务程序，以及一些常用工具。
用于提供一些filestar/filecoin/lotus相关非代码侵入性的相关常用工具或小程序。


star: 是jsonRPC服务程序，常驻进程。后面可以添加扩展很多拓展功能。
star-monitor: 用于监控 worker工作状态，自动根据woker工作状态，调 miner 封装命令。


大家有什么想要的filestar/filecoin/lotus相关小程序和工具，欢迎提issue！


# 编译安装
make all
make install


```
//查看star版本
star --version
```

```
//启动star
nohup star daemon  > ~/nohup.star.output 2>&1 &
//查看日志 
//tail -f ~/nohup.star.output
```

```
//启动 用于监听，自动封装随机扇区
nohup star-monitor   > ~/star-monitor.output 2>&1 &
//查看日志
//tail -f ~/star-monitor.output
```

```
//停止star
star daemon stop

//停止star-monitor
ps axf |grep star-monitor
kill xxx
```

# star使用说明
## 配置监听ip
vi /root/.star/config.toml
如下编辑，star 监听的ip地址，注意：默认只监听127.0.0.1
```
# Default config:
#
[API]
ListenAddress = "/ip4/192.168.31.39/tcp/3333/http"
#  #  RemoteListenAddress = ""
#  #  Timeout = "30s"
#
```

## 获取Star版本
```
curl -X POST \
     -H "Content-Type: application/json" \
     -H "Authorization: Bearer $(cat ~/.star/token)" \
     --data '{ "jsonrpc": "2.0", "method": "Star.Version", "params": [], "id": 3 }' \
     'http://127.0.0.1:3333/rpc/v0'
```

## 查看miner信息
```
curl -X POST \
     -H "Content-Type: application/json" \
     -H "Authorization: Bearer $(cat ~/.star/token)" \
     --data '{ "jsonrpc": "2.0", "method": "Star.MinerInfo", "params": [], "id": 3 }' \
     'http://127.0.0.1:3333/rpc/v0'  |jq
```

## 查看workersList
```
curl -X POST \
     -H "Content-Type: application/json" \
     -H "Authorization: Bearer $(cat ~/.star/token)" \
     --data '{ "jsonrpc": "2.0", "method": "Star.WorkersList", "params": [], "id": 3 }' \
     'http://127.0.0.1:3333/rpc/v0'
```

## 获取系统运行时间
```
curl -X POST \
     -H "Content-Type: application/json" \
     -H "Authorization: Bearer $(cat ~/.star/token)" \
     --data '{ "jsonrpc": "2.0", "method": "Star.SysUptime", "params": [], "id": 3 }' \
     'http://127.0.0.1:3333/rpc/v0'
```
