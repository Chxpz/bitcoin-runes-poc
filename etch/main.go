package etch

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

var apiKey = os.Getenv("BLOCKCYPHER_API_KEY")

type TX struct {
	Inputs  []Input  `json:"inputs"`
	Outputs []Output `json:"outputs"`
}

type Input struct {
	Addresses []string `json:"addresses"`
}

type Output struct {
	Addresses []string `json:"addresses"`
	Value     int64    `json:"value"`
}

type TokenData struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type TXResponse struct {
	Tx     TX       `json:"tx"`
	Tosign []string `json:"tosign"`
	Hash   string   `json:"hash"`
}

func CreateEtchTransaction(fromAddress, toAddress string, tokenData TokenData, value int64) (*TXResponse, error) {
	url := "https://api.blockcypher.com/v1/btc/test3/txs/new?token=" + apiKey

	data, err := json.Marshal(tokenData)
	if err != nil {
		return nil, fmt.Errorf("error marshalling token data: %v", err)
	}

	tx := TX{
		Inputs: []Input{
			{Addresses: []string{fromAddress}},
		},
		Outputs: []Output{
			{Addresses: []string{toAddress}, Value: value},
			{Addresses: []string{string(data)}, Value: 0},
		},
	}

	txBytes, err := json.Marshal(tx)
	if err != nil {
		return nil, fmt.Errorf("error marshalling transaction: %v", err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(txBytes))
	if err != nil {
		return nil, fmt.Errorf("error sending transaction: %v", err)
	}
	defer resp.Body.Close()

	var result TXResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return nil, fmt.Errorf("error decoding response: %v", err)
	}

	result.Hash = result.Tx.Outputs[0].Addresses[0]

	fmt.Printf("Transaction Created: %s\n", result)
	return &result, nil
}
