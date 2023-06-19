/*
 * Copyright © 2023 shaohan.n <shaohan.niu.share@gmail.com>
 */

package main

import (
	"sig-gowallet/cmd/hdkmsctl/cli"
)

// hd-kms cli entrypoint.
func main() {
	// Usage:
	//  hdkms wallet generate --wallet-name master --passphrase 123456
	//  hdkms wallet list
	//  hdkms wallet import
	//  hdkms wallet export --wallet-name master --passphrase 123456
	//  hdkms wallet get [wallet name]
	//  hdkms wallet balance
	//  hdkms address get --chain ETH --index 0

	//  hdkms tx send --data ""
	//

	cli.Execute()
}
