package main

import (
	"constant"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strconv"
	"time"
	"util"
	"wallet"
)

type Order struct {
	Id string
}

/**
	用户下单
 */
func (order *Order) PlaceOrder(wlt wallet.Wallet) {
	readId := util.ReadId(constant.PostIdFile)
	fmt.Println("id:",readId[order.Id])
	args := Sign(wlt, readId[order.Id])
	//args := Sign(wlt, order.Id)
	args[constant.ChainCodeName] = constant.ChainCodeTrade
	args[constant.Version] = constant.TradeVersion
	args[constant.Function] = constant.Submit
	args[constant.AppId] = constant.AppIdGravity
	util.PostTest(constant.UrlInvoke,args)
}

/**
	用户完成订单
 */
func (order *Order) FinishOrder(wlt wallet.Wallet) {
	readId := util.ReadId(constant.OrderIdFile)
	args := Sign(wlt, readId[order.Id])
	args[constant.ChainCodeName] = constant.ChainCodeTrade
	args[constant.Version] = constant.TradeVersion
	args[constant.Function] = constant.Finish
	args[constant.AppId] = constant.AppIdGravity
	util.PostTest(constant.UrlInvoke,args)
}

/**
	商家确认订单
 */
func (order *Order) ConfirmOrder(wlt wallet.Wallet) {
	readId := util.ReadId(constant.OrderIdFile)
	args := Sign(wlt, readId[order.Id])
	args[constant.ChainCodeName] = constant.ChainCodeTrade
	args[constant.Version] = constant.TradeVersion
	args[constant.Function] = constant.Confirm
	args[constant.AppId] = constant.AppIdGravity
	util.PostTest(constant.UrlInvoke,args)
}

/**
	C获取订单列表
 */
func (order *Order) GetCustomOrder(wlt *wallet.Wallet) {
	args := make(map[string]string)
	args[constant.ChainCodeName] = constant.ChainCodeTrade
	args[constant.Version] = constant.TradeVersion
	args[constant.Function] = constant.GetTradeByConstumer
	args[constant.AppId] = constant.AppIdGravity

	args[constant.Args0] = base64.StdEncoding.EncodeToString(wlt.PublicKey)
	res := util.PostTest(constant.UrlQuery, args)
	order.analysisOrder(res,true)
}

/**
	B获取订单列表
 */
func (order *Order) GetBusinessOrder(pubkey string) {
	args := make(map[string]string)
	args[constant.ChainCodeName] = constant.ChainCodeTrade
	args[constant.Version] = constant.TradeVersion
	args[constant.Function] = constant.GetTradeByBusiness
	args[constant.AppId] = constant.AppIdGravity

	args[constant.Args0] = pubkey
	res := util.PostTest(constant.UrlQuery, args)
	order.analysisOrder(res,false)

}

func (order *Order) analysisOrder(res []byte,isCus bool) {
	var state string

	var array map[string]string
	json.Unmarshal(res, &array)
	ids := make(map[string]string)
	count := 1
	for id, str := range array {
		var trade Trade
		json.Unmarshal([]byte(str),&trade)
		fmt.Println("订单号:",id)
		//decode := util.AesAndBase64Decode(trade.Customer, wlt.AesKey)
		fmt.Println("用户:",trade.Customer)
		fmt.Println("商家:",trade.Business)
		fmt.Println("信息ID:",trade.InfoID)
		fmt.Println("标题:",trade.Title)
		fmt.Println("提单时间:",trade.SubmitTime)
		if trade.State >1{
			fmt.Println("确认时间:",trade.ConfirmTime)
		}
		if trade.State > 2 {
			fmt.Println("完成时间:",trade.FinishTIme)
		}
		switch trade.State {
		case 1:
			state = "用户已下单"
		case 2:
			state = "商家已确认"
		case 3:
			state = "订单完成"
		}

		fmt.Println("状态:",state)
		if trade.State < 3 {
			if isCus {
				ids[trade.Customer+strconv.Itoa(count)] = id
			}else{
				ids[trade.Business+strconv.Itoa(count)] = id
				fmt.Println(trade.Business+strconv.Itoa(count))
			}
			count++
			fmt.Println("ID缓存key：",trade.Customer+strconv.Itoa(count))
		}
		fmt.Println()
	}
	util.SaveId(ids,constant.OrderIdFile)
	fmt.Println("dat:",ids)

}

type Trade struct {
	Customer   string `json:"Constumer"`
	Business    string`json:"Business"`
	InfoID      string`json:"InfoID"`
	Title       string`json:"Title"`
	SubmitTime  time.Time`json:"SubmitTime"`
	ConfirmTime time.Time`json:"ConfirmTime"`
	FinishTIme  time.Time`json:"FinishTIme"`
	State       int`json:"State"`
}



