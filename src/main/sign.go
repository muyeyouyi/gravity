package main

import (
	"wallet"
	"crypto/ecdsa"
	"crypto/rand"
	"fmt"
	"encoding/base64"
)

/**
	签名,base64编码
 */
func Sign(wlt wallet.Wallet,content string) map[string]string{
	args:= make(map[string]string)
	args["args0"] = base64.StdEncoding.EncodeToString(wlt.PublicKey)
	args["args1"] = content
	r, s, err := ecdsa.Sign(rand.Reader, &wlt.PrivateKey, []byte(content))
	if err != nil {
		fmt.Println(err)
	}
	signature := append(r.Bytes(), s.Bytes()...)
	args["args2"] = base64.StdEncoding.EncodeToString(signature)
	args["accessToken"] = AccessToken
	return args
}