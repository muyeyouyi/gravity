package main
//
//import (
//	"wallet"
//	"util"
//	"fmt"
//	"encoding/json"
//	"constant"
//	"log"
//)
////type Response struct {
////	Status int `json:"status"`
////	Message string `json:"message"`
////	BlockNumber string `json:"blockNumber"`
////	Data []string `json:"data"`
////}
//
//type Register struct {
//	Nickname    string `json:"Nickname"`
//	Name        string `json:"Name"`
//	Age         string `json:"Age"`
//	Phonenum    string `json:"Phonenum"`
//	ID          string `json:"ID"`
//	CompanyID   string `json:"CompanyID"`
//	CompanyName string `json:"CompanyName"`
//}
//
///**
//	注册
// */
//func (user *Register) RegisterCommit(wlt wallet.Wallet) {
//	//字段加密
//	user.Name = util.AesAndBase58Encode(user.Name, wlt.aesKey)
//	user.Phonenum = util.AesAndBase58Encode(user.Phonenum, wlt.aesKey)
//	user.ID = util.AesAndBase58Encode(user.ID, wlt.aesKey)
//	user.Age = util.AesAndBase58Encode(user.Age, wlt.aesKey)
//	if user.CompanyID != "" {
//		user.CompanyID = util.AesAndBase58Encode(user.CompanyID, wlt.aesKey)
//	}
//	fmt.Println("加密：", *user)
//
//	//生成json
//	jsonByte, e := json.Marshal(user)
//	//fmt.Println(string(jsonByte))
//	if e != nil {
//		fmt.Println(e)
//	}
//
//	//签名
//	args := Sign(wlt, string(jsonByte))
//	args[constant.ChainCodeName] = constant.ChainCodeUser
//	args[constant.Version] = constant.UserVersion
//	args[constant.Function] = constant.Set
//	args[constant.AppId] = constant.AppIdGravity
//
//	//网络请求
//	util.PostTest(constant.UrlInvoke, args)
//
//}
//
///**
//	获取用户信息
// */
//func (user *Register) GetUserInfo(wlt wallet.Wallet) {
//	args := make(map[string]string)
//	args[constant.ChainCodeName] = constant.ChainCodeUser
//	args[constant.Version] = constant.UserVersion
//	args[constant.Function] = constant.Get
//	args[constant.AppId] = constant.AppIdGravity
//
//	args[constant.Args0] = util.Base58(wlt.publicKey)
//	res := util.PostTest(constant.UrlQuery, args)
//
//	user.analysis(res, wlt)
//
//}
//
//func (user *Register) analysis(res []byte, wlt wallet.Wallet) {
//	defer func() {
//		if e := recover(); e!= nil{
//			log.Fatalln("error:json解析异常")
//		}
//	}()
//	//解析
//	var userInfo Register
//	json.Unmarshal(res, &userInfo)
//	userInfo.ID = util.AesAndBase58Decode(userInfo.ID, wlt.aesKey)
//	userInfo.Name = util.AesAndBase58Decode(userInfo.Name, wlt.aesKey)
//	userInfo.Age = util.AesAndBase58Decode(userInfo.Age, wlt.aesKey)
//	userInfo.Phonenum = util.AesAndBase58Decode(userInfo.Phonenum, wlt.aesKey)
//	if userInfo.CompanyID != "" {
//		userInfo.CompanyID = util.AesAndBase58Decode(userInfo.CompanyID, wlt.aesKey)
//	}
//	fmt.Println("用户信息：")
//	fmt.Println("昵称：", userInfo.Nickname)
//	fmt.Println("姓名：", userInfo.Name)
//	fmt.Println("年龄：", userInfo.Age)
//	fmt.Println("电话：", userInfo.Phonenum)
//	fmt.Println("身份证号：", userInfo.ID)
//	if userInfo.CompanyName != "" {
//		fmt.Println("企业名称：", userInfo.CompanyName)
//		fmt.Println("法人登记证号：", userInfo.CompanyID)
//	}
//}
