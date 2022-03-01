package main

import (
	"context"
	"fmt"
	"log"

	"github.com/portto/solana-go-sdk/client"
	//"github.com/portto/solana-go-sdk/client/rpc"
	"github.com/portto/solana-go-sdk/rpc"
	"github.com/portto/solana-go-sdk/types"
)

func main() {
	// create a RPC client
	c := client.NewClient(rpc.TestnetRPCEndpoint)

	// get the current running Solana version
	response, err := c.GetVersion(context.TODO())
	if err != nil {
		panic(err)
	}

	fmt.Println("version", response.SolanaCore)
	// // create a new wallet using types.NewAccount()
	wallet := types.NewAccount()

	// // // display the wallet public and private keys
	fmt.Println("Wallet Address:", wallet.PublicKey.ToBase58())
	fmt.Println("Private Key:", wallet.PrivateKey)

	balance, err := c.GetBalance(
		context.TODO(),
		wallet.PublicKey.ToBase58(),
	)
	if err != nil {
		log.Fatalf("failed to get balance, err: %v", err)
	}
	fmt.Printf("balance: %v\n", balance)

}
