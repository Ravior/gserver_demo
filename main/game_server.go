package main

import (
	"fmt"
	"github.com/Ravior/gserver/os/glog"
	"github.com/Ravior/gserver/util/gconfig"
	"gserver_demo/main/internal"
	"gserver_demo/server/common/config/game_config"
	"gserver_demo/server/game_server"
	"net/http"
	"os"
)

// 游戏服务器
// 采用WebSocket协议
func main() {
	// 加载游戏服务器配置
	if err := gconfig.Global.Load("config/server/game_server.json"); err != nil {
		glog.Error("配置数据读取异常，登陆服务器启动失败", err)
		os.Exit(0)
		return
	}

	if len(os.Args) <= 1 {
		glog.Error("参数[服务器ID]缺失，无法启动游戏服服务器")
		os.Exit(0)
		return
	}

	// 获取传入的ServerId
	serverId := os.Args[1]

	// 初始化基础模块
	internal.InitBaseModule(serverId)

	// 初始化游戏配置表
	game_config.Init()

	// 打印系统信息
	internal.PrintSystemInfo()

	// 启动在线Pprof
	if gconfig.Global.Pprof.Port > 0 {
		go func() {
			httpAddr := fmt.Sprintf("%s:%d", gconfig.Global.Pprof.IP, gconfig.Global.Pprof.Port)
			glog.Infof("启动Pprof Http服务, Addr: %s", httpAddr)
			if err := http.ListenAndServe(httpAddr, nil); err != nil {
				glog.Warnf("启动Pprof Http服务失败, Addr: %s, Err: %v", httpAddr, err)
			}
		}()
	}

	// 初始化游戏逻辑服
	game_server.Init()

	// 等待重新加载配置信号
	internal.WaitReloadConfigSignal()

	// 等待关闭信息
	internal.WaitStopSignal()
}
