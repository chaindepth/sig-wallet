package bip32

import (
	"encoding/hex"
	"fmt"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/stretchr/testify/assert"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
	"golang.org/x/crypto/sha3"
	"strings"
	"testing"
)

func TestGenerateMasterKey(t *testing.T) {
	mnemonic := "range sheriff try enroll deer over ten level bring display stamp recycle"
	seed, _ := bip39.NewSeedWithErrorChecking(mnemonic, "")

	masterKey, err := bip32.NewMasterKey(seed)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("masterKey: %s", masterKey.PublicKey())

	childKey0, _ := masterKey.NewChildKey(0)
	t.Logf("childKey0: %s", childKey0.PublicKey())

	childKey1, _ := masterKey.NewChildKey(1)
	t.Logf("childKey1: %s", childKey1.PublicKey())

	childKey2, _ := masterKey.NewChildKey(2)
	t.Logf("childKey2: %s", childKey2.PublicKey())
}

func TestBip32Library(t *testing.T) {
	mnemonic := "tag volcano eight thank tide danger coast health above argue embrace heavy"
	seed, _ := bip39.NewSeedWithErrorChecking(mnemonic, "")

	masterKey, err := bip32.NewMasterKey(seed)
	if err != nil {
		t.Fatal(err)
	}

	// m/44'/60'/0'/0/0
	childKey0, _ := masterKey.NewChildKey(bip32.FirstHardenedChild + 44) // 44'
	t.Logf("-> %d", bip32.FirstHardenedChild+44)
	t.Logf("childKey0: %s", childKey0.PublicKey())

	childKey1, _ := childKey0.NewChildKey(bip32.FirstHardenedChild + 60) // 60'
	t.Logf("-> %d", bip32.FirstHardenedChild+60)
	childKey2, _ := childKey1.NewChildKey(bip32.FirstHardenedChild + 0) // 0'
	t.Logf("-> %d", bip32.FirstHardenedChild+0)
	childKey3, _ := childKey2.NewChildKey(0) // 0
	t.Logf("-> %d", 0)
	childKey4, _ := childKey3.NewChildKey(1) // 0
	t.Logf("-> %d", 0)

	t.Log(childKey4.Key)
	privateKey, address := encodeEthereum(childKey4.Key)
	t.Logf("privateKey: %s", privateKey)
	t.Logf("address: %s", address) // 0xc49926c4124cee1cba0ea94ea31a6c12318df947
	// assert.True(t, strings.EqualFold(address, "0xc49926c4124cee1cba0ea94ea31a6c12318df947"))
	assert.True(t, strings.EqualFold(address, "0x8230645ac28a4edd1b0b53e7cd8019744e9dd559"))
}

// encodeEthereum encodes the private key and address for Ethereum.
func encodeEthereum(privateKeyBytes []byte) (privateKey, address string) {
	_, pubKey := btcec.PrivKeyFromBytes(privateKeyBytes)

	publicKey := pubKey.ToECDSA()
	publicKeyBytes := append(publicKey.X.Bytes(), publicKey.Y.Bytes()...)

	// Ethereum uses the last 20 bytes of the keccak256 hash of the public key
	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKeyBytes)
	addr := hash.Sum(nil)
	addr = addr[len(addr)-20:]

	return hex.EncodeToString(privateKeyBytes), eip55checksum(fmt.Sprintf("0x%x", addr))
}

// eip55checksum implements the EIP55 checksum address encoding.
// this function is copied from the go-ethereum library: go-ethereum/common/types.go checksumHex method
func eip55checksum(address string) string {
	buf := []byte(address)
	sha := sha3.NewLegacyKeccak256()
	sha.Write(buf[2:])
	hash := sha.Sum(nil)
	for i := 2; i < len(buf); i++ {
		hashByte := hash[(i-2)/2]
		if i%2 == 0 {
			hashByte = hashByte >> 4
		} else {
			hashByte &= 0xf
		}
		if buf[i] > '9' && hashByte > 7 {
			buf[i] -= 32
		}
	}
	return string(buf[:])
}
