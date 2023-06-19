package main

import (
	"fmt"
	"github.com/shniu/gokits/blockchain"
	"testing"
)

func TestVersion(t *testing.T) {
	fmt.Println(blockchain.Version())
}
