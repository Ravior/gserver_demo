package controller

import (
	"github.com/Ravior/gserver/net/gnet"
	"gserver_demo/proto/pb"
)

var Role = role{}

type role struct {
}

func (r *role) InitRouter(router *gnet.Router) {
	g := router.Group("role")
	g.AddRoute("login", r.login)
	g.AddRoute("heartBeat", r.heartBeat)
}

func (r *role) login(req *gnet.Request, reqMsg *pb.C2SLogin) {

}

// 处理心跳包
func (r *role) heartBeat(req *gnet.Request, reqMsg *pb.C2SHeartbeat) {
	//player := manager.Session.GetPlayerByConnId(req.GetConnId())
	//if player != nil {
	//	player.SendMsg(&pb.S2CHeartbeat{ServerTime: gtime.Now()})
	//}
}
