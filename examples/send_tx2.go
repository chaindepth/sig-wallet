package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/hex"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
	"log"
	"math/big"
)

func main() {

	methodId := crypto.Keccak256([]byte("setA(uint256)"))[:4]
	fmt.Println("methodId: ", common.Bytes2Hex(methodId))
	paramValue := math.U256Bytes(new(big.Int).Set(big.NewInt(123)))
	fmt.Println("paramValue: ", common.Bytes2Hex(paramValue))
	input := append(methodId, paramValue...)
	fmt.Println("input: ", common.Bytes2Hex(input))

	// 使用 ethclient 连接到 ethereum network 上的某个 node
	client, err := ethclient.Dial("https://goerli.infura.io/v3/e9cb57d516ef46628b8fb5bc8eadfa1f")
	if err != nil {
		log.Fatal(err)
	}
	client.BlockNumber(context.Background())
	//block, err := client.BlockByHash(context.Background(), common.HexToHash("0x..."))
	//fmt.Printf(block.Hash().Hex())

	// 从一个 16 进制的字符串中还原出私钥
	privateKey, err := crypto.HexToECDSA("6849cffaa78b3f13d0c3864cfddc03900b7c34ac60c0ed3fef3a63c7d2510ed8")
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

	//if fromAddress.Bytes() != nil {
	//	os.Exit(0)
	//}

	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	value := big.NewInt(10000000000000000) // in wei (0.01 eth)
	gasLimit := uint64(21000)              // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	toAddress := common.HexToAddress("0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d")
	var data []byte
	txData := &types.LegacyTx{
		Nonce:    nonce,
		To:       &toAddress,
		Value:    value,
		Gas:      gasLimit,
		GasPrice: gasPrice,
		Data:     data,
	}
	// tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)
	tx := types.NewTx(txData)

	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	// ts := types.Transactions{signedTx}
	// rawTxBytes := ts.GetRlp(0)
	toBytes, _ := rlp.EncodeToBytes(signedTx)
	rawTxHex := hex.EncodeToString(toBytes)

	fmt.Printf(rawTxHex) // f86...772

	//err = client.SendTransaction(context.Background(), signedTx)
	//if err != nil {
	//	log.Fatal(err)
	//}

}
