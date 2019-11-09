package hook

import (
	"encoding/json"
	"fmt"
	"sync"
)

var scriptConf = scriptConfig{}

// ScriptConfig 库与脚本文件名的映射关系
type scriptConfig struct {
	sync.RWMutex
	data    bashMap
	repoMap map[string]repo
}
type bashMap map[string]string
type repo struct {
	Secret     string            `json:"secret"`
	ScriptPath string            `json:"script_path"`
	Event      map[string]string `json:"event"`
}

func (c *scriptConfig) Get(key string) (cmd string, err error) {
	c.RLock()
	defer c.RUnlock()

	cmd, ok := c.data[key]
	if !ok {
		err = fmt.Errorf("找不到对应command, key: %s", key)
	}
	return
}

// Set 目前计划不允许重复设置
func (c *scriptConfig) Set(key, val string) (cmd string, err error) {
	c.Lock()
	defer c.Unlock()

	cmd, ok := c.data[key]
	if !ok {
		c.data[key] = val
		cmd = val
	} else {
		err = fmt.Errorf("当前库 %s 已经存在对应的command: %s", key, cmd)
	}
	return
}

// Flash 重新加载配置
func (c *scriptConfig) Flash(data []byte) (err error) {
	c.Lock()
	defer c.Unlock()
	tmp := bashMap{}

	err = json.Unmarshal(data, &tmp)
	if err != nil {
		return err
	}
	c.data = tmp
	return
}
