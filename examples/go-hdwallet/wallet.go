package main

import (
	"fmt"
	"github.com/foxnut/go-hdwallet"
)

var (
	testMnemonic = "range sheriff try enroll deer over ten level bring display stamp recycle"
)

func main() {
	master, _ := hdwallet.NewKey(hdwallet.Mnemonic(testMnemonic))

	wallet, _ := master.GetWallet(hdwallet.CoinType(hdwallet.BTC), hdwallet.AddressIndex(1))
	address, _ := wallet.GetAddress()
	addressP2WPKH, _ := wallet.GetKey().AddressP2WPKH()
	addressP2WPKHInP2SH, _ := wallet.GetKey().AddressP2WPKHInP2SH()
	fmt.Println("BTC: ", address, addressP2WPKH, addressP2WPKHInP2SH)
}
