package main

import (
	"wallet"
	"encoding/json"
	"fmt"
	"util"
	"constant"
)

/**
	帖子信息
 */
type Post struct {
	Title       string `json:"Title"`
	Content     string `json:"Content"`
	City        string `json:"City"`
	Price       string `json:"Price"`
}

func (post *Post) PostCommit(wlt wallet.Wallet) {
	//生成json
	jsonByte, e := json.Marshal(post)
	if e != nil {
		fmt.Println(e)
	}

	args := Sign(wlt, string(jsonByte))
	util.PostTest(constant.BaseUrl,args)
}
