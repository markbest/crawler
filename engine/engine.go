package engine

import (
	"fmt"
	"github.com/markbest/crawler/conf"
	"github.com/markbest/crawler/utils"
	"log"
	"os"
)

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

type Engine struct {
	Scheduler   Scheduler
	WorkerCount int
}

func init() {
	if err := conf.InitConfig(); err != nil {
		panic(err)
	}
}

func (e *Engine) Run(seeds ...Request) {
	out := make(chan ParseResult)
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		e.CreateWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	itemCount := 0
	esClient := utils.NewES(conf.Conf.Zhenai.ElasticUrl, log.New(os.Stdout, "", log.LstdFlags))
	esStorage := NewStorage(esClient, conf.Conf.Zhenai.ElasticIndex, conf.Conf.Zhenai.ElasticType)
	for {
		result := <-out
		if result.Items != nil {
			for _, item := range result.Items {
				go func() {
					esStorage.Save(&item)
				}()
				itemCount++
				fmt.Printf("Got #%d item %v\n", itemCount, item)

			}
		}

		for _, request := range result.Requests {
			e.Scheduler.Submit(request)
		}
	}
}

func (e *Engine) CreateWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	go func() {
		for {
			ready.WorkerReady(in)
			request := <-in
			result, err := Work(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
