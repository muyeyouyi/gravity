package main

import (
	"wallet"
	"util"
	"constant"
)

type Order struct {
	Id string
}

func (order *Order) PlaceOrder(wlt wallet.Wallet) {
	args := Sign(wlt, order.Id)
	args["accessToken"] = AccessToken
	args["chaincodeName"] = ""
	args["version"] = "1"
	util.PostTest(constant.BaseUrl,args)
}

func (order *Order) ConfirmOrder(wlt wallet.Wallet) {
	args := Sign(wlt, order.Id)
	args["accessToken"] = AccessToken
	args["chaincodeName"] = ""
	args["version"] = "1"
	util.PostTest(constant.BaseUrl,args)
}


