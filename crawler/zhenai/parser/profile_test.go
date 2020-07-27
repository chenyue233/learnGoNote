package parser

import (
	"testing"
	"io/ioutil"
	"video_sever/crawler/model"
)

func TestParseProfile(t *testing.T) {
	contents,err := ioutil.ReadFile("profile_test_data.html")
	if err != nil{
		panic(err)
	}
	result := ParseProfile(contents,"凤求凰")
	if len(result.Items) != 1{
		t.Errorf("Items should contain 1" + "element;but was %v",result.Items)
	}
	profile := result.Items[0].(model.Profile)

	expected := model.Profile{
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
	if profile != expected{
		t.Errorf("execpted %v;but was %v",expected,profile)
	}
}
