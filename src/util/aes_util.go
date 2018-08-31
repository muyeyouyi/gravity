package util

import (
	"crypto/aes"
	"crypto/cipher"
	"bytes"
	"encoding/base64"
)

func AesEncrypt(origData, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil
	}
	blockSize := block.BlockSize()
	origData = PKCS5Padding(origData, blockSize)
	// origData = ZeroPadding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	// 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
	// crypted := origData
	blockMode.CryptBlocks(crypted, origData)
	return crypted
}

func AesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	origData = PKCS5UnPadding(origData)
	// origData = ZeroUnPadding(origData)
	return origData, nil
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

/**
	对字段加密和base58编码
 */
func AesAndBase64Encode(info string,key []byte) string {
	encrypt := AesEncrypt([]byte(info), key)
	encodeString := base64.StdEncoding.EncodeToString(encrypt)
	return encodeString
}

/**
	base58解码和解密
 */
func AesAndBase64Decode(info string,key []byte) string  {
	decodeString, e := base64.StdEncoding.DecodeString(info)
	decrypt, e := AesDecrypt(decodeString, key)
	LogE(e)
	return string(decrypt)
}

func Base64(pubkey []byte) string  {
	return base64.StdEncoding.EncodeToString(pubkey)
}