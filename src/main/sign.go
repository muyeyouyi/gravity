package main

//import (
//	"constant"
//	"crypto/ecdsa"
//	"crypto/rand"
//	"fmt"
//	"util"
//	"wallet"
//)

///**
//	签名,base64编码
// */
//func Sign(wlt wallet.Wallet,content string) map[string]string{
//	args:= make(map[string]string)
//	args[constant.Args0] = string(util.Base58Encode(wlt.publicKey))
//	args[constant.Args1] = content
//	r, s, err := ecdsa.Sign(rand.Reader, &wlt.privateKey, []byte(content))
//	if err != nil {
//		fmt.Println(err)
//	}
//	signature := append(r.Bytes(), s.Bytes()...)
//	args[constant.Args2] = string(util.Base58Encode(signature))
//	return args
//}