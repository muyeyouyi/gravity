package main

import (
	"fmt"
	"util"
)

func main() {
	wallet := util.NewWallet()
	sign := wallet.Sign("1")
	fmt.Println("base58编码:",sign)
	verify := wallet.Verify(wallet.GetPubKey(), "1", sign)
	fmt.Println("签名验证:",verify)
	//cli()
	//newWallet := wallet.NewWallet()
	//fmt.Println("原始钱包:", newWallet)
	//prikey := newWallet.GetPrivateKey()
	//fmt.Println("导出私钥:",prikey)
	//wlt := wallet.ImportWallet(prikey)
	//fmt.Println("导入钱包:",wlt)

}



//func cli() {
//	cli := &Cli{}
//	cli.Run()
//}
