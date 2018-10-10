package util

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math/big"
)

const walletFile = "wallet.dat"

type Wallet struct {
	privateKey ecdsa.PrivateKey
	publicKey  []byte
	aesKey     []byte
}

/**
	私钥导出
 */
func (wlt *Wallet) GetPriKey() string {
	encode := Base58Encode(wlt.privateKey.D.Bytes())
	str := string(encode)
	return str
}

/**
	公钥导出
 */
func (wlt *Wallet) GetPubKey() string {
	return Base58(wlt.publicKey)
}

func (wlt *Wallet) Sign(info string) string {
	r, s, err := ecdsa.Sign(rand.Reader, &wlt.privateKey, []byte(info))
	if err != nil {
		fmt.Println(err)
	}
	signature := append(r.Bytes(), s.Bytes()...)
	//fmt.Println(signature)
	//fullSign := append([]byte{0x00},signature...)
	//fmt.Println("原始签名：",signature)
	sign := string(Base58Encode(signature))
	return sign
}

/**
	验证签名
 */
func (wlt *Wallet) Verify(pubkey, json, sign string) bool {
	curve := elliptic.P256()
	//pubkeyByte := Base58Decode([]byte(pubkey))
	pubkeyByte := wlt.publicKey

	//拆分签名文件
	r := big.Int{}
	s := big.Int{}

	fullSign := Base58Decode([]byte(sign))
	//trueSign := fullSign[1:]

	//trueSign := Base58Decode([]byte(sign))
	//fmt.Println("base58解码:",trueSign)
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

/**
 */
func NewWallet() *Wallet {
	privateKey, publicKey := newKeyPair()
	aesKey := sha256.Sum256(privateKey.D.Bytes())
	return &Wallet{privateKey, publicKey, aesKey[:]}
}

/**
 */
func ImportWallet(p string) *Wallet {
	decode := Base58Decode([]byte(p))
	bigint := new(big.Int)
	bigint.SetBytes(decode)
	privateKey, publicKey := importPrikey(bigint)
	aesKey := sha256.Sum256(privateKey.D.Bytes())
	return &Wallet{privateKey, publicKey, aesKey[:]}
}

/**
 */
func newKeyPair() (ecdsa.PrivateKey, []byte) {
	curve := elliptic.P256()

	private, _ := ecdsa.GenerateKey(curve, rand.Reader)
	pubKey := bytePubkey(private)
	return *private, pubKey
}

/**
 */
func bytePubkey(private *ecdsa.PrivateKey) []byte {
	pubKey := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)
	return pubKey
}

/**
 */
func importPrikey(k *big.Int) (ecdsa.PrivateKey, []byte) {
	private := new(ecdsa.PrivateKey)
	c := elliptic.P256()
	private.PublicKey.Curve = c
	private.D = k
	private.PublicKey.X, private.PublicKey.Y = c.ScalarBaseMult(k.Bytes())
	pubKey := bytePubkey(private)
	return *private, pubKey
}
