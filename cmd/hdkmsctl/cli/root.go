/*
 * Copyright © 2023 shaohan.n shaohan.niu.share@gmail.com
 */

package cli

import (
	"errors"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

var (
	walletName string
	passphrase string
	cfgPath    string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:              "hdkms [command] [args] [flags]",
	Short:            "Key management for HD wallet.",
	Long:             `hdkms manages hierarchical deterministic wallets from the command line in go.`,
	TraverseChildren: true,
	Run: func(cmd *cobra.Command, args []string) {
		// fmt.Println("rootCmd.Run ...")
		_ = cmd.Help()
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// fmt.Println("rootCmd.init ...")

	// pFlags := rootCmd.PersistentFlags()

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.abc.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	// rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	// rootCmd.Flags().StringVarP(&walletName, "wallet-name", "n", ".kms", "Wallet name, default is `.kms`")
	rootCmd.PersistentFlags().StringVarP(&walletName, "wallet-name", "n", "kms", "Wallet name")
	rootCmd.PersistentFlags().StringVarP(&passphrase, "passphrase", "p", "", "Wallet passphrase, default is empty.")
	rootCmd.PersistentFlags().StringVarP(&cfgPath, "config", "c", "~/.hdkms", "Local config file path")

	initConfigFilePath()
}

// Config file directory structure
func initConfigFilePath() {
	if strings.HasPrefix(cfgPath, "~/") {
		home, err := os.UserHomeDir()
		if err != nil {
			panic(err)
		}

		cfgPath = strings.Replace(cfgPath, "~", home, 1)
	}

	if _, err := os.Stat(cfgPath); errors.Is(err, os.ErrNotExist) {
		err := os.MkdirAll(cfgPath, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
}
