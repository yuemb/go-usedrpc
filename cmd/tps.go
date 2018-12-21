package main

import (
	"fmt"
	"github.com/usechain/go-usedrpc"
)

func main() {
	c := usedrpc.NewUseRPC("http://127.0.0.1:8545")
	blockNumber, _ := c.UseBlockNumber()
	blockStart, _ := c.UseGetBlockByNumber(blockNumber - 30, false)
	blockNow, _ := c.UseGetBlockByNumber(blockNumber, false)
	timeslot := blockNow.Timestamp - blockStart.Timestamp
	var totalTxNum = 0
	for  i := blockStart.Number; i > blockNumber; i++ {
		num, _ := c.UseGetBlockTransactionCountByNumber(i)
		totalTxNum += num
	}
	fmt.Println("The tps :", totalTxNum+timeslot)
}