package main

import (
	"constant"
	"util"
	"wallet"
	"encoding/json"
	"fmt"
)

/**
	帖子信息
 */
type Post struct {
	Title       string `json:"Title"`
	Content     string `json:"Content"`
	CompanyName string `json:"CompanyName"`
	City        string `json:"City"`
	Price       int    `json:"Price"`
}

/**
	帖子信息
 */
func (post *Post) PostCommit(wlt wallet.Wallet) {
	//生成json
	jsonByte, e := json.Marshal(post)
	if e != nil {
		fmt.Println(e)
	}

	args := Sign(wlt, string(jsonByte))
	args[constant.Function] = constant.Set
	args[constant.Version] = constant.InfoVersion
	args[constant.ChainCodeName] = constant.ChainCodeInfo
	args[constant.AppId] = constant.AppIdGravity
	util.PostTest(constant.UrlInvoke, args)
}

/**
	查询一个商家的所有帖子
 */
func (post *Post) GetPosts(pubkey string) {
	args := make(map[string]string)
	args[constant.ChainCodeName] = constant.ChainCodeInfo
	args[constant.Version] = constant.InfoVersion
	args[constant.Function] = constant.GetByOwner
	args[constant.AppId] = constant.AppIdGravity
	args[constant.AccessToken] = AccessToken

	args[constant.Args0] = pubkey
	util.PostTest(constant.UrlQuery, args)
}

/**
	查询一个帖子详情
 */
func (post *Post) GetPostDetail(ID string) {
	args := make(map[string]string)
	args[constant.ChainCodeName] = constant.ChainCodeInfo
	args[constant.Version] = constant.InfoVersion
	args[constant.Function] = constant.Get
	args[constant.AppId] = constant.AppIdGravity
	args[constant.AccessToken] = AccessToken

	args[constant.Args0] = ID
	util.PostTest(constant.UrlQuery, args)
}


