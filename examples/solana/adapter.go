package solana

import (
	"crypto/ed25519"
	"encoding/base64"
	"encoding/hex"
	"github.com/blocto/solana-go-sdk/client"
	"github.com/btcsuite/btcutil/base58"
)

// PubKey represents a public key, which can be converted to an address.
type PubKey interface {
	// Address returns the address of the public key.
	Address() (string, error)

	// Hex() string
	// Bytes() []byte
	// Marshal() ([]byte, error)
	// Unmarshal([]byte) error
}

const PublicKeyLength = 32

// SolPubKey represents a Solana public key.
type SolPubKey [PublicKeyLength]byte

func (p SolPubKey) Address() (string, error) {
	return base58.Encode(p[:]), nil
}

// SolPubKeyFromBytes creates a SolPubKey from bytes.
func SolPubKeyFromBytes(b []byte) PubKey {
	var pubKey SolPubKey
	// TODO: More checks
	copy(pubKey[:], b)
	return pubKey
}

// SolPubKeyFromString creates a SolPubKey from a base58-encoded string.
func SolPubKeyFromString(s string) PubKey {
	return SolPubKeyFromBytes(base58.Decode(s))
}

// EthPubKey represents an Ethereum public key.
type EthPubKey []byte

func (p EthPubKey) Address() (string, error) {
	// TODO: Implement
	return "", nil
}

// EthPubKeyFromBytes creates an EthPubKey from bytes.
func EthPubKeyFromBytes(b []byte) PubKey {
	var pubKey EthPubKey
	copy(pubKey, b)
	return pubKey
}

type Keypair interface {
	// GenerateKey returns the master public key of the keypair. The param curve is the curve
	// used for the public key, e.g. ed25519, secp256k1.
	GenerateKey(curve string)

	// DeriveChildKey returns the public key of the derived child keypair.
	// m/44'/60'/0'/0/0
	DeriveChildKey(masterPubKey string, childIndex uint32) PubKey
}

type LocalKeypair struct {
	pubKey PubKey // Generated master public key

	curve   string // e.g. ed25519, secp256k1 ...
	network string // e.g. ETH, SOL ...
}

func (k *LocalKeypair) GenerateKey(curve string) {
	// TODO More ...
	_, priKey, _ := ed25519.GenerateKey(nil)
	pubKey := priKey.Public().(ed25519.PublicKey)

	var _pk SolPubKey
	copy(_pk[PublicKeyLength-len(pubKey):], pubKey)

	k.pubKey = _pk
}

type MpcKeypair struct {
	pubKey PubKey
}

/// Broadcast Tx

type RawTx struct {
	chain string // e.g. ETH, SOL ...

	message string // Unsigned transaction message, hex-encoded or base64-encoded

	innerTx interface{} // innerTx represents the inner transaction, which associates with the chain itself.
}

// SignedTx represents a signed transaction, which can be broadcast to the chain.
type SignedTx struct {
	message  string // Fully-signed transaction message, hex-encoded or base64-encoded
	encoding string // hex, base64 ...

	chain       string // e.g. ETH, SOL ...
	forkedChain string // Which chain forked from
}

func (tx SignedTx) Bytes() []byte {
	switch tx.encoding {
	case "base64":
		b, _ := base64.StdEncoding.DecodeString(tx.message)
		return b
	case "hex":
		b, _ := hex.DecodeString(tx.message)
		return b
	default:
		return []byte(tx.message)
	}
}

type TxReceipt struct {
	Status    string // success, failed
	TxHash    string
	BlockHash string
	Height    string
}

type TxBuilder interface {
	Build() RawTx
}

type TxSigner interface {
	Sign(tx RawTx) (SignedTx, error)
}

type TxBroadcaster interface {
	Broadcast(tx SignedTx) error
}

type ChainAdapterClient interface {
	TxBuilder
	TxSigner
	TxBroadcaster

	GetTx(txHash string) TxReceipt
}

type SolanaAdapter struct {
	chain string // SOL

	client *client.Client // RPC Client

	// mpcClient *mpc.Client // MPC Client
}

func (a SolanaAdapter) Build() RawTx {
	// TODO
	return RawTx{chain: a.chain, message: ""}
}

func Sign() (SignedTx, error) {
	// TODO
	return SignedTx{}, nil
}
