package main

import (
	"constant"
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

	util.PostTest(constant.UrlQuery, args)
}

/**
	执行匹配
 */
func (match *Match) Match(city, id, lowPrice, highPrice string) {
	args := make(map[string]string)
	args[constant.ChainCodeName] = constant.ChainCodeInfo
	args[constant.Version] = constant.InfoVersion
	args[constant.Function] = constant.Matching
	args[constant.AppId] = constant.AppIdGravity

	args[constant.Args0] = id
	args[constant.Args1] = city
	args[constant.Args2] = lowPrice
	args[constant.Args3] = highPrice
	util.PostTest(constant.UrlQuery, args)
}
