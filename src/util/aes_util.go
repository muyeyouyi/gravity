package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

func aesEncrypt(origData, key []byte) []byte {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil
	}
	blockSize := block.BlockSize()
	origData = pKCS5Padding(origData, blockSize)
	// origData = ZeroPadding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	crypted := make([]byte, len(origData))
	// 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
	// crypted := origData
	blockMode.CryptBlocks(crypted, origData)
	return crypted
}

func aesDecrypt(crypted, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	origData = pKCS5UnPadding(origData)
	// origData = ZeroUnPadding(origData)
	return origData, nil
}

func pKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func pKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

/**
对字段加密和base58编码
*/
func AesAndBase58Encode(info string, key []byte) string {
	encrypt := aesEncrypt([]byte(info), key)
	encodeString := Base58(encrypt)
	return encodeString
}

/**
base58解码和解密
*/
func AesAndBase58Decode(info string, key []byte) string {
	decodeString := Base58Decode([]byte(info))
	decrypt, e := aesDecrypt(decodeString, key)
	LogE(e)
	return string(decrypt)
}

func Base58(pubkey []byte) string {
	return string(Base58Encode(pubkey))
}
