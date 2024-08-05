package main

import (
	"fmt"

	"github.com/Chxpz/bitcoin-runes-poc/db"
	"github.com/Chxpz/bitcoin-runes-poc/etch"
	"github.com/Chxpz/bitcoin-runes-poc/mint"
)

func main() {
	db.InitDB()
	defer db.Dbpool.Close()

	fromAddress := "tb1qe5mfmh5p8z355zwmnq6j59r8qjseecmcf8dkvd"
	toAddress := "tb1qe5mfmh5p8z355zwmnq6j59r8qjseecmcf8dkvd"
	tokenData := etch.TokenData{
		ID:          "unique-token-id-1234",
		Name:        "MyToken",
		Description: "This is a test token",
	}
	value := int64(1000)

	txResponse, err := etch.CreateEtchTransaction(fromAddress, toAddress, tokenData, value)
	if err != nil {
		fmt.Printf("Error creating etch transaction: %v\n", err)
		db.SaveTransaction(txResponse.Hash, txResponse.Hash, "pending")
		return
	}

	fmt.Printf("Etch Transaction Created: %s\n", txResponse.Hash)

	err = mint.SendTransaction(txResponse)
	if err != nil {
		fmt.Printf("Error sending mint transaction: %v\n", err)
		db.SaveTransaction(txResponse.Hash, txResponse.Hash, "pending")
		return
	}

	fmt.Printf("Mint Transaction Sent for Tx: %s\n", txResponse.Hash)
	db.SaveTransaction(txResponse.Hash, txResponse.Hash, txResponse.Hash)
}
