package main

import (
	"constant"
	"encoding/json"
	"fmt"
	"strconv"
	"util"
)

type Match struct {
}

/**
	查询合约列表
 */
func (match *Match) GetMatchList() {
	args := make(map[string]string)
	args[constant.ChainCodeName] = constant.ChainCodeMatching
	args[constant.Version] = constant.MatchingVersion
	args[constant.Function] = constant.MatchingList
	args[constant.AppId] = constant.AppIdGravity

	res := util.PostTest(constant.UrlQuery, args)
	var ccs []string
	json.Unmarshal(res, &ccs)
	ids := make(map[string]string)
	for index, value := range ccs {
		ids[strconv.Itoa(index)] = value
		fmt.Println("合约"+strconv.Itoa(index)+":", value,"  ID:",strconv.Itoa(index))
	}
	util.SaveId(ids, constant.ChainCodeFile)
}

/**
	执行匹配
 */
func (match *Match) Match(city, id, lowPrice, highPrice string) {
	ccid := util.ReadId(constant.ChainCodeFile)
	args := make(map[string]string)
	args[constant.ChainCodeName] = constant.ChainCodeInfo
	args[constant.Version] = constant.InfoVersion
	args[constant.Function] = constant.Matching
	args[constant.AppId] = constant.AppIdGravity

	args[constant.Args0] = ccid[id]
	args[constant.Args1] = city
	args[constant.Args2] = lowPrice
	args[constant.Args3] = highPrice
	res := util.PostTest(constant.UrlQuery, args)

	var array map[string]string
	json.Unmarshal(res, &array)
	ids := make(map[string]string)
	for index, value := range array {
		var info Info
		json.Unmarshal([]byte(value), &info)
		fmt.Println("信息", index, ":")
		fmt.Println("公司：", info.CompanyName)
		fmt.Println("标题：", info.Title)
		fmt.Println("内容：", info.Content)
		fmt.Println("城市：", info.City)
		fmt.Println("价格：", info.Price)
		fmt.Println("时间戳：", info.PublishTime)
		ids[info.CompanyName +index] = info.ID
		fmt.Println("ID缓存key：", info.CompanyName +index)
		fmt.Println()

	}
	util.SaveId(ids, constant.PostIdFile)
}
