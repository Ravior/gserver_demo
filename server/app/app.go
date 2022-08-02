package app

import (
	"errors"
	"github.com/Ravior/gserver/net/gnet"
	"github.com/Ravior/gserver/os/glog"
	"github.com/Ravior/gserver/util/geventbus"
	"github.com/Ravior/gserver/util/gutil"
	"github.com/golang/protobuf/proto"
	"go.uber.org/atomic"
)

var (
	MyApp = &app{
		eventBus:     geventbus.New(),
		ServerStatus: ServerStatusInit,
	}
)

const (
	ServerStatusInit      = 0 // 服务器初始化
	ServerStatusRunning   = 1 // 服务器运行中
	ServerStatusCloseWait = 2 // 服务器关闭等待中
	ServerStatusClose     = 3 // 服务器关闭
)

type app struct {
	serverStartFunc  func()        // 服务器启动回调
	serverCloseFunc  func()        // 关闭服务器回调
	serverQuitFunc   func()        // 服务器退出回调(kill -3调用)
	reloadConfigFunc func()        // 重新加载配置回调
	server           gnet.ISocket  // 服务器对象
	eventBus         geventbus.Bus // 全局事件总线
	OnlineNum        atomic.Int32  // 在线人数
	ServerStatus     int           // 服务器状态
}

// PostEvent 发布事件
func (a *app) PostEvent(event string, args ...interface{}) {
	a.eventBus.Publish(event, args...)
}

// ListenEvent 订阅事件
func (a *app) ListenEvent(event string, fn interface{}) {
	_ = a.eventBus.SubscribeAsync(event, fn, false)
}

// UnListenEvent 取消订阅事件
func (a *app) UnListenEvent(event string, fn interface{}) {
	_ = a.eventBus.Unsubscribe(event, fn)
}

// SetReloadConfigFunc 设置重新加载配置回调
func SetReloadConfigFunc(f func()) {
	MyApp.reloadConfigFunc = f
}

// GetReloadConfigFunc 获取重新加载配置回调
func GetReloadConfigFunc() func() {
	return MyApp.reloadConfigFunc
}

// SetServerStartFunc 设置程序关闭回调
func SetServerStartFunc(f func()) {
	MyApp.serverStartFunc = f
}

// GetServerStartFunc 获取程序关闭回调
func GetServerStartFunc() func() {
	return MyApp.serverStartFunc
}

// SetServerCloseFunc 设置程序关闭回调
func SetServerCloseFunc(f func()) {
	MyApp.serverCloseFunc = f
}

// GetServerCloseFunc 获取程序关闭回调
func GetServerCloseFunc() func() {
	return MyApp.serverCloseFunc
}

// SetServerQuitFunc 设置程序退出回调
func SetServerQuitFunc(f func()) {
	MyApp.serverQuitFunc = f
}

// GetServerQuitFunc 获取程序退出回调
func GetServerQuitFunc() func() {
	return MyApp.serverQuitFunc
}

// SetServer 设置当前服务器对象
func SetServer(s gnet.ISocket) {
	MyApp.server = s
}

// GetServer 获取当前服务器对象
func GetServer() gnet.ISocket {
	return MyApp.server
}

// ConnSendMsg Conn发送消息
func ConnSendMsg(conn gnet.IConnection, msg proto.Message) error {
	// 所有服务器加载配置
	msgID, msgData := gutil.ProtoMarshal(msg)

	if conn == nil {
		glog.Error("connection is nil")
		return errors.New("connection is nil")
	}

	// 调用SendMsg发包
	if err := conn.SendMsg(msgID, msgData); err != nil {
		glog.Error("SendMsg error !")
		return err
	}
	return nil
}

// BroadcastMsg 广播消息
func BroadcastMsg(msg proto.Message) {
	_server := GetServer()
	if _server == nil {
		return
	}

	msgID, msgData := gutil.ProtoMarshal(msg)
	glog.Debugf("[BroadcastMsg]MsgID:%d", msgID)
	_server.GetConnMgr().BroadcastMsg(msgID, msgData)
}
