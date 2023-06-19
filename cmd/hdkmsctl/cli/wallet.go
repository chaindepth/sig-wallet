/**
 * Copyright © 2023 shaohan.n shaohan.niu.share@gmail.com
 */

package cli

import (
	"fmt"
	impl "sig-gowallet/internal"

	"github.com/spf13/cobra"
)

func init() {
	// fmt.Println("wallet.init ...")
	rootCmd.AddCommand(walletCmd)

	walletCmd.AddCommand(generateCmd)
	walletCmd.AddCommand(listCmd)
	walletCmd.AddCommand(importCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// generateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// generateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// walletCmd represents the wallet command, which related to wallet management, and
// subcommands as below: generate, list, import, export, balance ...
var walletCmd = &cobra.Command{
	Use:              "wallet [command]",
	Short:            "wallet subcommand",
	Long:             `All functions related to wallet management, e.g. generate, list, import, export etc.`,
	TraverseChildren: true,
}

// generateCmd represents the generate wallet command
var generateCmd = &cobra.Command{
	Use:   "generate [flags]",
	Short: "Generate a new HD wallet.",
	Long:  `Generate a new HD wallet, and you can specify the wallet name, wallet passphrase, etc..`,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("generate wallet called")
		// fmt.Println("Wallet name:", walletName)
		// fmt.Println("Passphrase:", passphrase)
		//
		// fmt.Println(cmd.Use, args)

		// check arguments
		err := impl.GenerateWallet(walletName, passphrase)
		if err == nil {
			fmt.Println("Wallet generated:", walletName)
		}
	},
	TraverseChildren: true,
}

// listCmd represents the list wallet command
var listCmd = &cobra.Command{
	Use:              "list [flags]",
	Short:            "List wallets.",
	Long:             `List all available wallets`,
	TraverseChildren: true,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO call listWallets
	},
}

// importCmd represents the list wallet command
var importCmd = &cobra.Command{
	Use:              "import [flags]",
	Short:            "Import a mnemonic or private key.",
	Long:             `Import a mnemonic or private key to create a wallet.`,
	TraverseChildren: true,
	Run: func(cmd *cobra.Command, args []string) {
		// TODO call import mnemonic or private key
	},
}
