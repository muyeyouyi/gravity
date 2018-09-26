package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"math/big"
	"util"
)

const walletFile = "wallet.dat"

type Wallet struct {
	PrivateKey ecdsa.PrivateKey
	PublicKey  []byte
	AesKey []byte
}

/**
	私钥导出
 */
func (wlt *Wallet) Export() string  {
	encode := util.Base58Encode(wlt.PrivateKey.D.Bytes())
	str := string(encode)
	return str
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
	导入钱包
 */
func ImportWallet(p string) *Wallet  {
	decode := util.Base58Decode([]byte(p))
	bigint := new(big.Int)
	bigint.SetBytes(decode)
	privateKey, publicKey := importPrikey(bigint)
	aesKey := sha256.Sum256(privateKey.D.Bytes())
	return &Wallet{privateKey, publicKey,aesKey[:]}
}
/**
	生成一对公、私钥
 */
func newKeyPair() (ecdsa.PrivateKey, []byte) {
	curve := elliptic.P256()

	private, _ := ecdsa.GenerateKey(curve, rand.Reader)
	pubKey := getPubkey(private)
	return *private, pubKey
}

/**
	由私钥生成公钥
 */
func getPubkey(private *ecdsa.PrivateKey) []byte {
	pubKey := append(private.PublicKey.X.Bytes(), private.PublicKey.Y.Bytes()...)
	return pubKey
}

/**
	导入私钥
 */
func importPrikey(k *big.Int) (ecdsa.PrivateKey,[]byte) {
	private := new(ecdsa.PrivateKey)
	c := elliptic.P256()
	private.PublicKey.Curve = c
	private.D = k
	private.PublicKey.X, private.PublicKey.Y = c.ScalarBaseMult(k.Bytes())
	pubKey := getPubkey(private)
	return *private,pubKey
}


