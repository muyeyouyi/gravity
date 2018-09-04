package main

import (
	"os"
	"fmt"
	"flag"
	"util"
	"wallet"
	"encoding/base64"
	"strconv"
)

const (
	createWallet      = "createwallet"
	getWallets        = "getwallets"
	name              = "name"
	registerCustom    = "registerc"
	registerBussiness = "registerb"
	getUser           = "getuser"
	walletName        = "w"
	id                = "id"
	nickname          = "nn"
	age               = "age"
	tel               = "tel"
	businessName      = "bn"
	businessId        = "bid"
	createPost        = "createpost"
	title             = "title"
	content           = "content"
	city              = "city"
	price             = "price"
	getChainCodeList  = "getchaincode"
	useChainCode      = "usechaincode"
	getBusinessOrder  = "getbusorder"
	getCustomerOrder  = "getcusorder"
	placeOrder        = "placeorder"
	finishOrder       = "finishorder"
	confirmOrder      = "confirmorder"
	getPost           = "getpost"
	getPostDetail     = "getpostdetail"
	lowPrice          = "lowp"
	highPrice         = "highp"
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
	regBBussinessId := registerBusCmd.String(businessId, "", "在-bid后输入组织机构代码")
	regBBussinessName := registerBusCmd.String(businessName, "", "在-bn后输入公司名称")

	//获取用户信息
	getUserCmd := flag.NewFlagSet(getUser, flag.ExitOnError)
	getUserWallet := getUserCmd.String(walletName, "", "在-w后输入钱包名")

	//创建帖子
	createPostCmd := flag.NewFlagSet(createPost, flag.ExitOnError)
	postTitle := createPostCmd.String(title, "", "在-title后输入标题")
	postContent := createPostCmd.String(content, "", "在-content后输入帖子内容")
	postBusinessName := createPostCmd.String(businessName, "", "在-bn后输入公司名称")
	postCity := createPostCmd.String(city, "", "在-city后输入城市")
	postPrice := createPostCmd.String(price, "", "在-price后输入价格")
	postWallet := createPostCmd.String(walletName, "", "在-w后输入钱包名称")

	//查询商家帖子列表
	getPostsCmd := flag.NewFlagSet(getPost, flag.ExitOnError)
	getPostsWallet := getPostsCmd.String(title, "", "在-w后输入钱包名称")

	//查询帖子详情
	getPostDetailCmd := flag.NewFlagSet(getPostDetail, flag.ExitOnError)
	getPostDetailId := getPostDetailCmd.String(id, "", "在-id后输入信息ID")

	//获取链码列表
	getChainCodeCmd := flag.NewFlagSet(getChainCodeList, flag.ExitOnError)

	//使用链码匹配
	useChainCodeCmd := flag.NewFlagSet(useChainCode, flag.ExitOnError)
	useChainCodeId := useChainCodeCmd.String(id, "", "在-w后输入钱包名称")
	useChainCodeCity := useChainCodeCmd.String(city, "", "在-w后输入钱包名称")
	useChainCodeLowPrice := useChainCodeCmd.String(lowPrice, "", "在-lowp后输入最低价格")
	useChainCodeHighPrice := useChainCodeCmd.String(highPrice, "", "在-highp后输入最高价格")

	//C获取订单列表
	getCusOrderCmd := flag.NewFlagSet(getCustomerOrder, flag.ExitOnError)
	getCusOrderWallet := getCusOrderCmd.String(walletName, "", "在-w后输入钱包名称")

	//B获取订单列表
	getBusOrderCmd := flag.NewFlagSet(getBusinessOrder, flag.ExitOnError)
	getBusOrderWallet := getBusOrderCmd.String(walletName, "", "在-w后输入钱包名称")

	//C下单
	placeOrderCmd := flag.NewFlagSet(placeOrder, flag.ExitOnError)
	placeOrderWallet := placeOrderCmd.String(walletName, "", "在-w后输入钱包名称")
	placeOrderId := placeOrderCmd.String(id, "", "在-id后输入帖子ID")

	//B确认订单
	confirmOrderCmd := flag.NewFlagSet(confirmOrder, flag.ExitOnError)
	confirmOrderWallet := confirmOrderCmd.String(walletName, "", "在-w后输入钱包名称")
	confirmOrderId := confirmOrderCmd.String(id, "", "在-id后输入订单ID")

	//C下单
	finishOrderCmd := flag.NewFlagSet(finishOrder, flag.ExitOnError)
	finishOrderWallet := finishOrderCmd.String(walletName, "", "在-w后输入钱包名称")
	finishOrderId := finishOrderCmd.String(id, "", "在-id后输入订单ID")

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
		err = createPostCmd.Parse(os.Args[2:])
	case getChainCodeList:
		err = getChainCodeCmd.Parse(os.Args[2:])
	case getCustomerOrder:
		err = getCusOrderCmd.Parse(os.Args[2:])
	case getBusinessOrder:
		err = getBusOrderCmd.Parse(os.Args[2:])
	case placeOrder:
		err = placeOrderCmd.Parse(os.Args[2:])
	case confirmOrder:
		err = confirmOrderCmd.Parse(os.Args[2:])
	case getUser:
		err = getUserCmd.Parse(os.Args[2:])
	case getPost:
		err = getPostsCmd.Parse(os.Args[2:])
	case getPostDetail:
		err = getPostDetailCmd.Parse(os.Args[2:])
	case useChainCode:
		err = useChainCodeCmd.Parse(os.Args[2:])
	case finishOrder:
		err = finishOrderCmd.Parse(os.Args[2:])
	}

	util.LogE(err)
	//------------------------------------------------钱包------------------------------------------------------------

	/**
		创建钱包
	 */
	if createWalletCmd.Parsed() {
		if *createWalletData == "" {
			createWalletCmd.Usage()
			os.Exit(1)
		}
		cli.createWallet(*createWalletData)
	}

	/**
		打印钱包
	 */
	if getWalletsCmd.Parsed() {
		cli.printWallets()
	}
	//------------------------------------------------用户资料------------------------------------------------------------

	/**
		用户注册
	 */
	if registerCusCmd.Parsed() {
		if *regWallet != "" && *regName != "" && *regNickName != "" && *regID != "" && *regAge != "" && *regTel != "" {
			userInfo := &Register{*regNickName, *regName, *regAge, *regTel, *regID, "", ""}
			fmt.Println(*userInfo)
			cli.register(*regWallet, userInfo)
		} else {
			registerCusCmd.Usage()
			os.Exit(1)
		}
	}
	/**
		商家注册
	 */
	if registerBusCmd.Parsed() {
		if *regBWallet != "" && *regBName != "" && *regBNickName != "" && *regBID != "" && *regBAge != "" && *regBTel != "" && *regBBussinessName != "" && *regBBussinessId != "" {
			userInfo := &Register{*regBNickName, *regBName, *regBAge, *regBTel, *regBID, *regBBussinessId, *regBBussinessName}
			fmt.Println(*userInfo)
			cli.register(*regBWallet, userInfo)
		} else {
			registerCusCmd.Usage()
			os.Exit(1)
		}
	}

	/**
		获取用户信息
	 */
	if getUserCmd.Parsed() {
		if *getUserWallet != "" {
			wlt, e := wallet.ExamWallet(*getUserWallet)
			if e != nil {
				util.LogE(e)
			}else{
				reg := &Register{}
				reg.GetUserInfo(wlt)
			}
		} else {
			getUserCmd.Usage()
			os.Exit(1)
		}
	}
	//------------------------------------------------合约------------------------------------------------------------

	/**
		获取合约列表
	 */
	if getChainCodeCmd.Parsed() {
		match := &Match{}
		match.GetMatchList()
	}

	/**
		匹配
	 */
	if useChainCodeCmd.Parsed() {
		if *useChainCodeCity != "" && *useChainCodeId != "" && *useChainCodeLowPrice != "" && *useChainCodeHighPrice != "" {
			match := &Match{}
			match.Match(*useChainCodeCity, *useChainCodeId, *useChainCodeLowPrice, *useChainCodeHighPrice)
		} else {
			useChainCodeCmd.Usage()
			os.Exit(1)
		}
	}
	//------------------------------------------------帖子------------------------------------------------------------

	/**
		商家发布信息
	 */
	if createPostCmd.Parsed() {
		if *postCity != "" && *postContent != "" && *postPrice != "" && *postTitle != "" && *postWallet != "" && *postBusinessName != "" {
			price, e := strconv.Atoi(*postPrice)
			if e != nil {
				fmt.Println(e)
			} else {
				post := &Post{*postTitle, *postContent, *postBusinessName, *postCity, price}
				fmt.Println(*post)
				cli.createPost(*postWallet, post)
			}
		}
	}

	/**
		获取一个商家所有帖子列表
	 */
	if getPostsCmd.Parsed() {
		if *getPostsWallet != "" {
			post := &Post{}
			post.GetPosts(*getPostsWallet)
		}
	}

	/**
		查询帖子详情
	 */
	if getPostDetailCmd.Parsed() {
		if *getPostDetailId != "" {
			post := &Post{}
			post.GetPostDetail(*getPostDetailId)
		}
	}

	//------------------------------------------------订单------------------------------------------------------------

	/*
		用户获取订单列表
	 */
	if getCusOrderCmd.Parsed() {
		if *getCusOrderWallet != "" {
			wlt := ExamWallet(cli, *getCusOrderWallet)
			if wlt != nil {
				getInfo := &Order{}
				getInfo.GetCustomOrder(base64.StdEncoding.EncodeToString(wlt.PublicKey))
			}
		} else {
			getCusOrderCmd.Usage()
			os.Exit(1)
		}
	}

	/**
		商家获取订单列表
	 */
	if getBusOrderCmd.Parsed() {
		if *getBusOrderWallet != "" {
			wlt := ExamWallet(cli, *getBusOrderWallet)
			if wlt != nil {
				getInfo := &Order{}
				getInfo.GetBusinessOrder(base64.StdEncoding.EncodeToString(wlt.PublicKey))
			}
		} else {
			getBusOrderCmd.Usage()
			os.Exit(1)
		}
	}

	/**
		用户下单
	 */
	if placeOrderCmd.Parsed() {
		if *placeOrderWallet != "" && *placeOrderId != "" {
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

	/**
		商家确认订单
	 */
	if confirmOrderCmd.Parsed() {
		if *confirmOrderId != "" && *confirmOrderWallet != "" {
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

	/**
		用户完成订单
	 */
	if finishOrderCmd.Parsed() {
		if *finishOrderWallet != "" && *finishOrderId != "" {
			wlt := ExamWallet(cli, *finishOrderWallet)
			if wlt != nil {
				order := &Order{*finishOrderId}
				order.FinishOrder(*wlt)
			}

		} else {
			finishOrderCmd.Usage()
			os.Exit(1)
		}
	}
}

/**
	读取钱包信息
 */
func ExamWallet(cli *Cli, wltName string) *wallet.Wallet {
	wlt, e := wallet.ExamWallet(wltName)
	if e != nil {
		util.LogE(e)
		cli.printUsage()
		return nil
	} else {
		return &wlt
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
	fmt.Println("-----------------------------------钱包---------------------------------------------------------")
	fmt.Println("    （创建钱包）createwallet -name mike ")
	fmt.Println("    （打印全部钱包名称） getwallets")
	fmt.Println("---------------------------------用户资料-------------------------------------------------------")
	fmt.Println("    （用户注册）registerc -w mike -name mike -nn mk -age 18 -tel 13812345678 -id 110101199001010000")
	fmt.Println("    （商家注册）registerb -w mike -name mike -nn mk -age 18 -tel 13812345678 -id 110101199001010000 -bid 50001000-3 -bn 北京城市网邻信息技术有限公司")
	fmt.Println("    （查询用户信息）getuser -w mike")
	fmt.Println("-----------------------------------信息-------------------------------------------------------")
	fmt.Println("    （发布信息）createpost -w mike -title 北京地区搬家 -bn 哥俩好搬家公司 -content 负责朝阳区搬家业务 -price 200 -city 北京 ")
	fmt.Println("    （根据商家获取信息列表）getbusorder -w mike")
	fmt.Println("    （根据信息ID查询信息）getorderdetail -id id")
	fmt.Println("-----------------------------------合约-------------------------------------------------------")
	fmt.Println("    （获取合约列表）getchaincode")
	fmt.Println("    （匹配）usechaincode -id 链码id -city 北京 -lowp 50 -highp 1000")
	fmt.Println("-----------------------------------订单-------------------------------------------------------")
	fmt.Println("    （用户获取订单列表）getcusorder -w mike")
	fmt.Println("    （商家获取订单列表）getbusorder -w mike")
	fmt.Println("    （用户下单）placeorder -w mike -id postId")
	fmt.Println("    （商家确认订单）confirmorder -w mike -id orderId")
	fmt.Println("    （用户完成订单）finishorder -w mike -id orderId")
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
	} else {
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
	} else {
		post.PostCommit(wlt)
	}

}
