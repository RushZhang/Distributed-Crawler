package client

import (
	"crawler/engine"
	"log"
	"crawler_distributed/rpcsupport"
	"crawler_distributed/config"
)

func ItemSaver(host string) (chan engine.Item, error) {
	rpcClient, err := rpcsupport.NewClient(host)
	if err != nil { return nil, err}

	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <- out
			log.Printf("持久化得到ITEM %d: %v", itemCount, item)
			itemCount++

			//拿到item之后就建立rpc连接
			result := ""
			err := rpcClient.Call(config.ItemSaverRpc, item, &result)
			if err != nil {
				log.Printf("Saver在存%v时遇到了问题%v\n", item, err)
			}
		}
	}()
	return out, nil
}
