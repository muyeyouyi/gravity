package main

import (
	"wallet"
	"util"
	"fmt"
	"encoding/json"
)

type Register struct {
	Nickname string `json:"Nickname"`
	Name     string `json:"Name"`
	Age      string `json:"Age"`
	Phonenum string `json:"Phonenum"`
	ID       string `json:"ID"`
	CompanyID   string `json:"CompanyID"`
	CompanyName string `json:"CompanyName"`
}

/**
	注册
 */
func (user *Register)RegisterCommit(wlt wallet.Wallet) {
	user.Name = util.AesAndBase64Encode(user.Name,wlt.AesKey)
	user.Phonenum = util.AesAndBase64Encode(user.Phonenum,wlt.AesKey)
	user.ID = util.AesAndBase64Encode(user.ID,wlt.AesKey)
	user.Age = util.AesAndBase64Encode(user.Age,wlt.AesKey)
	if user.CompanyID != "" {
		user.CompanyID = util.AesAndBase64Encode(user.CompanyID,wlt.AesKey)
	}
	fmt.Println("加密：",*user)
	jsonByte, e := json.Marshal(user)
	//fmt.Println(string(jsonByte))
	if e != nil {
		fmt.Println(e)
	}
	args := Sign(wlt, string(jsonByte))


	for key, value := range args {
		fmt.Println("key:",key,"  value:",value)
	}
	verify := util.Verify(args["args0"], args["args1"], args["args2"])
	fmt.Println("veryfy:",verify)

	//user.Name = util.AesAndBase64Decode(user.Name,wlt.AesKey)
	//user.Phonenum = util.AesAndBase64Decode(user.Phonenum,wlt.AesKey)
	//user.ID = util.AesAndBase64Decode(user.ID,wlt.AesKey)
	//user.Age = util.AesAndBase64Decode(user.Age,wlt.AesKey)
	//if user.CompanyID != "" {
	//	user.CompanyID = util.AesAndBase64Decode(user.CompanyID,wlt.AesKey)
	//}
	//fmt.Println("解密:",*user)
	//todo 网络请求


}

