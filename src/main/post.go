package main

import (
	"constant"
	"util"
	"wallet"
	"encoding/json"
	"fmt"
	"time"
	"strconv"
)

//type PostListResponse struct {
//	Status int `json:"status"`
//	Message string `json:"message"`
//	BlockNumber string `json:"blockNumber"`
//	Data []string `json:"data"`
//}

type Info struct {
	ID          string
	PubKey      string
	Title       string
	Content     string
	CompanyName string
	City        string
	Price       int
	PublishTime time.Time
}


/**
	帖子信息
 */
type Post struct {
	Title       string `json:"Title"`
	Content     string `json:"Content"`
	CompanyName string `json:"CompanyName"`
	City        string `json:"City"`
	Price       int    `json:"Price"`
}


/**
	帖子信息
 */
func (post *Post) PostCommit(wlt wallet.Wallet) {
	//生成json
	jsonByte, e := json.Marshal(post)
	if e != nil {
		fmt.Println(e)
	}

	args := Sign(wlt, string(jsonByte))
	args[constant.Function] = constant.Set
	args[constant.Version] = constant.InfoVersion
	args[constant.ChainCodeName] = constant.ChainCodeInfo
	args[constant.AppId] = constant.AppIdGravity
	util.PostTest(constant.UrlInvoke, args)
}

/**
	查询一个商家的所有帖子
 */
func (post *Post) GetPosts(pubkey string) {
	args := make(map[string]string)
	args[constant.ChainCodeName] = constant.ChainCodeInfo
	args[constant.Version] = constant.InfoVersion
	args[constant.Function] = constant.GetByOwner
	args[constant.AppId] = constant.AppIdGravity

	args[constant.Args0] = pubkey
	res := util.PostTest(constant.UrlQuery, args)
	post.analysis(res)
}
func (post *Post) analysis(res []byte)  {
	var array map[string]string
	json.Unmarshal(res,&array)
	count := 1
	ids := make(map[string]string)
	for id, value := range array {
		info := analysisInfo([]byte(value), count)
		key := info.CompanyName + strconv.Itoa(count)
		ids[key] = id
		count++
	}
	util.SaveId(ids, constant.OrderIdFile)

}

func analysisInfo(value []byte, count int) Info {
	var info Info
	json.Unmarshal(value, &info)
	if count != 0 {
		fmt.Println("信息", count, ":")
	}
	fmt.Println("公司：", info.CompanyName)
	fmt.Println("标题：", info.Title)
	fmt.Println("内容：", info.Content)
	fmt.Println("城市：", info.City)
	fmt.Println("价格：", info.Price)
	fmt.Println("时间戳：", info.PublishTime)
	fmt.Println()
	return info
}

/**
	查询一个帖子详情
 */
func (post *Post) GetPostDetail(ID string) {
	ids := util.ReadId(constant.OrderIdFile)

	args := make(map[string]string)
	args[constant.ChainCodeName] = constant.ChainCodeInfo
	args[constant.Version] = constant.InfoVersion
	args[constant.Function] = constant.Get
	args[constant.AppId] = constant.AppIdGravity

	args[constant.Args0] = ids[ID]
	res := util.PostTest(constant.UrlQuery, args)
	analysisInfo(res,0)
}


