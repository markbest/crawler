package main

import (
	"github.com/markbest/crawler/conf"
	"github.com/markbest/crawler/engine"
	"github.com/markbest/crawler/scheduler"
	"github.com/markbest/crawler/zhenai"
)

var targetUrl = "http://www.zhenai.com/zhenghun"

func main() {
	if err := conf.InitConfig(); err != nil {
		panic(err)
	}

	e := engine.Engine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 100,
	}
	e.Run(engine.Request{
		Url:       targetUrl,
		ParseFunc: zhenai.ParseCityList,
	})
}
