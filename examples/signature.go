package main

import (
	"crypto/ecdsa"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	eth_common "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	eth_crypto "github.com/ethereum/go-ethereum/crypto"
	"log"
	"math/big"
	// "sig-gowallet/crypto"
)

func main2() {
	fmt.Println("sign...")

	// hex to private key
	privateKey, err := eth_crypto.HexToECDSA("6849cffaa78b3f13d0c3864cfddc03900b7c34ac60c0ed3fef3a63c7d2510ed8")
	// privateKey, err := eth_crypto.HexToECDSA("289c2857d4598e37fb9647507e47a309d6133539bf21a8b9cb6df88fd5232032")
	if err != nil {
		fmt.Printf("%v", err)
	}

	// fmt.Println("%v", privateKey)

	// var tx = map[string]interface{}{}
	var tx = make(map[string]string)
	tx["nonce"] = "0x0"
	tx["gasLimit"] = "0x5208"
	tx["gasPrice"] = "0x3b9aca00"
	tx["to"] = "0x5f5aB1692181B2c4dE255B07Be1a4A78Ea95DBD3"
	tx["value"] = "0x16345785d8a0000"
	tx["data"] = ""
	tx["chainId"] = "5"

	txBytes, err := json.Marshal(tx)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Printf("%s\n", txBytes)

	txData := sha256.Sum256(txBytes)

	// txBytes := sha256.Sum256([]byte("ethereum"))
	sign, err := eth_crypto.Sign(txData[:], privateKey)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("sign length: ", len(sign))
	fmt.Println("sign hex: ", hex.EncodeToString(sign))

	// Valid
	signCopy := decodeHex(hex.EncodeToString(sign))

	var p *ecdsa.PublicKey
	p, ok := privateKey.Public().(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	pubKeyBytes := eth_crypto.FromECDSAPub(p)
	verify := eth_crypto.VerifySignature(pubKeyBytes, txData[:], signCopy[:len(signCopy)-1])
	fmt.Println("verify result: ", verify)

	// addr := eth_crypto.PubkeyToAddress(*p)
	addr := eth_common.HexToAddress("0x5f5aB1692181B2c4dE255B07Be1a4A78Ea95DBD3")

	txD := &types.LegacyTx{
		Nonce:    1,
		To:       &addr,
		Value:    big.NewInt(100000000000000000),
		Gas:      21000,
		GasPrice: big.NewInt(1000000000),
		Data:     []byte(""),
	}

	signer := types.NewEIP155Signer(big.NewInt(888))
	txSign, err := types.SignNewTx(privateKey, signer, txD)
	if err != nil {
		log.Fatal(err)
	}
	v, r, s := txSign.RawSignatureValues()
	fmt.Printf("tx sign V=%d,R=%d,S=%d\n", v, r, s)
	signedTx, _ := txSign.MarshalJSON()
	fmt.Printf("%s\n", signedTx)
	fmt.Println(hex.EncodeToString(signedTx))
}

func decodeHex(h string) []byte {
	bytes, err := hex.DecodeString(h)
	if err != nil {
		log.Fatal(err)
	}
	return bytes
}
