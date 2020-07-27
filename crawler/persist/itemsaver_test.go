package persist

import (
	"testing"
	"video_sever/crawler/model"
	"github.com/olivere/elastic"
	"context"
	"encoding/json"
)

func TestItemSaver(t *testing.T) {
	profile := model.Profile{
		Name:"凤求凰",
		Gender: "女",
		Age: 47,
		Height: 156,
		Income: "3千以下",
		Marriage: "丧偶",
		Education: "高中及以",
		Occupation: "销售总监",
		WorkPlace: "百色田林",
		Constellation: "魔羯",
		House: "和家人同",
		Car: "未买",
	}
	id,err := save(profile)
	if err != nil{
		panic(err)
	}
	clien,err := elastic.NewClient(elastic.SetSniff(false))
	if err != nil{
		panic(err)
	}
	resp,err := clien.Get().Index(
		"dating_profile").
		Type("zhenai").
		Id(id).
		Do(context.Background())
	if err != nil{
		panic(err)
	}
	t.Logf("%+v",resp)
	var actual model.Profile
	err = json.Unmarshal(*resp.Source,&actual)
	if err != nil{
		panic(err)
	}
	if actual != profile{
		t.Errorf("got %v;expected %v",actual,profile)
	}
}
