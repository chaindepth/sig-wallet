package main

import (
	"context"
	"encoding/hex"
	"fmt"
	"github.com/blocto/solana-go-sdk/client"
	"github.com/blocto/solana-go-sdk/common"
	"github.com/blocto/solana-go-sdk/pkg/hdwallet"
	"github.com/blocto/solana-go-sdk/program/system"
	"github.com/blocto/solana-go-sdk/rpc"
	"github.com/blocto/solana-go-sdk/types"
	"github.com/tyler-smith/go-bip39"
)

// TestHexPrivateKey its address is BKHLRxLPiNALhf5UejxqJKjGUWCf5uaWfMP6jxNsi1pM
const TestHexPrivateKey = "df6db68d2c28db792f729d459ba62bca536c0ec9b0eb433958880a72f8b2578c9943cb3b63f80ffd41f6284e29ce2f30048aa1a429f0157577325054317b5022"

// TestBytesSlicePrivateKey its address is DsqkS4YYodw1PPeUZTgaaYrpZ5wythns9QmFBV3VnjiC
var TestBytesSlicePrivateKey = []byte{100, 25, 253, 115, 217, 126, 60, 157, 59, 25, 74, 187, 227, 173, 85, 139, 84, 2, 43, 150, 46, 178, 115, 65, 160, 171, 187, 52, 73, 29, 70, 93, 191, 82, 103, 111, 156, 1, 99, 114, 80, 18, 41, 78, 202, 246, 46, 3, 177, 153, 169, 56, 191, 92, 165, 116, 114, 89, 250, 160, 135, 146, 221, 69}
var feePayer, _ = types.AccountFromBytes(TestBytesSlicePrivateKey)

// FromWalletPrivateKey 2PKXDnSYQf8jbMMGi66V7optX3UfVaqSiRkoPQACDDZD
var FromWalletPrivateKey = []byte{173, 127, 250, 43, 227, 134, 117, 153, 53, 211, 73, 215, 67, 137, 82, 83, 64, 112, 105, 208, 11, 135, 10, 187, 130, 81, 42, 73, 68, 236, 204, 221, 20, 147, 104, 18, 162, 111, 44, 226, 159, 100, 229, 207, 202, 60, 124, 64, 215, 254, 38, 153, 167, 197, 164, 222, 109, 184, 214, 240, 89, 245, 203, 236}
var fromWallet, _ = types.AccountFromBytes(FromWalletPrivateKey)

const targetAddress = "B3HSvqGvWVCppZh6A7g7ga94eZS4DRFQNFTPi65f8snc"

// const solReceiver = "2PKXDnSYQf8jbMMGi66V7optX3UfVaqSiRkoPQACDDZD"

func main() {
	// getVersion()
	// generateWallet()
	// importWallet()
	// requestAirdrop()
	// getBalance()
	// transferSol()
	getTransaction("313wTtj2kGPKWNCUiAxwL57MXQks1FFTNyTtoWLmmvPyC6JREFy7vu7rQ87FjRFptMZMtoenwfSyoRwcV5VuXypL")
	// getBlockAndParseTransactions()
	// issueToken()
}

func issueToken() {
	// Create mint.

}

func getTransaction(txHash string) {
	c := client.NewClient(rpc.DevnetRPCEndpoint)

	tx, err := c.GetTransaction(context.Background(), txHash)
	if err != nil {
		panic(err)
	}

	fmt.Println("Transaction hash:", tx.Transaction.Signatures[0])
	fmt.Println("Transaction message:", tx.Transaction.Message)
}

func getBlockAndParseTransactions() {
	c := client.NewClient(rpc.DevnetRPCEndpoint)

	slot, err := c.GetSlot(context.Background())
	if err != nil {
		panic(err)
	}
	fmt.Println("Current slot:", slot)

	block, err := c.GetBlock(context.TODO(), slot)
	if err != nil {
		panic(err)
	}

	fmt.Println("Block hash:", block.Blockhash)
	fmt.Println("Block time:", block.BlockTime)
	fmt.Println("Block height:", block.BlockHeight)
	fmt.Println("Previous block hash:", block.PreviousBlockhash)
	fmt.Println("Parent slot:", block.ParentSlot)
	fmt.Println("Block transactions:", len(block.Transactions))

	// for i := range block.Transactions {
	// blockTx := block.Transactions[i]
	// }
}

func transferSolWithFeePayer() {
	c := client.NewClient(rpc.DevnetRPCEndpoint)
	response, err := c.GetLatestBlockhash(context.TODO())
	if err != nil {
		panic(err)
	}

	rent, err := c.GetMinimumBalanceForRentExemption(context.TODO(), 0)
	if err != nil {
		panic(err)
	}

	fmt.Println("Normal account minimum rent:", rent)

	// wallet, _ := types.AccountFromBytes(TestBytesSlicePrivateKey)
	a := types.NewAccount()
	fmt.Println("New wallet address:", a.PublicKey.ToBase58())

	tx, err := types.NewTransaction(types.NewTransactionParam{
		Signers: []types.Account{feePayer},
		Message: types.NewMessage(types.NewMessageParam{
			FeePayer:        feePayer.PublicKey,
			RecentBlockhash: response.Blockhash,
			Instructions: []types.Instruction{
				system.Transfer(system.TransferParam{
					From: feePayer.PublicKey,
					// To:   a.PublicKey,
					To: common.PublicKeyFromString(targetAddress),
					// Amount: 1e6, // 0.001 SOL
					// Amount: rent,
					Amount: 1,
				}),
			},
		}),
	})
	if err != nil {
		panic(err)
	}
	txhash, err := c.SendTransaction(context.Background(), tx)
	if err != nil {
		panic(err)
	}
	fmt.Println("Transaction hash:", txhash)
}

func transferSol() {
	c := client.NewClient(rpc.DevnetRPCEndpoint)
	response, err := c.GetLatestBlockhash(context.TODO())
	if err != nil {
		panic(err)
	}

	wallet, _ := types.AccountFromBytes(TestBytesSlicePrivateKey)

	tx, err := types.NewTransaction(types.NewTransactionParam{
		Signers: []types.Account{wallet},
		Message: types.NewMessage(types.NewMessageParam{
			FeePayer:        wallet.PublicKey,
			RecentBlockhash: response.Blockhash,
			Instructions: []types.Instruction{
				system.Transfer(system.TransferParam{
					From:   wallet.PublicKey,
					To:     common.PublicKeyFromString(targetAddress),
					Amount: 1e6, // 0.001 SOL
				}),
			},
		}),
	})
	if err != nil {
		panic(err)
	}
	txhash, err := c.SendTransaction(context.Background(), tx)
	if err != nil {
		panic(err)
	}
	fmt.Println("Transaction hash:", txhash)
	// 4thPdnk9Zu7xGYA3SKpVuRGV36TdWpKp8pT1W1EMVz82H5GWfjgYwHTjaBxyrtqGBJyMj1RmgkhMga5a8tKzNLuu

	// 4AXFHJM4Sqe8giv1hSRM4zt3gSeFE2wBtqHRtNZzuYhJvTZ2hjMwNB9DvFRBVzcUcsbd3FkgBmEbQXbmQ2imSeNG
	// 313wTtj2kGPKWNCUiAxwL57MXQks1FFTNyTtoWLmmvPyC6JREFy7vu7rQ87FjRFptMZMtoenwfSyoRwcV5VuXypL
}

func getBalance() {
	c := client.NewClient(rpc.DevnetRPCEndpoint)

	balance, err := c.GetBalance(
		context.TODO(), // request context
		"DsqkS4YYodw1PPeUZTgaaYrpZ5wythns9QmFBV3VnjiC", // wallet to fetch balance for
	)
	if err != nil {
		panic(err)
	}

	fmt.Println("Wallet Balance in Lamport:", balance)
	fmt.Println("Wallet Balance in SOL:", float64(balance)/1e9)
}

// Only works on devnet.
func requestAirdrop() {
	c := client.NewClient(rpc.DevnetRPCEndpoint)

	// account, err := types.AccountFromHex(TestHexPrivateKey)
	account, err := types.AccountFromBytes(TestBytesSlicePrivateKey)
	if err != nil {
		panic(err)
	}

	fmt.Println("Account address:", account.PublicKey.ToBase58())

	txHash, err := c.RequestAirdrop(context.TODO(), account.PublicKey.ToBase58(), 100)
	if err != nil {
		panic(err)
	}

	fmt.Println("Airdrop transaction hash:", txHash)
}

func importWallet() {
	// Import a wallet with base58 private key.
	account, err := types.AccountFromBase58("28WJTTqMuurAfz6yqeTrFMXeFd91uzi9i1AW6F5KyHQDS9siXb8TquAuatvLuCEYdggyeiNKLAUr3w7Czmmf2Rav")
	if err != nil {
		panic(err)
	}

	fmt.Println("Imported base58 wallet address:", account.PublicKey.ToBase58())

	// Import a wallet with bytes slice private key.
	account, err = types.AccountFromBytes([]byte{
		56, 125, 59, 118, 230, 173, 152, 169, 197, 34,
		168, 187, 217, 160, 119, 204, 124, 69, 52, 136,
		214, 49, 207, 234, 79, 70, 83, 224, 1, 224, 36,
		247, 131, 83, 164, 85, 139, 215, 183, 148, 79,
		198, 74, 93, 156, 157, 208, 99, 221, 127, 51,
		156, 43, 196, 101, 144, 104, 252, 221, 108,
		245, 104, 13, 151,
	})
	fmt.Println("Imported bytes slice wallet address:", account.PublicKey.ToBase58())

	// Import a wallet with hex string private key.
	account, err = types.AccountFromHex("387d3b76e6ad98a9c522a8bbd9a077cc7c453488d631cfea4f4653e001e024f78353a4558bd7b794fc64a5d9c9dd063dd7f339c2bc4659068fcdd6cf5680d97")
	fmt.Println("Imported hex string wallet address:", account.PublicKey.ToBase58())

	// bip39 mnemonic
	mnemonic := "pill tomorrow foster begin walnut borrow virtual kick shift mutual shoe scatter"
	seed := bip39.NewSeed(mnemonic, "")
	account, err = types.AccountFromSeed(seed[:32])
	fmt.Println("Imported mnemonic wallet address:", account.PublicKey.ToBase58())

	// bip44 mnemonic
	mnemonic = "neither lonely flavor argue grass remind eye tag avocado spot unusual intact"
	seed = bip39.NewSeed(mnemonic, "") // (mnemonic, password)
	path := `m/44'/501'/0'/0'`
	derivedKey, _ := hdwallet.Derived(path, seed)
	account, _ = types.AccountFromSeed(derivedKey.PrivateKey)
	fmt.Printf("%v => %v\n", path, account.PublicKey.ToBase58())

	// others
	for i := 1; i < 10; i++ {
		// path := fmt.Sprintf(`m/44'/501'/%d'/0'`, i)
		path := fmt.Sprintf(`m/44'/501'/0'/%d'`, i)
		derivedKey, _ := hdwallet.Derived(path, seed)
		account, _ := types.AccountFromSeed(derivedKey.PrivateKey)
		fmt.Printf("%v => %v\n", path, account.PublicKey.ToBase58())
	}
}

func generateWallet() {
	// Create a new wallet using types.NewWallet()
	account := types.NewAccount()
	fmt.Println("New wallet address:", account.PublicKey.ToBase58())
	fmt.Println("New wallet secret key:", account.PrivateKey)
	fmt.Println(hex.EncodeToString(account.PrivateKey))
}

func getVersion() {
	// Create a rpc client.
	c := client.NewClient(rpc.DevnetRPCEndpoint)

	response, err := c.GetVersion(context.TODO())
	if err != nil {
		panic(err)
	}

	fmt.Println("Solana devnet RPC version:", response.SolanaCore)
}
