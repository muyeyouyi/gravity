package main

import (
	"wallet"
	"encoding/json"
	"fmt"
	"constant"
	"util"
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

func (post *Post) PostCommit(wlt wallet.Wallet) {
	//生成json
	jsonByte, e := json.Marshal(post)
	if e != nil {
		fmt.Println(e)
	}

	args := Sign(wlt, string(jsonByte))
	args[constant.Args0] = constant.Set
	args[constant.Function] = constant.Invoke
	args[constant.ChainCodeName] = constant.Info
	util.PostTest(constant.BaseUrl, args)
}

func (post *Post) GetPosts(pubkey string) {
	args := make(map[string]string)
	args[constant.Function] = constant.Query
	args[constant.ChainCodeName] = constant.Info
	args[constant.Args0] = constant.GetByOwner
	args[constant.Args1] = pubkey
	util.PostTest(constant.BaseUrl, args)
}
