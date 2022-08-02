package manager

import (
	"github.com/Ravior/gserver/net/gnet"
	"github.com/Ravior/gserver/os/glog"
	"sync"
)

// Session 会话管理器
var Session = &session{}

type session struct {
	connMap sync.Map
}

func (s *session) AddPlayer(roleId int32, conn gnet.IConnection) {
	s.connMap.Store(conn.GetConnID(), roleId)
}

func (s *session) RemovePlayer(connId uint32) {
	glog.Infof("[链接管理池]准备删除连接对象，ConnID:%d", connId)
	s.connMap.Delete(connId)
}
