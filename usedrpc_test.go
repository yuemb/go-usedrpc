// Copyright 2018 The go-usedrpc Authors
// This file is part of the go-usedrpc library.
//
// The go-usedrpc library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-usedrpc library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-usedrpc library. If not, see <http://www.gnu.org/licenses/>.

package usedrpc

import (
	"fmt"
	"testing"
	"math/big"
)

func TestRPCClient(t *testing.T) {
	c := NewUseRPC("http://47.100.195.101:8545")
	coinbase, err := c.UseCoinbase()
	fmt.Println("The block number:", err, coinbase)
}

func TestRPCQueryAddress(t *testing.T) {
	c := NewUseRPC("http://127.0.0.1:8545")
	flag, err := c.UseQueryAddress("0x6da8c30181d22e69fb17fa498f5cba5b09ecd572", "latest")
	fmt.Println("The address auentication stat is", err, flag)
}

func TestRPCMiner(t *testing.T) {
	c := NewUseRPC("http://127.0.0.1:8545")
	flag, err := c.UseIsMiner("0x6da8c30181d22e69fb17fa498f5cba5b09ecd572", "latest")
	fmt.Println("Is a miner", err, flag)
}

func TestRPCMainTransaction(t *testing.T) {
	c := NewUseRPC("http://127.0.0.1:8545")
	tx := T {
		From: "0x6da8c30181d22e69fb17fa498f5cba5b09ecd572",
		To:   "0xfffffffffffffffffffffffffffffffff0000001",
		Value: big.NewInt(0),
		Data:  "",
	}
	flag, err := c.UseSendSubTransaction(tx, "0x6da8c30181d22e69fb17fa498f5cba5b09ecd572", "latest")
	fmt.Println("The tx hash:", err, flag)
}

func TestRPCTransaction(t *testing.T) {
	c := NewUseRPC("http://127.0.0.1:8545")
	accounts, _ := c.UseAccounts()
	tx := T {
		From: accounts[0],
		To:   "0xfffffffffffffffffffffffffffffffff0000001",
		Value: big.NewInt(0),
		Data:  "",
	}
	flag, err := c.UseSendTransaction(tx)
	fmt.Println("The tx hash:", err, flag)

	flag, err = c.UseSendTransaction(tx)
	fmt.Println("The tx hash:", err, flag)
}