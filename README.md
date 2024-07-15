# 萤石云监控播放

- 需要安装`ffplay`作为播放客户端

> 安装
```shell
go install github.com/balrogsxt/ezviz-monitor@latest
```

> 运行

不设置环境变量的情况下
```shell
monitor -key=xxxxx -secret=xxxxx -device=xxxxxx -ffplay_path=ffplay
```

设置环境变量后可直接运行 `monitor`
