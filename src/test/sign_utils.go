package main

import (
	"bytes"
	"crypto/ecdsa"
	"crypto/elliptic"
	"math/big"
)

var b58Alphabet = []byte("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")

/**
	验证签名
 */
func Verify(pubkey, json, sign string) bool {
	curve := elliptic.P256()
	pubkeyByte := Base58Decode([]byte(pubkey))

	//拆分签名文件
	r := big.Int{}
	s := big.Int{}

	fullSign := Base58Decode([]byte(sign))
	sigLen := len(fullSign)
	r.SetBytes(fullSign[:(sigLen / 2)])
	s.SetBytes(fullSign[(sigLen / 2):])

	//拆分公钥
	x := big.Int{}
	y := big.Int{}
	keyLen := len(pubkeyByte)
	x.SetBytes(pubkeyByte[:(keyLen / 2)])
	y.SetBytes(pubkeyByte[(keyLen / 2):])
	//还原为原始公钥
	rawPubKey := ecdsa.PublicKey{curve, &x, &y}
	//公钥、签名文件、原始数据确认签名有效性
	if ecdsa.Verify(&rawPubKey, []byte(json), &r, &s) == false {
		return false
	}
	return true
}

func Base58Decode(input []byte) []byte {
	result := big.NewInt(0)
	zeroBytes := 0

	for b := range input {
		if b == 0x00 {
			zeroBytes++
		}
	}

	payload := input[zeroBytes:]
	for _, b := range payload {
		charIndex := bytes.IndexByte(b58Alphabet, b)
		result.Mul(result, big.NewInt(58))
		result.Add(result, big.NewInt(int64(charIndex)))
	}

	decoded := result.Bytes()
	decoded = append(bytes.Repeat([]byte{byte(0x00)}, zeroBytes), decoded...)

	return decoded[1:]
}
