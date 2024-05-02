package main

import (
	"context"
	"fmt"
	"github.com/blocto/solana-go-sdk/client"
	"github.com/blocto/solana-go-sdk/common"
	"github.com/blocto/solana-go-sdk/program/metaplex/token_metadata"
	"github.com/blocto/solana-go-sdk/rpc"
	"github.com/blocto/solana-go-sdk/types"
	"testing"
)

func TestFindPDA(t *testing.T) {
	pubKey, nonce, _ := common.FindProgramAddress([][]byte{[]byte("test")}, common.PublicKeyFromString("2PKXDnSYQf8jbMMGi66V7optX3UfVaqSiRkoPQACDDZD"))
	fmt.Println(pubKey, nonce)
}

func TestGetMint(t *testing.T) {
	// 4Sfpc2gVBwawgZ9uW4tEA1NqRuZ9LQ8YmZk3MXZz9R4v
	// GeoWLwRr25dUpdDDqyFqcgk7m2MaJgw2VnSJGUfjFaXv
	c := client.NewClient(rpc.DevnetRPCEndpoint)
	getTokenMint(c, "4Sfpc2gVBwawgZ9uW4tEA1NqRuZ9LQ8YmZk3MXZz9R4v")
}

func TestCreateTokenMint(t *testing.T) {
	c := client.NewClient(rpc.DevnetRPCEndpoint)
	tokenMintAccount := types.NewAccount()
	createTokenMint(c, tokenMintAccount, 8)
}

func TestA(t *testing.T) {
	t.Log(len("1gBDpsAR2Eknv3PQZh7tQCvLopy6L3L1Eq5FhMCkqU7"))
	t.Log(len("GqXWxJw6eq8zW7uRXuRJTE3dSwAPJ3mdcPqt2xgExZtt"))
}

func TestDoTransfer(t *testing.T) {
	c := client.NewClient(rpc.DevnetRPCEndpoint)

	destinationAccount := getDestinationAccount()
	mintAuthAccount := getMintAuthorityAccount()
	feePayerAccount := getFeePayerAccount()

	tokenMintPubKey := common.PublicKeyFromString("1gBDpsAR2Eknv3PQZh7tQCvLopy6L3L1Eq5FhMCkqU7")
	ata := common.PublicKeyFromString("GqXWxJw6eq8zW7uRXuRJTE3dSwAPJ3mdcPqt2xgExZtt")
	destinationAta := common.PublicKeyFromString("FhYMQuUgnJcte4fHXfUZTPgSh4KFyZcKftrDxRgvQXpa")

	doTransfer(
		c,
		destinationAccount,
		mintAuthAccount,
		feePayerAccount,
		tokenMintPubKey,
		ata,
		destinationAta,
		25000000000,
		8,
	)
}

func TestDoTransfer2(t *testing.T) {
	c := client.NewClient(rpc.DevnetRPCEndpoint)

	destinationAccount := getMintAuthorityAccount()
	mintAuthAccount := getDestinationAccount()
	feePayerAccount := getDestinationAccount()

	tokenMintPubKey := common.PublicKeyFromString("1gBDpsAR2Eknv3PQZh7tQCvLopy6L3L1Eq5FhMCkqU7")
	destinationAta := common.PublicKeyFromString("GqXWxJw6eq8zW7uRXuRJTE3dSwAPJ3mdcPqt2xgExZtt")
	ata := common.PublicKeyFromString("FhYMQuUgnJcte4fHXfUZTPgSh4KFyZcKftrDxRgvQXpa")

	doTransfer(
		c,
		destinationAccount,
		mintAuthAccount,
		feePayerAccount,
		tokenMintPubKey,
		ata,
		destinationAta,
		2500000000,
		8,
	)
}

func TestTokenMetadata(t *testing.T) {
	c := client.NewClient(rpc.MainnetRPCEndpoint)

	mintAccount := common.PublicKeyFromString("EPjFWdd5AufqSSqeM2qN1xzybapC8G4wEGGkZwyTDt1v")

	metadataAccount, _ := token_metadata.GetTokenMetaPubkey(mintAccount)
	fmt.Println(metadataAccount.ToBase58())

	accountInfo, err := c.GetAccountInfo(context.Background(), metadataAccount.ToBase58())
	if err != nil {
		t.Fatal(err)
	}

	metadata, err := token_metadata.MetadataDeserialize(accountInfo.Data)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(metadata)
}
