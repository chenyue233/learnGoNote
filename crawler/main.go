package main

import (
	"video_sever/crawler/engine"
	"video_sever/crawler/zhenai/parser"
	"video_sever/crawler/scheduler"
	"video_sever/crawler/persist"
)

func main()  {
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan: persist.ItemSaver(),
	}
		e.Run(engine.Request{
			Url: "http://www.zhenai.com/zhenghun",
			ParserFunc: parser.ParseCityList,
		})
	}

