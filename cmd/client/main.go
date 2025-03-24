package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

const target = "0000" // The target prefix for the hash (PoW difficulty)

type QuoteRequest struct {
	Challenge string `json:"challenge"`
	Nonce     string `json:"nonce"`
}

type QuoteResponse struct {
	Quote string `json:"quote"`
}

type ChallengeResponse struct {
	Challenge string `json:"challenge"`
}

func generateProofOfWork(challenge string) (string, string) {
	nonce := 0
	var hash string
	for {
		data := challenge + strconv.Itoa(nonce)
		hash = fmt.Sprintf("%x", sha256.Sum256([]byte(data)))
		if hash[:len(target)] == target {
			break
		}
		nonce++
	}
	return hash, strconv.Itoa(nonce)
}

func main() {

	// Step 1: Get the challenge from the server
	resp, err := http.Get("http://localhost:8098/api/v1/challenge")
	if err != nil {
		fmt.Println("Error fetching challenge:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// Read the response containing the challenge
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading challenge response:", err)
		os.Exit(1)
	}

	// Parse the challenge response
	var challengeResp ChallengeResponse
	if err := json.Unmarshal(body, &challengeResp); err != nil {
		fmt.Println("Error unmarshaling challenge response:", err)
		os.Exit(1)
	}

	// Step 2: Solve the Proof of Work challenge
	challenge := challengeResp.Challenge
	_, nonce := generateProofOfWork(challenge)

	// Prepare the request to send to the server with the challenge and nonce
	quoteRequest := QuoteRequest{
		Challenge: challenge,
		Nonce:     nonce,
	}

	// Marshal the request into JSON
	requestData, err := json.Marshal(quoteRequest)
	if err != nil {
		fmt.Println("Error marshaling request:", err)
		os.Exit(1)
	}

	// Step 3: Send the PoW solution to the server to get the quote
	resp, err = http.Post("http://localhost:8098/api/v1/quote", "application/json", bytes.NewBuffer(requestData))
	if err != nil {
		fmt.Println("Error sending request:", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// Read the response from the server
	body, err = io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		os.Exit(1)
	}

	// If the response is successful, parse and display the quote
	if resp.StatusCode == http.StatusOK {
		var quoteResponse QuoteResponse
		if err := json.Unmarshal(body, &quoteResponse); err != nil {
			fmt.Println("Error unmarshaling response:", err)
			os.Exit(1)
		}

		// Print the quote received from the server
		fmt.Println("Received quote:", quoteResponse.Quote)
	} else {
		fmt.Println("Failed to get quote. Status:", resp.Status)
		fmt.Println("Response body:", string(body))
	}
}
