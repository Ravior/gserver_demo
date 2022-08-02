package level_up

import (
	"fmt"
	"github.com/Ravior/gserver/util/gconfig"
)

// 升级配置
var (
	Config     = &config{}
	configFile = "config/game_config/level_up.json"
)

type config struct {
	items map[int32]*ConfigItem
}

func (c *config) Load() error {
	c.items = make(map[int32]*ConfigItem)
	_items := make([]*ConfigItem, 0)
	if err := gconfig.LoadJsonConfigFromBasePath(configFile, &_items); err != nil {
		fmt.Printf("Load ConfigFile [%s] Fail, Error:%s\r\n", configFile, err)
		return err
	}

	for _, item := range _items {
		c.items[item.ID] = item
	}
	return nil
}

func (c *config) MaxLevel() int32 {
	// 写死最大等级100级
	return 100
}

func (c *config) Get(id int32) *ConfigItem {
	if id >= c.MaxLevel() {
		id = c.MaxLevel()
	}
	if item, ok := c.items[id]; ok {
		return item
	}
	return nil
}
