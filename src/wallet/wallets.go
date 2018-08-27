package wallet

import (
	"os"
	"io/ioutil"
	"log"
	"encoding/gob"
	"crypto/elliptic"
	"bytes"
	"util"
	"errors"
)

type Wallets struct {
	Wallets map[string]*Wallet
}

/**
	创建钱包，从本地读取缓存
 */
func NewWallets() (*Wallets, error) {
	wallets := Wallets{}
	wallets.Wallets = make(map[string]*Wallet)
	err := wallets.LoadFromFile()

	return &wallets, err
}

/**
	读取本地钱包信息
 */
func (ws *Wallets) LoadFromFile() error {
	if _, err := os.Stat(walletFile); os.IsNotExist(err) {
		return err
	}

	fileContent, err := ioutil.ReadFile(walletFile)
	if err != nil {
		log.Panic(err)
	}

	var wallets Wallets
	gob.Register(elliptic.P256())
	decoder := gob.NewDecoder(bytes.NewReader(fileContent))
	err = decoder.Decode(&wallets)
	if err != nil {
		log.Panic(err)
	}

	ws.Wallets = wallets.Wallets

	return nil
}

/**
	钱包持久化
 */
func (ws Wallets) SaveToFile() {
	var content bytes.Buffer

	gob.Register(elliptic.P256())

	encoder := gob.NewEncoder(&content)
	err := encoder.Encode(ws)
	if err != nil {
		log.Panic(err)
	}

	err = ioutil.WriteFile(walletFile, content.Bytes(), 0644)
	if err != nil {
		log.Panic(err)
	}
}
/**
	创建一个新钱包
 */
func (ws *Wallets) CreateWallet(nickName string) {
	names := ws.GetWalletNames()
	for _, name := range names {
		if nickName == name{
			 util.LogE(errors.New("钱包名称重复，请重新输入"))
			 return
		}
	}
	wallet := NewWallet()
	ws.Wallets[nickName] = wallet
	util.LogD("钱包\""+nickName+"\"创建成功")
}

/**
	根据钱包名字获取钱包
 */
func (ws *Wallets) GetWallet(address string) *Wallet {

	wallet := ws.Wallets[address]
	return wallet
}

/**
	获取全部钱包名字
 */
func (ws *Wallets) GetWalletNames() []string{
	var names []string

	for key := range ws.Wallets {
		names = append(names,key)
	}
	return names
}

func ExamWallet(walletName string) (Wallet,error) {
	wallets, _ := NewWallets()
	wlt := wallets.GetWallet(walletName)
	if	wlt == nil{
		return Wallet{},errors.New("钱包输入错误")
	}else{
		return *wlt,nil
	}
}
