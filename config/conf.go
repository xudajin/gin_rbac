package config

import (
	"fmt"
	"log"

	"github.com/BurntSushi/toml"
)

type Conf struct {
	DB     database `toml:"database"`
	Mode   string   `toml:"mode"`
	JwtKey string   `toml:"jwt_key"`
}
type database struct {
	LoginName     string `toml:"name"`
	LoginPassword string `toml:"password"`
	Type          string `toml:"database_type"`
	DatabaseName  string `toml:"database_name"`
}

var Set *Conf

func init() {
	var conf Conf
	_, err := toml.DecodeFile("./config/conf.toml", &conf) // 以入口文件为当前路径
	if err != nil {
		log.Fatalf("获取配置失败: %s\n", err)
		return
	}
	Set = &conf
	fmt.Println("获取系统配置成功")
}
