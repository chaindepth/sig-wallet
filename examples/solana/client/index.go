package main

import (
	"context"
	"fmt"
	"github.com/blocto/solana-go-sdk/client"
	"github.com/blocto/solana-go-sdk/rpc"
)

const testAddress = "DsqkS4YYodw1PPeUZTgaaYrpZ5wythns9QmFBV3VnjiC"

func main() {
	devNet()
	mainnet()
}

func mainnet() {
	c := client.NewClient(rpc.MainnetRPCEndpoint)
	balance, _ := c.GetBalance(context.Background(), "toly.sol")
	fmt.Println("Wallet Balance in Lamport:", balance)
}

func devNet() {
	c := client.NewClient(rpc.DevnetRPCEndpoint)
	balance, _ := c.GetBalance(context.Background(), testAddress)
	fmt.Println("Wallet Balance in Lamport:", balance)
}
