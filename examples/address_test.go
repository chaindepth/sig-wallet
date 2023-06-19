package main

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"testing"
)

func TestAddress(t *testing.T) {
	// 从一个 16 进制的字符串中还原出私钥
	privateKey, err := crypto.HexToECDSA("a051c2bb65b362305ff331b86c461d02abc881baddff8370dc3a5196e59806af")
	if err != nil {
		log.Fatal(err)
	}

	// 通过私钥得到公钥
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	// 通过公钥得到地址
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Printf("%s, %s, %s", fromAddress, publicKeyECDSA.X, publicKeyECDSA.Y)
}
