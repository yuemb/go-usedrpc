package main

import (
	"math/big"
	"fmt"
	"github.com/usechain/go-usedrpc"
	"time"
)

func main() {
	c := usedrpc.NewUseRPC("http://127.0.0.1:8545")
	accounts, _ := c.UseAccounts()
	tx := usedrpc.T {
		From: accounts[0],
		To:   "0xfffffffffffffffffffffffffffffffff0000001",
		Value: big.NewInt(0),
		Data:  "",
	}
	for {
		flag, err := c.UseSendTransaction(tx)
		fmt.Println("The tx hash:", err, flag)
		time.Sleep(50*time.Millisecond)
	}
}



