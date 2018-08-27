package util

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"net/url"
)

func GetTest() {
	//构造url
	u, err := url.Parse("http://www.baidu.com?")
	if err != nil {
		fmt.Println("url parse fail")
		return
	}
	q := u.Query()
	q.Set("name", "wnw")
	q.Set("sex", "wowam")
	u.RawQuery = q.Encode()
	fmt.Println("url:",u.String())
	//发起get请求
	resp, err1 := http.Get(u.String())
	if err1 != nil || resp.StatusCode != http.StatusOK {
		fmt.Println("get fail:", err1)
		return
	}
	defer resp.Body.Close()
	//读取响应体
	body, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		fmt.Println("read body fail")
		return
	}
	fmt.Println(string(body))
	//解析数据
}
