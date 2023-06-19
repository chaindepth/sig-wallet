package main

func main33() {
	// 请求参数编码
	// 也可能没有，比如 transfer eth

	// 构造交易对象
	//to := common.HexToAddress("0x5f5aB1692181B2c4dE255B07Be1a4A78Ea95DBD3")
	//
	//txData := &types.LegacyTx{
	//	Nonce:    uint64(20),
	//	To:       &to,
	//	Value:    big.NewInt(10000000000000000),
	//	Gas:      uint64(21000),
	//	GasPrice: big.NewInt(80569709814),
	//	Data:     nil,
	//}
	//rawTx := types.NewTx(txData)
	//jsonRawTx, _ := rawTx.MarshalJSON()
	//fmt.Println("rawTx: ", string(jsonRawTx))
	//
	//// 交易签名
	//chainId := int64(5)
	//signer := types.NewEIP155Signer(big.NewInt(chainId))
	//prv, err := crypto.HexToECDSA("6849cffaa78b3f13d0c3864cfddc03900b7c34ac60c0ed3fef3a63c7d2510ed8")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//signedTx, err := types.SignTx(rawTx, signer, prv)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//signedTxBytes, _ := signedTx.MarshalJSON();
	//fmt.Println("signed tx: ", string(signedTxBytes))
	//
	//fmt.Println(common.Bytes2Hex(signedTxBytes))

	//client, err := ethclient.Dial("https://goerli.infura.io/v3/e9cb57d516ef46628b8fb5bc8eadfa1f")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//err = client.SendTransaction(context.Background(), signedTx)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//fmt.Println("Tx sent, txHash: ", signedTx.Hash().Hex())
}
