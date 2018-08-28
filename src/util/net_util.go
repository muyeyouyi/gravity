package util

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"strings"
)

func PostTest(questUrl string,args map[string]string)[]byte {
	fullUrl := questUrl
	for key, value := range args {
		fullUrl+= key
		fullUrl+= "="
		fullUrl+= value
		fullUrl += "&"
	}
	//fmt.Println("url0:",fullUrl)
	fullUrl = string(([]byte(fullUrl))[:len(fullUrl)-1])
	fmt.Println("url:",fullUrl)

	resp, err := http.Post(fullUrl,
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
	fmt.Println("response:",string(body))
	return body
}
