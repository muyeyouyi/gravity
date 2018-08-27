package main

import (
	"os"
	"fmt"
	"flag"
	"util"
	"wallet"
)

const (
	createWallet      = "createwallet"
	getWallets        = "getwallets"
	name              = "name"
	registerCustom    = "registerc"
	registerBussiness = "registerb"
	walletName            = "w"
	id                = "id"
	nickname          = "nn"
	age               = "age"
	tel               = "tel"
	bussinessName     = "bn"
	bussinessId       = "bid"
)

type Cli struct {
}

func (cli *Cli) Run() {
	//输出提示信息
	cli.validateArgs()

	//cmd创建一个钱包
	createWalletCmd := flag.NewFlagSet(createWallet, flag.ExitOnError)
	createWalletData := createWalletCmd.String(name, "", "在-name后输入名字")

	//打印钱包地址
	getWalletsCmd := flag.NewFlagSet(getWallets, flag.ExitOnError)

	//C注册基本信息
	registerCusCmd := flag.NewFlagSet(registerCustom, flag.ExitOnError)
	regName := registerCusCmd.String(name, "", "在-name后输入名字")
	regID := registerCusCmd.String(id, "", "在-id后输入身份证号")
	regTel := registerCusCmd.String(tel, "", "在-tel后输入手机号")
	regAge := registerCusCmd.String(age, "", "在-age后输入年龄")
	regNickName := registerCusCmd.String(nickname, "", "在-nn后输入昵称")
	regWallet := registerCusCmd.String(walletName, "", "在-w后输入钱包名")

	//B注册基本信息
	registerBusCmd := flag.NewFlagSet(registerBussiness, flag.ExitOnError)
	regBName := registerBusCmd.String(name, "", "在-name后输入名字")
	regBID := registerBusCmd.String(id, "", "在-id后输入身份证号")
	regBTel := registerBusCmd.String(tel, "", "在-tel后输入手机号")
	regBAge := registerBusCmd.String(age, "", "在-age后输入年龄")
	regBNickName := registerBusCmd.String(nickname, "", "在-nn后输入昵称")
	regBWallet := registerBusCmd.String(walletName, "", "在-w后输入钱包名")
	regBBussinessId := registerBusCmd.String(bussinessId, "", "在-bid后输入组织机构代码")
	regBBussinessName := registerBusCmd.String(bussinessName, "", "在-w后输入公司名称")

	//截取命令行内容
	var err error
	switch os.Args[1] {
	case createWallet:
		err = createWalletCmd.Parse(os.Args[2:])
	case getWallets:
		err = getWalletsCmd.Parse(os.Args[2:])
	case registerCustom:
		err = registerCusCmd.Parse(os.Args[2:])
	case registerBussiness:
		err = registerBusCmd.Parse(os.Args[2:])
	}

	util.LogE(err)

	if createWalletCmd.Parsed() {
		if *createWalletData == "" {
			createWalletCmd.Usage()
			os.Exit(1)
		}
		cli.createWallet(*createWalletData)
	}

	if getWalletsCmd.Parsed() {
		cli.printWallets()
	}

	if registerCusCmd.Parsed()  {
		if *regWallet != "" && *regName != "" && *regNickName != ""  && *regID != "" && *regAge != "" && *regTel != ""{
			userInfo := &Register{*regNickName, *regName, *regAge, *regTel, *regID, "", ""}
			fmt.Println(*userInfo)
			cli.register(*regWallet,userInfo)
		}else{
			registerCusCmd.Usage()
			os.Exit(1)
		}
	}

	if registerBusCmd.Parsed()  {
		if *regBWallet != "" && *regBName != "" && *regBNickName != ""  && *regBID != "" && *regBAge != "" && *regBTel != "" && *regBBussinessName != "" && *regBBussinessId != ""{
			userInfo := &Register{*regBNickName, *regBName, *regBAge, *regBTel, *regBID, *regBBussinessId, *regBBussinessName}
			fmt.Println(*userInfo)
			cli.register(*regBWallet,userInfo)
		}else{
			registerCusCmd.Usage()
			os.Exit(1)
		}
	}

}

/**
	验证命令行参数
 */
func (cli *Cli) validateArgs() {
	if len(os.Args) < 2 {
		cli.printUsage()
		os.Exit(1)
	}
}

/**
	打印提示信息
 */
func (cli *Cli) printUsage() {
	fmt.Println("用法：")
	fmt.Println("    ", createWallet, " -", name, " mike （创建一个钱包）")
	fmt.Println("    ", getWallets, "                 (打印全部钱包名称)")
	fmt.Println("    registerc -w mike -name mike -nn mk -age 18 -tel 13812345678 -id 110101199001010000")
	fmt.Println("    registerb -w mike -name mike -nn mk -age 18 -tel 13812345678 -id 110101199001010000 -bid 50001000-3 -bn 北京城市网邻信息技术有限公司")
}

/**
	创建钱包
 */
func (cli *Cli) createWallet(name string) {
	wallets, e := wallet.NewWallets()
	util.LogE(e)
	wallets.CreateWallet(name)
	wallets.SaveToFile()
}
func (cli *Cli) printWallets() {
	ws, e := wallet.NewWallets()
	util.LogE(e)
	names := ws.GetWalletNames()
	for _, name := range names {
		println("钱包名:" + name)
	}
}

/**
	用户注册
 */
func (cli *Cli) register(walletName string, user *Register) {
	wlt, e := wallet.ExamWallet(walletName)
	if e != nil {
		util.LogE(e)
		cli.printUsage()
	}else{
		user.RegisterCommit(wlt)
	}

}



