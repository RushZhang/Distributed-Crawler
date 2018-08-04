package main

import (
	"crawler_distributed/rpcsupport"
	"gopkg.in/olivere/elastic.v5"
	"crawler_distributed/persist"
	"fmt"
	"crawler_distributed/config"
)

func main() {

	err := serveRpc(fmt.Sprintf(":%d",config.ItemSaverPort), config.ElasticIndex)
	if err != nil {
		panic (err)
	}

}

func serveRpc(host, index string) error {
	esClient, err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil {
		return err
	}

	err = rpcsupport.NewServer(host, &persist.ItemSaverService{esClient, index})
	if err == nil {
		fmt.Println("ItemSaver的存储服务器在",config.ItemSaverPort,"端口监听中。。。")
	}
	return err
}
