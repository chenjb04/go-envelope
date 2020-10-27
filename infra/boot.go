package infra

import "github.com/tietang/props/kvs"

//应用程序启动管理器
type BootApplication struct {
	conf           kvs.ConfigSource
	starterContext StarterContext
}

//初始化BootApplication
func New(conf kvs.ConfigSource) *BootApplication {
	b := &BootApplication{
		conf:           conf,
		starterContext: StarterContext{},
	}
	b.starterContext[KeyProps] = conf
	return b
}

func (b *BootApplication) Start() {
	// 初始化Starter
	b.init()
	// 安装Starter
	b.setup()
	// 启动starter
	b.start()
}

func (b *BootApplication) init() {
	for _, starter := range StarterRegister.AllStarters() {
		starter.Init(b.starterContext)
	}
}

func (b *BootApplication) setup() {
	for _, starter := range StarterRegister.AllStarters() {
		starter.Setup(b.starterContext)
	}
}

func (b *BootApplication) start() {
	for i, starter := range StarterRegister.AllStarters() {
		if starter.StartBlocking() {
			// 如果最后一个是可阻塞的，那么直接启动并阻塞
			if i+1 == len(StarterRegister.AllStarters()) {
				starter.Start(b.starterContext)
			} else { //如果不是，使用协程异步启动，防止阻塞后面的starter
				go starter.Start(b.starterContext)
			}
		} else {
			starter.Start(b.starterContext)
		}

	}
}
