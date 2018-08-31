package main

import (
	"wallet"
	"crypto/ecdsa"
	"crypto/rand"
	"fmt"
	"encoding/base64"
	"constant"
)

/**
	签名,base64编码
 */
func Sign(wlt wallet.Wallet,content string) map[string]string{
	args:= make(map[string]string)
	args[constant.Args1] = base64.StdEncoding.EncodeToString(wlt.PublicKey)
	args[constant.Args2] = content
	r, s, err := ecdsa.Sign(rand.Reader, &wlt.PrivateKey, []byte(content))
	if err != nil {
		fmt.Println(err)
	}
	signature := append(r.Bytes(), s.Bytes()...)
	args[constant.Args3] = base64.StdEncoding.EncodeToString(signature)
	args[constant.AccessToken] = AccessToken
	return args
}