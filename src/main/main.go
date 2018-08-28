package main

import (
	"crypto/sha256"
	"util"
	"fmt"
)

func main() {
	orgs := make(map[string]string)
	orgs["name"] = "dfz"
	orgs["age"] = "18"
	orgs["tel"] = "13811111111"
	util.PostTest("https://www.baidu.com?",orgs)
	//cli()
}

func test() {
	//s := "dfzdfz"
	//key := sha256.Sum256([]byte(s))
	//encrypt := util.AesEncrypt([]byte(s), key[:])
	//fmt.Println(encrypt)
	//toString := base64.StdEncoding.EncodeToString(encrypt)
	////base58Encode := util.Base58Encode(encrypt)
	//util.LogD(toString)
	////base58Decode := util.Base58Decode(base58Encode)
	//decodeString, _ := base64.StdEncoding.DecodeString(toString)
	//fmt.Println(decodeString)
	//bytes, _ := util.AesDecrypt(encrypt, key[:])
	//util.LogD(string(bytes))
	////decode := util.AesAndBase64Decode(s, key[:])
	////fmt.Println(decode)
	////encode := util.AesAndBase64Encode(decode, key[:])
	////fmt.Println(encode)
	//
	////util.GetTest()
}

func cli() {
	cli := &Cli{}
	cli.Run()
}

func testAes() {
	sum256 := sha256.Sum256([]byte("dfz"))
	s := "abcdefg"
	bytes := util.AesEncrypt([]byte(s), sum256[:])
	fmt.Println(bytes)
	decrypt, _ := util.AesDecrypt(bytes, sum256[:])
	fmt.Println(string(decrypt))
}
