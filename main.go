package main

import (
	// "context"

	// "fmt"
	// "log"

	//"log"

	"fmt"
	"solana-wallet/mintNFT"

	// "github.com/portto/solana-go-sdk/client"
	// "github.com/portto/solana-go-sdk/program/sysprog"
	// "github.com/portto/solana-go-sdk/rpc"
	"github.com/portto/solana-go-sdk/types"
)

//var Alice = ("5TdJYb9AhEhSqLvAG4MPJddHHQSX4breqCSWnx2Ast5k")
var AliceSK, _ = types.AccountFromBytes([]byte{162, 128, 223, 203, 33, 217, 35, 50, 114, 79, 106, 50, 93, 174, 66, 2, 47, 22, 191, 158, 233, 41, 109, 52, 49, 255, 214, 3, 3, 182, 50, 185, 2, 212, 203, 14, 73, 174, 65, 37, 136, 138, 5, 84, 53, 62, 136, 198, 69, 3, 211, 0, 20, 214, 9, 140, 211, 24, 14, 197, 109, 104, 35, 65})

func main() {
	// // create a RPC client
	// c := client.NewClient(rpc.TestnetRPCEndpoint)

	// // get the current running Solana version
	// response, err := c.GetVersion(context.TODO())
	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("version", response.SolanaCore)
	// // // // create a new wallet using types.NewAccount()
	// wallet := types.NewAccount()
	fmt.Println("ALice is the fee payer------------------------", AliceSK.PublicKey)
	// // // // // display the wallet public and private keys
	// fmt.Println("Wallet Address:", wallet.PublicKey.ToBase58())
	// fmt.Println("Private Key:", wallet.PrivateKey)

	// //fund new account----------------------------------------------------------------------------------------
	// sig, err := c.RequestAirdrop(
	// 	context.TODO(),
	// 	wallet.PublicKey.ToBase58(), // address
	// 	1e9,                         // lamports (1 SOL = 10^9 lamports)
	// )
	// if err != nil {
	// 	log.Fatalf("failed to request airdrop, err: %v", err)
	// }
	// fmt.Println("-----------------------------------------air drop succeeded-------------------------------")
	// fmt.Println(sig)

	// recentBlockhashResponse, err := c.GetRecentBlockhash(context.Background())
	// if err != nil {
	// 	log.Fatalf("failed to get recent blockhash, err: %v", err)
	// }

	// tx, err := types.NewTransaction(types.NewTransactionParam{
	// 	Signers: []types.Account{AliceSK},
	// 	Message: types.NewMessage(types.NewMessageParam{
	// 		FeePayer:        AliceSK.PublicKey,
	// 		RecentBlockhash: recentBlockhashResponse.Blockhash,
	// 		Instructions: []types.Instruction{
	// 			sysprog.Transfer(sysprog.TransferParam{
	// 				From:   AliceSK.PublicKey,
	// 				To:     wallet.PublicKey,
	// 				Amount: 1e9,
	// 			}),
	// 		},
	// 	}),
	// })
	// if err != nil {
	// 	log.Println("failed to new a transaction, err: %v", err)
	// }

	// // send tx
	// txhash, err := c.SendTransaction(context.Background(), tx)
	// if err != nil {
	// 	log.Println("failed to send tx, err: %v", err)
	// }

	// log.Println("txhash:", txhash)

	// fmt.Println("------------------wallet pk------------------", wallet.PublicKey)
	// //get balance --------------------------------------------------------------------------------------------------------------
	// balance, err := c.GetBalance(
	// 	context.TODO(),
	// 	wallet.PublicKey.ToBase58(),
	// )
	// if err != nil {
	// 	log.Fatalf("failed to get balance, err: %v", err)
	// }
	// fmt.Println("-------------------------------------------balance is --------------------", balance)
	// fmt.Printf("balance: %v\n", balance)

	// //solanaHandler.HandleCalls()
	// fmt.Println("------------------------------------------------------------------------------------------------------------------------")
	mintNFT.MainCode()
}
