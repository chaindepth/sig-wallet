package main

import (
	"encoding/base64"
	"testing"
)

func TestSystemProgram(t *testing.T) {
	b, _ := base64.StdEncoding.DecodeString("c29sYW5hX3N5c3RlbV9wcm9ncmFt")
	t.Log(string(b))
}

func TestTransferSolWithFeePayer(t *testing.T) {
	transferSolWithFeePayer()
}
