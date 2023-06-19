package main

import (
	"fmt"

	"github.com/miguelmota/go-ethereum-hdwallet"
)

func main() {
	fmt.Println("hd testing ...")

	// Test mnemonic: salt seven nothing auction catch clerk climb hub hurry eager shine offer clock poem panda goddess below repair address marriage toy social coyote rice
	var mnemonic = "salt seven nothing auction catch clerk climb hub hurry eager shine offer clock poem panda goddess below repair address marriage toy social coyote rice"
	wallet, _ := hdwallet.NewFromMnemonic(mnemonic)

	path, _ := hdwallet.ParseDerivationPath("m/44'/60'/0'/0/0")
	fmt.Println(path)

	p := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/100")
	fmt.Println(p)

	a, _ := wallet.Derive(p, false)
	fmt.Println(a)

	privhex, _ := wallet.PrivateKeyHex(a)
	pubhex, _ := wallet.PublicKeyHex(a)
	addr, _ := wallet.Address(a)
	fmt.Println(privhex, pubhex, addr)

	// fmt.Println(wallet.PrivateKey())

	// fmt.Println(hdwallet.DefaultRootDerivationPath)

	wallet.Accounts()
}
