package persist

import (
	"log"
	"github.com/olivere/elastic"
	"context"
	"video_sever/crawler/engine"
	"errors"
)


func ItemSaver() chan engine.Item {
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for{
			item := <-out
			log.Printf("Item Saver:got item #%d:%v",itemCount,item)
			itemCount++
			_,err := save(item)
			if err != nil{
				log.Printf("Item saver:err" + "saving item %v:%v",item ,err)
			}
		}
	}()
	return out
}

func save(item engine.Item) (id string,err error) {
	clien ,err := elastic.NewClient(
		// must turn off sniff in docker
		elastic.SetSniff(false))
	if err != nil{
		return "",err
	}
	if item.Type ==""{
		return "",errors.New("must supply Type ")
	}
	resp,err := clien.Index().
		Index("dating_profile").
		Type(item.Type).
		Id(item.Id).
		BodyJson(item).Do(context.Background())
	if err != nil{
		return "",err
	}
	return resp.Id,nil
}
