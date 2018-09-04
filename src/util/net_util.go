package util

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"strings"
	"net/url"
	"constant"
	"encoding/json"
)


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
	fmt.Println("paramsParse:",paramsA)
	fmt.Println()

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
	return body
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
