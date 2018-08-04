package persist

import (
	"crawler/engine"
	"gopkg.in/olivere/elastic.v5"
	"errors"
	"context"
	"log"
	"fmt"
)

type ItemSaverService struct {
	ESClient *elastic.Client
	Index    string
}


//这个就是符合rpc规范的函数，第一个是args，第二个是*result
func (s *ItemSaverService) Save(item engine.Item, result *string) error {
	fmt.Println("##############")
	err := SaveToElastic(s.ESClient, item, s.Index)
	fmt.Printf("RPC的Save函数： Item %v saved.", item)
	if err == nil {
		*result = "ok"
	} else {
		log.Printf("在rpc Save存储 %v 的时候遇到错误 %v", item, err)
	}
	return err
}


func SaveToElastic(client *elastic.Client, item engine.Item, index string) (err error) {
	if item.Type == "" {
		return errors.New("必须提供item的类型（ES的表名）")
	}

	//client.Index就是存数据的意思,再后边的Index相当于数据库名，type相当于表名，不加id就相当于让系统自动分配id,加id就自己分配
	IndexService := client.Index().Index(index).Type(item.Type).BodyJson(item)
	if item.Id != "" {
		IndexService.Id(item.Id)
	}

	_,  err = IndexService.Do(context.Background())

	if err != nil {
		return err
	}
	//fmt.Printf("%+v", resp)  //这样子打印结构体可以把结构体的字段名也打印出来

	return nil
}