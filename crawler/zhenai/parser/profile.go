package parser

import (
	"video_sever/crawler/engine"
	"regexp"
	"strconv"
	"video_sever/crawler/model"
)

var GenderRe = regexp.MustCompile(`<div class="m-title" data-v-8b1eac0c>(\D+)的动态</div>`)
var AgeRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>(\d+)岁</div>`)
var HeightRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>(\d+)cm</div>`)
var WeightRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>(\d+)kg</div>`)
var IncomeRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>月收入:([^<]+)</div>`)
var marriageRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>(未婚|离异|丧偶)</div>`)
var EducationRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>(\D+)(士|科|中|学|下)</div>`)
var OccupationRe = regexp.MustCompile(`月收入:([^<]+)</div><div class="m-btn purple" data-v-8b1eac0c>([^<]+)</div>`)
var WorkPlaceRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>工作地:([^<]+)</div>`)
var ConstellationRe = regexp.MustCompile(`<div class="m-btn purple" data-v-8b1eac0c>(\D+)座([^<]+)</div>`)
var HouseRe = regexp.MustCompile(`<div class="m-btn pink" data-v-8b1eac0c>(\D+)(房|住+)</div>`)
var CarRe = regexp.MustCompile(`<div class="m-btn pink" data-v-8b1eac0c>(\D+)车</div>`)


func ParseProfile(contents []byte,name string) engine.ParserResult {
	profile := model.Profile{}
	profile.Name = name
	age,err := strconv.Atoi(extractString(contents,AgeRe,1))
	if err == nil{
		profile.Age = age
	}
	height,err := strconv.Atoi(extractString(contents,HeightRe,1))
	if err == nil{
		profile.Height = height
	}
	weight,err := strconv.Atoi(extractString(contents,WeightRe,1))
	if err == nil{
		profile.Weight = weight
	}
	profile.Marriage = extractString(contents,marriageRe,1)
	profile.Gender = extractString(contents,GenderRe,1)
	profile.Income = extractString(contents,IncomeRe,1)
	profile.Education = extractString(contents,EducationRe,1)
	profile.Occupation = extractString(contents,OccupationRe,2)
	profile.WorkPlace = extractString(contents,WorkPlaceRe,1)
	profile.Constellation = extractString(contents,ConstellationRe,1)
	profile.House = extractString(contents,HouseRe,1)
	profile.Car = extractString(contents,CarRe,1)
	if profile.Occupation == profile.Education{
		profile.Occupation = ""
	}
	if profile.Gender == "她"{
		profile.Gender = "女"
	}
	if profile.Gender == "他" {
		profile.Gender = "男"
	}
	result := engine.ParserResult{
		Items: [] interface{}{profile},
	}
	return result
}
func extractString(contents []byte, re *regexp.Regexp,num int) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2{
		return  string(match[num])
	}else {
		return ""
	}
}