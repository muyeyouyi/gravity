package main

import (
	"util"
	"constant"
)

type GetInfo struct {
	Pubkey string
}

/**
	获取链码列表
 */
func (info *GetInfo) GetChainCodeList()  {
	args := make(map[string]string)
	args["accessToken"] = AccessToken
	args["chaincodeName"] = ""
	args["version"] = "1"
	util.PostTest(constant.BaseUrl,args)
}

/**
	C端拉取订单
 */
func (info *GetInfo) GetCustomOrder() {
	args := make(map[string]string)
	args["args0"] = info.Pubkey
	args["accessToken"] = AccessToken
	args["chaincodeName"] = ""
	args["version"] = "1"

	util.PostTest(constant.BaseUrl,args)
}

/**
	拉取B端
 */
func (info *GetInfo) GetBusinessOrder() {
	args := make(map[string]string)
	args["args0"] = info.Pubkey
	args["accessToken"] = AccessToken
	args["chaincodeName"] = ""
	args["version"] = "1"
	util.PostTest(constant.BaseUrl,args)

}

/**
	拉取B端帖子列表
 */
func (info *GetInfo) GetBusinessPost() {
	args := make(map[string]string)
	args["args0"] = info.Pubkey
	args["accessToken"] = AccessToken
	args["chaincodeName"] = ""
	args["version"] = "1"
	util.PostTest(constant.BaseUrl,args)
}
