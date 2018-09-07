package constant

const (
	BaseUrl      = "https://baas.58.com/chaincode/"
	UrlInvoke    = BaseUrl + "invoke"
	UrlQuery     = BaseUrl + "query"
	AppIdGravity = "499169"
	AppSecret    = "1834b876583b310fb4f2ea363f5a62c4"

	Function      = "function"
	Version       = "version"
	AppId         = "appId"
	Args0         = "args0"
	Args1         = "args1"
	Args2         = "args2"
	Args3         = "args3"
	ChainCodeName = "chaincodeName"
	AccessToken   = "accessToken"

	//ChainCodeName
	ChainCodeUser     = "user"
	ChainCodeInfo     = "info"
	ChainCodeTrade    = "trade"
	ChainCodeMatching = "matching"
	//Function
	Set          = "set"
	Get          = "get"
	GetByOwner   = "getByOwner"
	Matching     = "matching"
	MatchingList = "matchingList"
	Submit       = "submit"
	Confirm       = "confirm"
	Finish       = "finish"
	GetTradeByConstumer       = "getTradeByConstumer"
	GetTradeByBusiness       = "getTradeByBusiness"

	//chainCodeVersion
	UserVersion     = "12"
	InfoVersion     = "10"
	MatchingVersion = "9"
	TradeVersion    = "9"

	//缓存文件
	OrderIdFile = "orderid.dat"
	ChainCodeFile = "chaincodeid.dat"
)
