package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/rpc"
)

func main() {
	// Replace with your Ethereum node URL
	ethNodeURL := "http://192.168.253.108:8545"

	// Connect to the Ethereum node
	client, err := ethclient.Dial(ethNodeURL)
	if err != nil {
		log.Fatal(err)
	}

	// Replace with your private key
	privateKeyHex := "0x3b8e0b56040a15e75c4fcaba0f753802c73db4a1e29379fd63c5abf735386c4e"
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		log.Fatal(err)
	}

	// Create an Ethereum sender account from the private key
	senderAddress := crypto.PubkeyToAddress(privateKey.PublicKey)
	senderNonce, err := client.PendingNonceAt(context.Background(), senderAddress)
	if err != nil {
		log.Fatal(err)
	}

	// Replace with the recipient's Ethereum address
	recipientAddress := "0xabe3f423136650def254f8467d294114e91818b"

	// Convert ETH to Wei (1 ETH = 10^18 Wei)
	amount := new(big.Int)
	amount.SetString("32000000000000000000", 10) // 32 ETH in Wei

	// Create an Ethereum transaction
	gasLimit := uint64(21000) // Limit of gas provided for the transaction
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	tx := types.NewTransaction(senderNonce, recipientAddress, amount, gasLimit, gasPrice, nil)

	// Sign the transaction with the sender's private key
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(big.NewInt(1)), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	// Send the transaction
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	// Print the transaction hash
	fmt.Printf("Transaction Hash: 0x%x\n", signedTx.Hash())
}
