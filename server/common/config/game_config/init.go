package game_config

import "gserver_demo/server/common/config/game_config/level_up"

func Init() {
	_ = level_up.Config.Load()
}

func Reload() {
	Init()
}
