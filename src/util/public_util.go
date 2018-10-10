package util
//
//import (
//	"crypto/ecdsa"
//	"crypto/elliptic"
//	"math/big"
//)
//
///**
//	验证签名
// */
//func Verify(pubkey, json, sign string) bool {
//	curve := elliptic.P256()
//	pubkeyByte := Base58Decode([]byte(pubkey))
//
//	//拆分签名文件
//	r := big.Int{}
//	s := big.Int{}
//	signByte:= Base58Decode([]byte(sign))
//	sigLen := len(signByte)
//	r.SetBytes(signByte[:(sigLen / 2)])
//	s.SetBytes(signByte[(sigLen / 2):])
//
//	//拆分公钥
//	x := big.Int{}
//	y := big.Int{}
//	keyLen := len(pubkeyByte)
//	x.SetBytes(pubkeyByte[:(keyLen / 2)])
//	y.SetBytes(pubkeyByte[(keyLen / 2):])
//	//还原为原始公钥
//	rawPubKey := ecdsa.PublicKey{curve, &x, &y}
//	//公钥、签名文件、原始数据确认签名有效性
//	if ecdsa.Verify(&rawPubKey, []byte(json), &r, &s) == false {
//		return false
//	}
//	return true
//}
//
//
