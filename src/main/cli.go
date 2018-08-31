package main

import (
	"os"
	"fmt"
	"flag"
	"util"
	"wallet"
	"encoding/base64"
)

const (
	createWallet      = "createwallet"
	getWallets        = "getwallets"
	name              = "name"
	registerCustom    = "registerc"
	registerBussiness = "registerb"
	getUser = "getuser"
	walletName        = "w"
	id                = "id"
	nickname          = "nn"
	age               = "age"
	tel               = "tel"
	bussinessName     = "bn"
	bussinessId       = "bid"
	createPost        = "createpost"
	title             = "title"
	content           = "content"
	city              = "city"
	price             = "price"
	getChainCodeList  = "getchaincode"
	getBusinessOrder  = "getbusorder"
	getCustomerOrder  = "getcusorder"
	getBusinessPost   = "getbuspost"
	placeOrder   = "placeorder"
	confirmOrder   = "confirmorder"
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
	regBBussinessName := registerBusCmd.String(bussinessName, "", "在-bn后输入公司名称")

	//获取用户信息
	getUserCmd := flag.NewFlagSet(getUser, flag.ExitOnError)
	getUserWallet := getUserCmd.String(name, "", "在-w后输入钱包名")

	//创建帖子
	postCmd := flag.NewFlagSet(createPost, flag.ExitOnError)
	postTitle := postCmd.String(title, "", "在-title后输入标题")
	postContent := postCmd.String(content, "", "在-content后输入帖子内容")
	postCity := postCmd.String(city, "", "在-city后输入城市")
	postPrice := postCmd.String(price, "", "在-price后输入价格")
	postWallet := postCmd.String(walletName, "", "在-w后输入钱包名称")

	//获取链码列表
	getChainCodeCmd := flag.NewFlagSet(getChainCodeList, flag.ExitOnError)

	//C获取订单列表
	getCusOrderCmd := flag.NewFlagSet(getCustomerOrder, flag.ExitOnError)
	getCusOrderWallet := getCusOrderCmd.String(walletName, "", "在-w后输入钱包名称")

	//B获取订单列表
	getBusOrderCmd := flag.NewFlagSet(getBusinessOrder, flag.ExitOnError)
	getBusOrderWallet := getBusOrderCmd.String(walletName, "", "在-w后输入钱包名称")

	//B获取帖子列表
	getBusPostCmd := flag.NewFlagSet(getBusinessPost, flag.ExitOnError)
	getBusPostWallet := getBusPostCmd.String(walletName, "", "在-w后输入钱包名称")

	//C下单
	placeOrderCmd := flag.NewFlagSet(placeOrder, flag.ExitOnError)
	placeOrderWallet := placeOrderCmd.String(walletName, "", "在-w后输入钱包名称")
	placeOrderId := placeOrderCmd.String(id, "", "在-id后输入帖子ID")

	//B确认订单
	confirmOrderCmd := flag.NewFlagSet(confirmOrder, flag.ExitOnError)
	confirmOrderWallet := confirmOrderCmd.String(walletName, "", "在-w后输入钱包名称")
	confirmOrderId := confirmOrderCmd.String(id, "", "在-id后输入订单ID")

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
	case createPost:
		err = postCmd.Parse(os.Args[2:])
	case getChainCodeList:
		err = getChainCodeCmd.Parse(os.Args[2:])
	case getCustomerOrder:
		err = getCusOrderCmd.Parse(os.Args[2:])
	case getBusinessOrder:
		err = getBusOrderCmd.Parse(os.Args[2:])
	case getBusinessPost:
		err = getBusPostCmd.Parse(os.Args[2:])
	case placeOrder:
		err = placeOrderCmd.Parse(os.Args[2:])
	case confirmOrder:
		err = confirmOrderCmd.Parse(os.Args[2:])
	case getUser:
		err = getUserCmd.Parse(os.Args[2:])
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

	if getUserCmd.Parsed()  {
		if *getUserWallet != ""{
			wlt, e := wallet.ExamWallet(*getUserWallet)
			if e != nil {
				reg := &Register{}
				reg.GetUserInfo(wlt)
			}
		}else{
			registerCusCmd.Usage()
			os.Exit(1)
		}
	}


	if postCmd.Parsed() {
		if *postCity  != "" && *postContent  != "" && *postPrice  != "" && *postTitle  != "" && *postWallet != ""{
			post := &Post{*postTitle, *postContent, *postCity, *postPrice}
			fmt.Println(*post)
			cli.createPost(*postWallet,post)
		}
	}
	if getChainCodeCmd.Parsed() {
		getInfo := &GetInfo{}
		getInfo.GetChainCodeList()
	}

	if getCusOrderCmd.Parsed() {
		if *getCusOrderWallet != "" {
			wlt := ExamWallet(cli, *getCusOrderWallet)
			if wlt != nil {
				getInfo := &GetInfo{base64.StdEncoding.EncodeToString(wlt.PublicKey)}
				getInfo.GetCustomOrder()
			}
		}else{
			getCusOrderCmd.Usage()
			os.Exit(1)
		}
	}

	if getBusOrderCmd.Parsed() {
		if *getBusOrderWallet != "" {
			wlt := ExamWallet(cli, *getBusOrderWallet)
			if wlt != nil {
				getInfo := &GetInfo{base64.StdEncoding.EncodeToString(wlt.PublicKey)}
				getInfo.GetBusinessOrder()
			}
		}else{
			getBusOrderCmd.Usage()
			os.Exit(1)
		}
	}

	if getBusPostCmd.Parsed() {
		if *getBusPostWallet != "" {
			wlt := ExamWallet(cli, *getBusPostWallet)
			if wlt != nil {
				getInfo := &GetInfo{base64.StdEncoding.EncodeToString(wlt.PublicKey)}
				getInfo.GetBusinessPost()
			}
		}else{
			getBusPostCmd.Usage()
			os.Exit(1)
		}
	}

	if placeOrderCmd.Parsed() {
		if *placeOrderWallet != "" && *placeOrderId != ""{
			wlt := ExamWallet(cli, *placeOrderWallet)
			if wlt != nil {
				order := &Order{*placeOrderId}
				order.PlaceOrder(*wlt)
			}

		} else {
			placeOrderCmd.Usage()
			os.Exit(1)
		}
	}

	if confirmOrderCmd.Parsed() {
		if *confirmOrderId != "" && *confirmOrderWallet != ""{
			wlt := ExamWallet(cli, *confirmOrderWallet)
			if wlt != nil {
				order := &Order{*confirmOrderId}
				order.ConfirmOrder(*wlt)
			}

		} else {
			confirmOrderCmd.Usage()
			os.Exit(1)
		}
	}
}

/**
	读取钱包信息
 */
func ExamWallet(cli *Cli, wltName string) *wallet.Wallet{
	wlt, e := wallet.ExamWallet(wltName)
	if e != nil {
		util.LogE(e)
		cli.printUsage()
		return nil
	} else {
		return &wlt
	}
}

type exist func()


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
	fmt.Println("-----------------------------------钱包---------------------------------------------------------")

	fmt.Println("    （创建钱包）createwallet -name mike ")
	fmt.Println("    （打印全部钱包名称） getwallets")
	fmt.Println("-----------------------------------注册/修改-------------------------------------------------------")
	fmt.Println("    （用户注册）registerc -w mike -name mike -nn mk -age 18 -tel 13812345678 -id 110101199001010000")
	fmt.Println("    （商家注册）registerb -w mike -name mike -nn mk -age 18 -tel 13812345678 -id 110101199001010000 -bid 50001000-3 -bn 北京城市网邻信息技术有限公司")
	fmt.Println("    （查询用户信息）getuser -w mike)")
	fmt.Println("-----------------------------------信息-------------------------------------------------------")
	fmt.Println("    （发布信息）createpost -w mike -title 哥俩好搬家公司 -content 负责朝阳区搬家业务 -price 200 -city 北京")
	fmt.Println("    （获取合约列表）getchaincode")
	fmt.Println("    （用户获取订单列表）getcusorder -w mike")
	fmt.Println("    （商家获取订单列表）getbusorder -w mike")
	fmt.Println("    （商家获取已发布信息）getbuspost -w mike")
	fmt.Println("    （用户提单）placeorder -w mike -id postId")
	fmt.Println("    （商家确认订单）confirmorder -w mike -id orderId")
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

/**
	创建帖子
 */
func (cli *Cli) createPost(walletName string, post *Post) {
	wlt, e := wallet.ExamWallet(walletName)
	if e != nil {
		util.LogE(e)
		cli.printUsage()
	}else{
		post.PostCommit(wlt)
	}

}



