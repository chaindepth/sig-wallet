package main

import (
	"fmt"
	"github.com/tyler-smith/go-bip39"
)

func main() {
	fmt.Println("bip39")

	entropy, _ := bip39.NewEntropy(256)
	mnemonic, _ := bip39.NewMnemonic(entropy)

	fmt.Println(mnemonic)

	// 我们只要拿到助记词和密码，就能得到 seed，有了 seed 就能得到主密钥
	seed1 := bip39.NewSeed(mnemonic, "hello")
	seed2 := bip39.NewSeed(mnemonic, "")

	fmt.Println(seed1)
	fmt.Println(seed2)

	seed3, _ := bip39.MnemonicToByteArray(mnemonic)
	seed4, _ := bip39.NewSeedWithErrorChecking(mnemonic, "")
	fmt.Println(seed3)
	fmt.Println(seed4)
}
