package main

import (
	"testing"
	"crawler_distributed/rpcsupport"
	"crawler/engine"
	"crawler/model"
	"time"
	"crawler_distributed/config"
)

func TestItemSaver(t *testing.T) {
	const host = ":1234"

	//1. start ItemSaverServer
	go serveRpc(host, "test1")
	time.Sleep(time.Second)//如果不睡，可能main函数就完了server还没起来


	//2. start ItemSaverClient
	client, err := rpcsupport.NewClient(host)
	if err != nil {
		panic(err)
	}
	//3. Call Save
	item := engine.Item{
		Url:  "http://album.zhenai.com/u/108906739",
		Type: "zhenai",
		Id:   "108906739",
		Payload: model.Profile{
			Age:        34,
			Height:     162,
			Weight:     57,
			Income:     "3001-5000元",
			Gender:     "女",
			Name:       "安静的雪",
			Xinzuo:     "牡羊座",
			Occupation: "人事/行政",
			Marriage:   "离异",
			House:      "已购房",
			Hokou:      "山东菏泽",
			Education:  "大学本科",
			Car:        "未购车",
		},
	}
	result := ""

	//这个就是调用了rpc规范的函数，第一个是调用的方法，第二个是args，第三个是*result
	err = client.Call(config.ItemSaverRpc, item, &result)
	if err != nil || result != "ok" {
		t.Errorf("result: %s; err :%s", result, err)
	}
}