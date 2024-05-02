package main

import (
	"context"
	"fmt"
	"github.com/blocto/solana-go-sdk/client"
	"github.com/blocto/solana-go-sdk/common"
	"github.com/blocto/solana-go-sdk/program/associated_token_account"
	"github.com/blocto/solana-go-sdk/program/system"
	"github.com/blocto/solana-go-sdk/program/token"
	"github.com/blocto/solana-go-sdk/rpc"
	"github.com/blocto/solana-go-sdk/types"
)

// TestBytesSlicePrivateKey its address is DsqkS4YYodw1PPeUZTgaaYrpZ5wythns9QmFBV3VnjiC
var TestBytesSlicePrivateKey = []byte{100, 25, 253, 115, 217, 126, 60, 157, 59, 25, 74, 187, 227, 173, 85, 139, 84, 2, 43, 150, 46, 178, 115, 65, 160, 171, 187, 52, 73, 29, 70, 93, 191, 82, 103, 111, 156, 1, 99, 114, 80, 18, 41, 78, 202, 246, 46, 3, 177, 153, 169, 56, 191, 92, 165, 116, 114, 89, 250, 160, 135, 146, 221, 69}
var feePayer, _ = types.AccountFromBytes(TestBytesSlicePrivateKey)

// TestHexPrivateKey its address is BKHLRxLPiNALhf5UejxqJKjGUWCf5uaWfMP6jxNsi1pM
const TestHexPrivateKey = "df6db68d2c28db792f729d459ba62bca536c0ec9b0eb433958880a72f8b2578c9943cb3b63f80ffd41f6284e29ce2f30048aa1a429f0157577325054317b5022"

var alice, _ = types.AccountFromHex(TestHexPrivateKey)

const mintAccount = "GeoWLwRr25dUpdDDqyFqcgk7m2MaJgw2VnSJGUfjFaXv"

var mintPubKey = common.PublicKeyFromString(mintAccount)

var aliceTokenATAPubkey = common.PublicKeyFromString("2nrpwrAFwiAxAygUTjaUw9kErAke1nRnpL9iymbdkmsD")

var bob = common.PublicKeyFromString("B3HSvqGvWVCppZh6A7g7ga94eZS4DRFQNFTPi65f8snc")

func main() {
	// createMint()
	// getMint()
	// createTokenAccount()
	// getTokenAccount()
	// mintTo()
	// getBalance()
	tokenTransfer()
}

func tokenTransfer() {
	c := client.NewClient(rpc.DevnetRPCEndpoint)

	//// crate an ATA for bob
	bobATA, _, err := common.FindAssociatedTokenAddress(bob, mintPubKey)
	fmt.Println("Bob ATA:", bobATA.ToBase58())
	// AaJRN7282LesVxMRUyr2mjQLWNt3ak4Ht36fVF3XF3fY
	//
	//// create Bob's ATA
	//res, err := c.GetLatestBlockhash(context.Background())
	//if err != nil {
	//	panic(err)
	//}
	//
	//tx, err := types.NewTransaction(types.NewTransactionParam{
	//	Message: types.NewMessage(types.NewMessageParam{
	//		FeePayer:        feePayer.PublicKey,
	//		RecentBlockhash: res.Blockhash,
	//		Instructions: []types.Instruction{
	//			associated_token_account.Create(associated_token_account.CreateParam{
	//				Funder:                 feePayer.PublicKey,
	//				Owner:                  bob,
	//				Mint:                   mintPubKey,
	//				AssociatedTokenAccount: bobATA,
	//			}),
	//		},
	//	}),
	//	Signers: []types.Account{feePayer},
	//})
	//
	//if err != nil {
	//	panic(err)
	//}
	//
	//txhash, err := c.SendTransaction(context.Background(), tx)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("create bob ata' txhash:", txhash)
	// 2x5VbiJbUmoVurcjtcb8E4fGjsdYJf7F4qqr5bzEgsetXwNrZptbD4pcdLSBg7R3zFRWrPMtjBFGoCVnAS2P92XT

	// ---

	res, err := c.GetLatestBlockhash(context.Background())
	if err != nil {
		panic(err)
	}

	tx, err := types.NewTransaction(types.NewTransactionParam{
		Message: types.NewMessage(types.NewMessageParam{
			FeePayer:        feePayer.PublicKey,
			RecentBlockhash: res.Blockhash,
			Instructions: []types.Instruction{
				token.TransferChecked(token.TransferCheckedParam{
					From:     aliceTokenATAPubkey,
					To:       bobATA,
					Mint:     mintPubKey,
					Auth:     alice.PublicKey,
					Signers:  []common.PublicKey{},
					Amount:   1e5,
					Decimals: 8,
				}),
			},
		}),
		Signers: []types.Account{feePayer, alice},
	})

	if err != nil {
		panic(err)
	}

	txhash, err := c.SendTransaction(context.Background(), tx)
	if err != nil {
		panic(err)
	}
	fmt.Println("txhash:", txhash)
	// gpbdVVSFJnUNTR99p6dyse6ZQoFrJNLPMBmeqdeCGyoThmFevdTu94zJn6aHTpivGkGDxkkJvKVPy6KKyvpPYmt
}

func getBalance() {
	c := client.NewClient(rpc.DevnetRPCEndpoint)

	tokenAmount, err := c.GetTokenAccountBalance(context.Background(), aliceTokenATAPubkey.ToBase58())
	if err != nil {
		panic(err)
	}

	fmt.Println("Balance:", tokenAmount.Amount)
	fmt.Println("Decimals:", tokenAmount.Decimals)
}

func mintTo() {
	c := client.NewClient(rpc.DevnetRPCEndpoint)

	res, err := c.GetLatestBlockhash(context.Background())
	if err != nil {
		panic(err)
	}

	tx, err := types.NewTransaction(types.NewTransactionParam{
		Message: types.NewMessage(types.NewMessageParam{
			FeePayer:        feePayer.PublicKey,
			RecentBlockhash: res.Blockhash,
			Instructions: []types.Instruction{
				token.MintToChecked(token.MintToCheckedParam{
					Mint:     mintPubKey,
					Auth:     alice.PublicKey,
					Signers:  []common.PublicKey{},
					To:       aliceTokenATAPubkey,
					Amount:   1e8,
					Decimals: 8,
				}),
			},
		}),
		Signers: []types.Account{feePayer, alice},
	})

	if err != nil {
		panic(err)
	}

	txhash, err := c.SendTransaction(context.Background(), tx)
	if err != nil {
		panic(err)
	}

	fmt.Println("txhash:", txhash)
	// K9ZVxM2WjdJhMcxHhFKxhF8QhHo86yM4rJm1U6MgGD7PxX9MkfzADZGNLzLGjUwdZqt7jn5dZfTEeNMoYWyzBi8
}

func getTokenAccount() {
	c := client.NewClient(rpc.DevnetRPCEndpoint)

	getAccountInfoResponse, err := c.GetAccountInfo(context.TODO(), "2nrpwrAFwiAxAygUTjaUw9kErAke1nRnpL9iymbdkmsD")
	if err != nil {
		panic(err)
	}
	tokenAccount, err := token.TokenAccountFromData(getAccountInfoResponse.Data)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%+v\n", tokenAccount)
	// {Mint:GeoWLwRr25dUpdDDqyFqcgk7m2MaJgw2VnSJGUfjFaXv Owner:BKHLRxLPiNALhf5UejxqJKjGUWCf5uaWfMP6jxNsi1pM Amount:0 Delegate:<nil> State:1 IsNative:<nil> DelegatedAmount:0 CloseAuthority:<nil>}
}

func createTokenAccount() {
	c := client.NewClient(rpc.DevnetRPCEndpoint)

	ata, _, err := common.FindAssociatedTokenAddress(alice.PublicKey, mintPubKey)
	if err != nil {
		panic(err)
	}
	// 2nrpwrAFwiAxAygUTjaUw9kErAke1nRnpL9iymbdkmsD
	fmt.Println("Associated token account:", ata.ToBase58())

	res, err := c.GetLatestBlockhash(context.Background())
	if err != nil {
		panic(err)
	}

	tx, err := types.NewTransaction(types.NewTransactionParam{
		Message: types.NewMessage(types.NewMessageParam{
			FeePayer:        feePayer.PublicKey,
			RecentBlockhash: res.Blockhash,
			Instructions: []types.Instruction{
				associated_token_account.Create(associated_token_account.CreateParam{
					Funder:                 feePayer.PublicKey,
					Owner:                  alice.PublicKey,
					Mint:                   mintPubKey,
					AssociatedTokenAccount: ata,
				}),
			},
		}),
		Signers: []types.Account{feePayer},
	})

	if err != nil {
		panic(err)
	}

	txhash, err := c.SendTransaction(context.Background(), tx)
	if err != nil {
		panic(err)
	}
	fmt.Println("txhash:", txhash)
	// 4qz3HipCfJhGnjCpiDqRTN9ZDvGJDoHLoxYTW93YRouFRQ3Z1yHLNnGoq1pYa7J1V8ihDo7YqsgQxxtWEung7KGP
}

func getMint() {
	c := client.NewClient(rpc.DevnetRPCEndpoint)
	accountInfoResp, err := c.GetAccountInfo(context.Background(), mintAccount)
	if err != nil {
		panic(err)
	}

	ma, err := token.MintAccountFromData(accountInfoResp.Data)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", ma)
}

func createMint() {
	c := client.NewClient(rpc.DevnetRPCEndpoint)

	// New mint account
	mint := types.NewAccount()
	fmt.Println("Mint account:", mint.PublicKey.ToBase58())

	// get rent
	rentExemptionBalance, err := c.GetMinimumBalanceForRentExemption(context.Background(), token.MintAccountSize)
	if err != nil {
		panic(err)
	}

	// recent block hash
	recentBlockHashResp, err := c.GetLatestBlockhash(context.Background())
	if err != nil {
		panic(err)
	}

	tx, err := types.NewTransaction(types.NewTransactionParam{
		Message: types.NewMessage(types.NewMessageParam{
			FeePayer:        feePayer.PublicKey,
			RecentBlockhash: recentBlockHashResp.Blockhash,
			Instructions: []types.Instruction{
				system.CreateAccount(system.CreateAccountParam{
					From:     feePayer.PublicKey,
					New:      mint.PublicKey,
					Owner:    common.TokenProgramID,
					Lamports: rentExemptionBalance,
					Space:    token.MintAccountSize,
				}),
				token.InitializeMint(token.InitializeMintParam{
					Decimals:   8,
					Mint:       mint.PublicKey,
					MintAuth:   alice.PublicKey,
					FreezeAuth: nil,
				}),
			},
		}),
		Signers: []types.Account{feePayer, mint},
	})

	if err != nil {
		panic(err)
	}

	txhash, err := c.SendTransaction(context.Background(), tx)
	if err != nil {
		panic(err)
	}

	fmt.Println("txhash:", txhash)
	// 2oKoPUuhtLuUikWEcJf2ZKms8GsXihsGXBr8ePejwsK6pU6Pq7ZdQUw3ovKuWeQjEcY8t49PtstPBTkrD6Q7S3pn
	// Mint account: GeoWLwRr25dUpdDDqyFqcgk7m2MaJgw2VnSJGUfjFaXv
}
