package internal

import (
	"errors"
	"fmt"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcutil/hdkeychain"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/tyler-smith/go-bip39"
	"sync"
)

var (
	WalletExistsErr = errors.New("wallet exists")
)

// type wallets map[string]*Wallet

var wallets = map[string]*Wallet{}

func loadWallets(path string) {

}

// Wallet is the underlying wallet struct.
type Wallet struct {
	mnemonic   string
	seed       []byte
	passphrase string

	masterKey *hdkeychain.ExtendedKey
	url       accounts.URL
	stateLock sync.RWMutex

	//paths       map[common.Address]accounts.DerivationPath
	//accounts    []accounts.Account
	//fixIssue172 bool
}

// newWallet create a wallet using seed.
func newWallet(seed []byte) (*Wallet, error) {
	masterKey, err := hdkeychain.NewMaster(seed, &chaincfg.MainNetParams)
	if err != nil {
		return nil, err
	}

	return &Wallet{
		seed:      seed,
		masterKey: masterKey,
	}, nil
}

// GenerateWallet is entrypoint for generating HD wallets.
func GenerateWallet(walletName string, passphrase string) error {
	log("Loading available wallets...")
	// TODO 启动 cli 程序时需要加载存储中的 wallets
	// keystore path: ~/.hdkms

	wallets[walletName] = &Wallet{}

	if _, ok := wallets[walletName]; ok {
		return WalletExistsErr
	}

	log("New entropy...")
	var err error
	var entropy []byte
	if entropy, err = bip39.NewEntropy(256); err != nil {
		return err
	}

	log("New mnemonic...")
	var mnemonic string
	if mnemonic, err = bip39.NewMnemonic(entropy); err != nil {
		return err
	}

	log("New seed...")
	seed := bip39.NewSeed(mnemonic, passphrase)

	log("New wallet...")
	var wallet *Wallet
	if wallet, err = newWallet(seed); err != nil {
		return err
	}

	wallet.mnemonic = mnemonic
	wallet.passphrase = passphrase
	wallets[walletName] = wallet

	log("Store key...")
	// TODO 要把 wallet 持久化才算成功

	return nil
}

func log(a ...any) {
	fmt.Println(a)
}
