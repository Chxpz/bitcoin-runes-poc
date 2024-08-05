package mint

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/Chxpz/bitcoin-runes-poc/etch"
)

var apiKey = os.Getenv("BLOCKCYPHER_API_KEY")

type SendTXRequest struct {
	Tx         *etch.TX `json:"tx"`
	Tosign     []string `json:"tosign"`
	Signatures []string `json:"signatures"`
}

func SignTransaction(tosign []string) ([]string, error) {
	// Simulate signing process
	// This function should use the private key to create the signature for each tosign item.
	// Here we are simulating it with dummy signatures for demonstration purposes.
	signatures := make([]string, len(tosign))
	for i := range tosign {
		signatures[i] = "dummy-signature-for-" + tosign[i] // Replace with actual signing logic
	}
	return signatures, nil
}

func SendTransaction(txResponse *etch.TXResponse) error {
	url := "https://api.blockcypher.com/v1/btc/test3/txs/send?token=" + apiKey

	signatures, err := SignTransaction(txResponse.Tosign)
	if err != nil {
		return fmt.Errorf("error signing transaction: %v", err)
	}

	sendTxReq := SendTXRequest{
		Tx:         &txResponse.Tx,
		Tosign:     txResponse.Tosign,
		Signatures: signatures,
	}

	txBytes, err := json.Marshal(sendTxReq)
	if err != nil {
		return fmt.Errorf("error marshalling transaction data: %v", err)
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(txBytes))
	if err != nil {
		return fmt.Errorf("error sending transaction: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("received non-200 response: %d - %s", resp.StatusCode, string(body))
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %v", err)
	}

	fmt.Println("Response Body:", string(body))

	var result map[string]interface{}
	err = json.NewDecoder(bytes.NewBuffer(body)).Decode(&result)
	if err != nil {
		return fmt.Errorf("error decoding response: %v", err)
	}

	if errMsg, ok := result["error"]; ok {
		return fmt.Errorf("API error: %v", errMsg)
	}

	fmt.Printf("Mint Transaction Sent for Tx: %s\n", txResponse.Hash)
	return nil
}
