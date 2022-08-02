package game_server

import (
	"github.com/Ravior/gserver/net/gnet"
	"github.com/Ravior/gserver/net/gwebsocket"
	"github.com/Ravior/gserver/os/glog"
	"gserver_demo/server/app"
	"gserver_demo/server/common/config/game_config"
	"gserver_demo/server/game_server/controller"
	"gserver_demo/server/game_server/manager"
	"os"
	"time"
)

func Init() {
	// 启动WebSocket服务器
	wsServer := gwebsocket.NewServer()
	wsServer.SetOnConnStart(onClientConnStart)
	wsServer.SetOnConnStop(onClientConnStop)

	app.SetServer(wsServer)
	app.SetServerStartFunc(onServerStart)
	app.SetServerCloseFunc(onServerClose)
	app.SetServerQuitFunc(onServerQuit)
	app.SetReloadConfigFunc(onReloadConfig)

	// 启动WebSocket服务器
	go wsServer.Run()

	// 1秒钟后如果服务器运行正常，则执行服务器启动回调
	time.AfterFunc(1*time.Second, func() {
		if app.GetServerStartFunc() != nil {
			app.GetServerStartFunc()()
		}
	})
}

// 初始化WebSocket消息路由
func intRouter(router *gnet.Router) {
	controller.Role.InitRouter(router)
}

func onClientConnStart(conn gnet.IConnection) {
	glog.Infof("ClientConnStart, Addr:%s, ConnId:%d, 当前连接数:%d", conn.RemoteAddr().String(), conn.GetConnID(), conn.GetSocket().GetConnMgr().Len())
}

func onClientConnStop(conn gnet.IConnection) {
	glog.Infof("ClientConnStop, Addr:%s, ConnId:%d, 当前连接数:%d", conn.RemoteAddr().String(), conn.GetConnID(), conn.GetSocket().GetConnMgr().Len())
	manager.Session.RemovePlayer(conn.GetConnID())
}

func onServerStart() {
	glog.Info("Game Server Start")
	// 初始化管理器
	manager.Init()
	glog.Info("服务器初始化完成, 可以进行链接")
	// 标记服务器未开启状态
	app.MyApp.ServerStatus = app.ServerStatusRunning
}

// 服务器关闭回调(kill -s SIGTERM 触发, 5分钟后停止服务器)
func onServerClose() {
	glog.Info("触发服务器关闭操作")

	// 服务器标记为关闭状态
	app.MyApp.ServerStatus = app.ServerStatusCloseWait

	ch := make(chan bool)

	// 5分钟后关闭服务器
	time.AfterFunc(300*time.Second, func() {
		// 执行服务器退出
		onServerQuit()
		ch <- true
	})

	select {
	case <-ch:
		glog.Info("服务器程序立马退出")
	}

	os.Exit(0)
}

// 服务器退出(kill -s SIGQUIT 触发, 立马停止服务器)
func onServerQuit() {
	glog.Info("=========###执行服务器退出回调-开始###=========")
	glog.Infof("服务器即将进行关闭, 当前连接数：%d", app.GetServer().GetConnMgr().Len())
	app.MyApp.ServerStatus = app.ServerStatusClose
	// 清理服务器资源
	manager.Clear()
	glog.Info("=========^^^执行服务器退出回调-结束^^^=========")
}

func onReloadConfig() {
	glog.Info("ReloadConfig")
	game_config.Reload()
}
