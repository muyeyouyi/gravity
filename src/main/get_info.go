package main

import (
	"util"
	"constant"
)

type GetInfo struct {
	Pubkey string
}

func (getInfo *GetInfo) GetChainCodeList()  {
	args := make(map[string]string)
	util.PostTest(constant.BaseUrl,args)
}
func (info *GetInfo) GetCustomOrder() {

}
func (info *GetInfo) GetBusinessOrder() {

}
