package main

import (
	"os"

	"github.com/danubiobwm/blockchain-go/wallet"
)

func main() {
	defer os.Exit(0)
	// cmd := cli.CommandLine{}
	// cmd.Run()

	w := wallet.MakeWallet()
	w.Address()

}
