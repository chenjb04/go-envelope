/*
统一的配置工具库，将各种配置源抽象或转换为类似properties格式的key/value，并提供统一的API来访问这些key/value。支持 properties 文件、
ini 文件、zookeeper k/v、zookeeper k/props、consul k/v、consul k/props等配置源，并且支持通过 Unmarshal从配置中抽出struct；
支持上下文环境变量的eval，${}形式；支持多种配置源组合使用。
*/
package base

import (
	"fmt"
	"github.com/tietang/props/kvs"
	"go-envelope/infra"
)

var props kvs.ConfigSource

type PropsStarter struct {
	infra.BaseStarter
}

func Props() kvs.ConfigSource {
	return props
}

func (p *PropsStarter) Init(ctx infra.StarterContext) {
	props = ctx.Props()
	fmt.Println("初始化配置")
}
