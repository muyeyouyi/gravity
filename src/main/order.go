package main

import (
	"wallet"
	"constant"
	"util"
)

type Order struct {
	Id string
}

/**
	用户下单
 */
func (order *Order) PlaceOrder(wlt wallet.Wallet) {
	args := Sign(wlt, order.Id)
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
	args := Sign(wlt, order.Id)
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
	args := Sign(wlt, order.Id)
	args[constant.ChainCodeName] = constant.ChainCodeTrade
	args[constant.Version] = constant.TradeVersion
	args[constant.Function] = constant.Confirm
	args[constant.AppId] = constant.AppIdGravity
	util.PostTest(constant.UrlInvoke,args)
}

/**
	C获取订单列表
 */
func (order *Order) GetCustomOrder(pubkey string) {
	args := make(map[string]string)
	args[constant.ChainCodeName] = constant.ChainCodeTrade
	args[constant.Version] = constant.TradeVersion
	args[constant.Function] = constant.GetTradeByConstumer
	args[constant.AppId] = constant.AppIdGravity
	args[constant.AccessToken] = AccessToken

	args[constant.Args0] = pubkey
	util.PostTest(constant.UrlQuery,args)
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
	args[constant.AccessToken] = AccessToken

	args[constant.Args0] = pubkey
	util.PostTest(constant.UrlQuery,args)

}


