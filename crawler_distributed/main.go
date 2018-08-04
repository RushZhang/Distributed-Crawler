package main

import (
	"crawler/engine"
	"crawler/zhenai/parser"
	"crawler/scheduler"
	"crawler_distributed/persist/client"
	"fmt"
	"crawler_distributed/config"
)



func main() {
	itemChan, err := client.ItemSaver(fmt.Sprintf(":%d",config.ItemSaverPort))
	if err != nil {  //如果itemSaver都起不来，那爬虫没有意义
		panic(err)
	}


	//要改为并发爬虫就写engine.ConcurrentEngine{}.Run()
	e := engine.ConcurrentEngine{
		Scheduler: &scheduler.SimpleScheduler{},
		WorkerCount: 10,
		ItemChan: itemChan,
	}
	//e := engine.SimpleEngine{}

	//e.Run(engine.Request{
	//	Url: "http://www.zhenai.com/zhenghun",
	//	ParserFunc: parser.ParseCityList,
	//})
	e.Run(engine.Request{
		Url: "http://www.zhenai.com/zhenghun/boertala",
		ParserFunc: parser.ParseCity,
	})

}





