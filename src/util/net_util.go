package util

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"strings"
	"net/url"
	"constant"
	"encoding/json"
	"bytes"
	"encoding/gob"
	"crypto/elliptic"
	"log"
	"os"
)
type Response struct {
	Status int `json:"status"`
	Message string `json:"message"`
	BlockNumber string `json:"blockNumber"`
	Data []string `json:"data"`
}


func PostAccessToken(questUrl string,args map[string]string)[]byte {
	fullUrl := questUrl
	for key, value := range args {
		fullUrl+= key
		fullUrl+= "="
		fullUrl+= value
		fullUrl += "&"
	}
	//fmt.Println("url0:",fullUrl)
	fullUrl = string(([]byte(fullUrl))[:len(fullUrl)-1])
	//fmt.Println("url:",fullUrl)

	u, _ := url.Parse(fullUrl)
	q := u.Query()
	u.RawQuery = q.Encode()
	//fmt.Println("urlparse:",u.String())

	resp, err := http.Post(u.String(),
		"application/x-www-form-urlencoded",
		strings.NewReader(""))
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println("response:",string(body))
	return body
}


func PostTest(questUrl string,args map[string]string)[]byte {
	args[constant.AccessToken] = getAccessToken()

	var params string
	for key, value := range args {
		params += key
		params += "="
		params += value
		params += "&"
	}
	//fmt.Println("url0:",params)
	params = string(([]byte(params))[:len(params)-1])
	fmt.Println("url:", questUrl)
	fmt.Println("params:", params)
	fmt.Println()

	//u, _ := url.Parse(params)
	//q := u.Query()
	//u.RawQuery = q.Encode()
	//fmt.Println("paramsParse:",u.String())

	//参数转译
	paramsA := ""
	values := url.Values{}
	for k, v := range args {
		values.Add(k, v)
	}
	paramsA = values.Encode()
	//fmt.Println("paramsParse:",paramsA)
	//fmt.Println()

	resp, err := http.Post(questUrl,
		"application/x-www-form-urlencoded",
		strings.NewReader(paramsA))
	if err != nil {
		fmt.Println(err)
	}


	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("response:",string(body))
	fmt.Println()

	defer func() {
		if e := recover(); e!= nil{
			//log.Fatalln("error:json解析异常")
		}
	}()
	var response Response
	json.Unmarshal(body, &response)
	if response.Status != 200 {
		panic("请求失败！")
	}
	info := response.Data[0]
	//if info != "" {
	//	fmt.Println("解析:",info)
	//}
	return []byte(info)
}

func getAccessToken() string{
	type Data struct {
		ExpireIn    int    `json:"expireIn"`
		AccessToken string `json:"accessToken"`
	}
	type Response struct {
		Code int  `json:"code"`
		Data Data `json:"data"`
	}
	orgs := make(map[string]string)
	orgs[constant.AppId] = constant.AppIdGravity
	orgs["appSecret"] = constant.AppSecret
	response := PostAccessToken("https://baas.58.com/token/clientCredentials?", orgs)
	var res Response
	json.Unmarshal(response, &res)
	return res.Data.AccessToken
}

type Ids struct {
	Ids map[string]string
}


func SaveId(ids map[string]string,fileName string)  {
	oldIds:= ReadId(fileName)
	if len(oldIds)== 0{
		oldIds = make(map[string]string)
	}else{
		for key, value := range ids {
			oldIds[key] = value
		}
	}


	idStruct := &Ids{ids}

	var content bytes.Buffer

	gob.Register(elliptic.P256())

	encoder := gob.NewEncoder(&content)
	err := encoder.Encode(idStruct)
	if err != nil {
		log.Panic(err)
	}

	err = ioutil.WriteFile(fileName, content.Bytes(), 0644)
	if err != nil {
		log.Panic(err)
	}
}

/**
	读取本地钱包信息
 */
func ReadId(fileName string) map[string]string{
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		return nil
	}

	fileContent, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Panic(err)
	}

	var idStruct Ids

	gob.Register(elliptic.P256())
	decoder := gob.NewDecoder(bytes.NewReader(fileContent))
	err = decoder.Decode(&idStruct)
	//fmt.Println("map",ids)
	if err != nil {
		log.Panic(err)
	}

	return idStruct.Ids
}
