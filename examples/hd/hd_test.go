package main

import (
	"fmt"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"github.com/stretchr/testify/assert"
	"log"
	"strings"
	"testing"
)

func TestCase1(t *testing.T) {
	mnemonic := "tag volcano eight thank tide danger coast health above argue embrace heavy"
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		log.Fatal(err)
	}

	path0 := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")
	account0, err := wallet.Derive(path0, false)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(account0.Address.Hex())
	assert.True(t, strings.EqualFold(account0.Address.Hex(), "0xc49926c4124cee1cba0ea94ea31a6c12318df947"))

	path1 := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/1")
	account1, err := wallet.Derive(path1, false)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(account1.Address.Hex())
	assert.True(t, strings.EqualFold(account1.Address.Hex(), "0x8230645ac28a4edd1b0b53e7cd8019744e9dd559"))

	// ETH
	// index 0: 0xc49926c4124cee1cba0ea94ea31a6c12318df947
	// index 1: 0x8230645ac28a4edd1b0b53e7cd8019744e9dd559

	// SOL
	// index 0: 92rsLkdmvz7Y4a6FkgAcAHu6RyMRn1dwug4NCrcgVBMT
	// index 1: GtsWe4XrepDGDoj7GqQnUES2UrTHZPUrtniFHa2MLCtd
}
