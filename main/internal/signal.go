package internal

import (
	"github.com/Ravior/gserver/os/glog"
	"gserver_demo/server/app"
	"os"
	"os/signal"
	"syscall"
)

// WaitStopSignal 阻塞，只有执行信号才执行
func WaitStopSignal() {
	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM)
	sig := <-osSignals
	glog.Infof("收到关闭信号, 准备关闭服务器, sig:%d", sig)
	if sig == syscall.SIGINT {
		glog.Info("Ctrl+C直接关闭服务器")
		close(osSignals)
		os.Exit(0)
	} else if sig == syscall.SIGQUIT {
		glog.Info("kill -s SIGQUIT 保存数据后关闭服务器")
		// 执行退出服务器回调
		f := app.GetServerQuitFunc()
		if f != nil {
			f()
		}
		close(osSignals)
	} else {
		glog.Info("kill -s SIGTERM 广播停服消息，5分钟后关闭服务器")
		// 执行关闭程序
		f := app.GetServerCloseFunc()
		if f != nil {
			f()
		}
		glog.Info("关闭服务器")
		close(osSignals)
		os.Exit(0)
	}
}

// WaitReloadConfigSignal 等待加载配置信号
func WaitReloadConfigSignal() {
	osSignals := make(chan os.Signal, 1)
	signal.Notify(osSignals, syscall.SIGUSR1)
	go func() {
		sig := <-osSignals
		glog.Debug("收到重新加载配置信号, sig:", sig)

		// 执行关闭程序
		f := app.GetReloadConfigFunc()
		if f != nil {
			f()
		}
	}()

}
