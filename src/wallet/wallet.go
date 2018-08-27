package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
)

const walletFile = "wallet.dat"

type Wallet struct {
	PrivateKey ecdsa.PrivateKey
	PublicKey  []byte
	AesKey []byte
}

/**
	创建一个新钱包
 */
func NewWallet() *Wallet {
	privateKey, publicKey := newKeyPair()
	aesKey := sha256.Sum256(privateKey.D.Bytes())
	return &Wallet{privateKey, publicKey,aesKey[:]}
}

/**
	生成一对公、私钥
 */
func newKeyPair() (ecdsa.PrivateKey, []byte) {
	curve := elliptic.P256()
	private, _ := ecdsa.GenerateKey(curve, rand.Reader)
	pubKey := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)
	return *private, pubKey
}

