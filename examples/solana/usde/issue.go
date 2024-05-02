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
	"time"
)

const LAMPORTS_PER_SOL = 1000000000

// feePayerPrivateKeyBytes its address is DsqkS4YYodw1PPeUZTgaaYrpZ5wythns9QmFBV3VnjiC
var feePayerPrivateKeyBytes = []byte{
	100, 25, 253, 115, 217, 126, 60, 157, 59, 25, 74, 187, 227, 173,
	85, 139, 84, 2, 43, 150, 46, 178, 115, 65, 160, 171, 187, 52, 73,
	29, 70, 93, 191, 82, 103, 111, 156, 1, 99, 114, 80, 18, 41, 78,
	202, 246, 46, 3, 177, 153, 169, 56, 191, 92, 165, 116, 114, 89,
	250, 160, 135, 146, 221, 69,
}

// mintAuthorityPrivateKeyBytes its address is 2PKXDnSYQf8jbMMGi66V7optX3UfVaqSiRkoPQACDDZD
var mintAuthorityPrivateKeyBytes = []byte{
	173, 127, 250, 43, 227, 134, 117, 153, 53, 211, 73, 215, 67, 137, 82, 83, 64, 112, 105,
	208, 11, 135, 10, 187, 130, 81, 42, 73, 68, 236, 204, 221, 20, 147, 104, 18, 162, 111, 44,
	226, 159, 100, 229, 207, 202, 60, 124, 64, 215, 254, 38, 153, 167, 197, 164, 222, 109,
	184, 214, 240, 89, 245, 203, 236,
}

// destinationPrivateKeyBytes its address is B3HSvqGvWVCppZh6A7g7ga94eZS4DRFQNFTPi65f8snc
var destinationPrivateKeyBytes = []byte{70, 162, 23, 39, 194, 223, 15, 181, 62, 254, 105, 41,
	136, 34, 182, 93, 20, 206, 5, 171, 123, 229, 201, 51, 34, 56, 173, 14, 232, 183, 95, 47,
	149, 42, 160, 170, 195, 129, 53, 126, 255, 215, 134, 176, 156, 102, 183, 236, 233, 56, 74,
	132, 105, 201, 225, 79, 6, 128, 4, 19, 37, 77, 177, 21,
}

func getFeePayerAccount() types.Account {
	feePayer, _ := types.AccountFromBytes(feePayerPrivateKeyBytes)
	return feePayer
}

func getFreezeAuthorityAccount() types.Account {
	freezeAuthority, _ := types.AccountFromBytes(mintAuthorityPrivateKeyBytes)
	return freezeAuthority
}

func getDestinationAccount() types.Account {
	destination, _ := types.AccountFromBytes(destinationPrivateKeyBytes)
	return destination
}

func getMintAuthorityAccount() types.Account {
	mintAuthority, _ := types.AccountFromBytes(mintAuthorityPrivateKeyBytes)
	return mintAuthority
}

func getBalance(c *client.Client, account types.Account) uint64 {
	lamports, _ := c.GetBalance(context.Background(), account.PublicKey.ToBase58())
	return lamports
}

func createTokenMint(c *client.Client, tokenMintAccount types.Account, decimals uint8) string {
	rentExemptionBalance, err := c.GetMinimumBalanceForRentExemption(context.Background(), token.MintAccountSize)
	if err != nil {
		panic(err)
	}

	recentBlockHashResp, err := c.GetLatestBlockhash(context.Background())
	if err != nil {
		panic(err)
	}

	feePayerAccount := getFeePayerAccount()
	mintAuthorityAccount := getMintAuthorityAccount()

	tx, err := types.NewTransaction(types.NewTransactionParam{
		Message: types.NewMessage(types.NewMessageParam{
			FeePayer:        feePayerAccount.PublicKey,
			RecentBlockhash: recentBlockHashResp.Blockhash,
			Instructions: []types.Instruction{
				// Create mint account instruction
				system.CreateAccount(system.CreateAccountParam{
					From:     feePayerAccount.PublicKey,
					New:      tokenMintAccount.PublicKey,
					Owner:    common.TokenProgramID,
					Lamports: rentExemptionBalance,
					Space:    token.MintAccountSize,
				}),
				// Initialize mint account instruction
				token.InitializeMint(token.InitializeMintParam{
					Decimals:   decimals,
					Mint:       tokenMintAccount.PublicKey,
					MintAuth:   mintAuthorityAccount.PublicKey,
					FreezeAuth: &mintAuthorityAccount.PublicKey,
				}),
			},
		}),
		Signers: []types.Account{feePayerAccount, tokenMintAccount},
	})

	if err != nil {
		panic(err)
	}

	fmt.Println("Creating token mint:", tokenMintAccount.PublicKey.ToBase58())
	return sendTransaction(c, tx)
}

func getTokenMint(c *client.Client, mintPubKey string) {
	accountInfo, err := c.GetAccountInfo(context.TODO(), mintPubKey)
	if err != nil {
		panic(err)
	}

	mintAccount, err := token.MintAccountFromData(accountInfo.Data)
	if err != nil {
		panic(err)
	}
	fmt.Println("√ Got token mint account from chain:", mintAccount)
}

func createTokenAccount(c *client.Client, walletAddress, tokenMintAccount types.Account) common.PublicKey {
	// mintAuthorityAccount := getMintAuthorityAccount()
	feePayerAccount := getFeePayerAccount()

	ataPubKey, nonce, err := common.FindAssociatedTokenAddress(walletAddress.PublicKey, tokenMintAccount.PublicKey)
	if err != nil {
		panic(err)
	}

	fmt.Println("Associated token account:", ataPubKey.ToBase58(), nonce)

	recentBlockHashResp, err := c.GetLatestBlockhash(context.Background())
	if err != nil {
		panic(err)
	}

	tx, err := types.NewTransaction(types.NewTransactionParam{
		Message: types.NewMessage(types.NewMessageParam{
			FeePayer:        feePayerAccount.PublicKey,
			RecentBlockhash: recentBlockHashResp.Blockhash,
			Instructions: []types.Instruction{
				associated_token_account.Create(associated_token_account.CreateParam{
					Funder:                 feePayerAccount.PublicKey,
					Owner:                  walletAddress.PublicKey,
					Mint:                   tokenMintAccount.PublicKey,
					AssociatedTokenAccount: ataPubKey,
				}),
			},
		}),
		Signers: []types.Account{feePayerAccount},
	})
	if err != nil {
		panic(err)
	}

	_ = sendTransaction(c, tx)
	return ataPubKey
}

func getTokenAccount(c *client.Client, ata common.PublicKey) {
	accountInfo, err := c.GetAccountInfo(context.TODO(), ata.ToBase58())
	if err != nil {
		panic(err)
	}

	tokenAccount, err := token.TokenAccountFromData(accountInfo.Data)
	if err != nil {
		panic(err)
	}
	fmt.Println("√ Got token account from chain:", tokenAccount)
}

func mintTo(c *client.Client, tokenMintAccount types.Account, ata common.PublicKey, supply uint64, decimals uint8) {
	feePayerAccount := getFeePayerAccount()
	mintAuthorityAccount := getMintAuthorityAccount()

	recentBlockHashResp, err := c.GetLatestBlockhash(context.Background())
	if err != nil {
		panic(err)
	}

	tx, err := types.NewTransaction(types.NewTransactionParam{
		Message: types.NewMessage(types.NewMessageParam{
			FeePayer:        feePayerAccount.PublicKey,
			RecentBlockhash: recentBlockHashResp.Blockhash,
			Instructions: []types.Instruction{
				token.MintToChecked(token.MintToCheckedParam{
					Mint:     tokenMintAccount.PublicKey,
					Auth:     mintAuthorityAccount.PublicKey,
					Signers:  []common.PublicKey{},
					To:       ata,
					Amount:   supply,
					Decimals: decimals,
				}),
			},
		}),
		Signers: []types.Account{feePayerAccount, mintAuthorityAccount},
	})

	if err != nil {
		panic(err)
	}

	sendTransaction(c, tx)
}

func getAtaBalance(c *client.Client, tokenMintAccount types.Account, ata common.PublicKey) {
	tokenAmount, err := c.GetTokenAccountBalance(context.Background(), ata.ToBase58())
	if err != nil {
		panic(err)
	}

	fmt.Println("Balance:", tokenAmount.Amount)
	fmt.Println("Decimals:", tokenAmount.Decimals)
	fmt.Println("UI Amount String:", tokenAmount.UIAmountString)
}

func transferToken(c *client.Client, tokenMintAccount, destinationAccount types.Account, ata common.PublicKey, amount uint64, decimals uint8) {
	// destinationAccount := getDestinationAccount()
	feePayerAccount := getFeePayerAccount()
	mintAuthAccount := getMintAuthorityAccount()

	destinationAta := createTokenAccount(c, destinationAccount, tokenMintAccount)
	fmt.Println("Destination ATA created:", destinationAta.ToBase58())
	time.Sleep(10 * time.Second)

	getTokenAccount(c, destinationAta)

	time.Sleep(10 * time.Second)

	doTransfer(c, destinationAccount, mintAuthAccount, feePayerAccount, tokenMintAccount.PublicKey, ata, destinationAta, amount, decimals)
}

func doTransfer(c *client.Client, destinationAccount, mintAuthAccount, feePayerAccount types.Account, tokenMintPubKye, ata, destinationAta common.PublicKey, amount uint64, decimals uint8) {
	recentBlockHashResp, err := c.GetLatestBlockhash(context.Background())
	if err != nil {
		panic(err)
	}

	tx, err := types.NewTransaction(types.NewTransactionParam{
		Message: types.NewMessage(types.NewMessageParam{
			FeePayer:        feePayerAccount.PublicKey,
			RecentBlockhash: recentBlockHashResp.Blockhash,
			Instructions: []types.Instruction{
				token.TransferChecked(token.TransferCheckedParam{
					From:     ata,
					To:       destinationAta,
					Mint:     tokenMintPubKye,
					Auth:     mintAuthAccount.PublicKey,
					Signers:  []common.PublicKey{},
					Amount:   amount,
					Decimals: decimals,
				}),
			},
		}),
		Signers: []types.Account{feePayerAccount, mintAuthAccount},
	})
	if err != nil {
		panic(err)
	}

	fmt.Println("Transferring token to:", destinationAccount.PublicKey.ToBase58(), "of its ATA:", destinationAta.ToBase58())

	sendTransaction(c, tx)
}

func sendTransaction(c *client.Client, tx types.Transaction) string {
	txHash, err := c.SendTransaction(context.Background(), tx)
	if err != nil {
		panic(err)
	}
	fmt.Println("Transaction hash:", txHash)
	return txHash
}

func main() {
	fmt.Println("------ Issuing a SPL Token Start ------")
	c := client.NewClient(rpc.DevnetRPCEndpoint)

	var decimals = uint8(8)
	var supply = uint64(10000000000000000)

	feePayerAccount := getFeePayerAccount()
	lamportsBalance := getBalance(c, feePayerAccount)
	fmt.Println("Fee payer is:", feePayerAccount.PublicKey.ToBase58(), "and its balance:", float64(lamportsBalance)/LAMPORTS_PER_SOL, "SOL")

	// 1. Create token mint account.
	fmt.Println("\n>> Step 1: Create token mint account")
	tokenMintAccount := types.NewAccount()
	createTokenMint(c, tokenMintAccount, decimals)

	time.Sleep(15 * time.Second)

	// 2. Double check if the token mint account is created successfully.
	fmt.Println("\n>> Step 2: Get token mint account")
	getTokenMint(c, tokenMintAccount.PublicKey.ToBase58())

	// 3. Create token account, we use ATA (Associated Token Account) here.
	fmt.Println("\n>> Step 3: Create token account")
	mintAuthorityAccount := getMintAuthorityAccount()
	ata := createTokenAccount(c, mintAuthorityAccount, tokenMintAccount)

	time.Sleep(15 * time.Second)

	// 4. Double check if the token account is created successfully.
	fmt.Println("\n>> Step 4: Get token account")
	getTokenAccount(c, ata)

	// 5. Mint to the token account.
	fmt.Println("\n>> Step 5: Mint to the token account")
	mintTo(c, tokenMintAccount, ata, supply, decimals)

	time.Sleep(15 * time.Second)

	// 6. Query the token account balance.
	fmt.Println("\n>> Step 6: Get token account balance")
	getAtaBalance(c, tokenMintAccount, ata)

	time.Sleep(15 * time.Second)

	// 7. Transfer token to another account.
	fmt.Println("\n>> Step 7: Transfer token to another account")
	transferToken(c, tokenMintAccount, getDestinationAccount(), ata, 1000000000, decimals)

	fmt.Println("------ Issuing a SPL Token End ------")
}
