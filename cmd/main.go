package main

import (
	"fmt"
	"state-machine/conf"
	"state-machine/service"

	"github.com/gin-gonic/gin"
	"github.com/jessevdk/go-flags"
)

func main() {
	systemInit()
}

var opts struct {
	ConfigFile string `short:"c" required:"true" name:"config file"`
}

func systemInit() {
	if _, err := flags.Parse(&opts); err != nil {
		panic(err)
	}
	fmt.Printf("opts.ConfigFile: %v\n", opts.ConfigFile)
	// 读取配置文件
	conf.Init(opts.ConfigFile)

	// 开启http接口 调用任务
	r := gin.New()
	{
		// 默认任务
		r.POST("v0/default", service.DefaultTask)
	}
	err := r.Run(conf.Configure.Port)
	if err != nil {
		panic(err)
	}
}
