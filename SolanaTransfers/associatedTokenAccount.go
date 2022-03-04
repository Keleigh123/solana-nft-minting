package SolanaTransfers

import (
	"context"
	//	"fmt"

	"log"

	"github.com/portto/solana-go-sdk/client"
	"github.com/portto/solana-go-sdk/common"

	"github.com/portto/solana-go-sdk/program/assotokenprog"
	//"github.com/portto/solana-go-sdk/program/sysprog"
	"github.com/portto/solana-go-sdk/program/tokenprog"

	// 	//"github.com/portto/solana-go-sdk/program/tokenprog"
	"github.com/portto/solana-go-sdk/rpc"
	"github.com/portto/solana-go-sdk/types"
)

var DestSign, _ = types.AccountFromBytes([]byte{173, 163, 231, 61, 78, 211, 41, 172, 87, 28, 11, 116, 211, 90, 18, 72, 247, 185, 100, 173, 197, 247, 46, 235, 254, 215, 211, 37, 23, 189, 224, 27, 152, 106, 81, 182, 39, 130, 143, 103, 89, 226, 143, 169, 180, 189, 166, 60, 47, 82, 18, 127, 244, 140, 72, 104, 161, 94, 82, 15, 149, 254, 135, 41})
var DestinationWallet = "BFxzxbnu4BghTW67cDa3wYB2zLZA5quGWQmiCeCmWx7i"

func GenerateTokenAccount(MintKey string) (string, error) {
	// 	log.Println("------------------------------------------check point 11-------------------------------", AliceDestSK.PublicKey)
	p := client.NewClient(rpc.TestnetRPCEndpoint)

	// destinationWallet := types.NewAccount()
	// fmt.Println("Wallet Address:", destinationWallet.PublicKey.ToBase58())
	// fmt.Println("Private Key:", destinationWallet.PrivateKey)

	// //fund new account----------------------------------------------------------------------------------------
	// sig, err := p.RequestAirdrop(
	// 	context.TODO(),
	// 	destinationWallet.PublicKey.ToBase58(), // address
	// 	1e9,                                    // lamports (1 SOL = 10^9 lamports)
	// )
	// if err != nil {
	// 	log.Fatalf("failed to request airdrop, err: %v", err)
	// }
	// fmt.Println("-----------------------------------------air drop succeeded for wallet 2-------------------------------")
	// fmt.Println(sig)
	// var AliceDestSK = destinationWallet.PublicKey.ToBase58()
	destination, _, err := common.FindAssociatedTokenAddress(common.PublicKeyFromString(DestinationWallet), common.PublicKeyFromString(MintKey))
	if err != nil {
		log.Fatalf("failed to find a valid ata, err: %v", err)
	}

	var destTokenAddress = destination.ToBase58()
	// 	var dest = destination.String()
	// rentExemptionBalance, err := c.GetMinimumBalanceForRentExemption(context.Background(), tokenprog.TokenAccountSize)
	// if err != nil {
	// 	log.Fatalf("get min balacne for rent exemption, err: %v", err)
	// }

	res, err := p.GetRecentBlockhash(context.Background())
	if err != nil {
		log.Fatalf("get recent block hash error, err: %v\n", err)
	}
	log.Println("------------------------------------------check point 12-------------------------------")
	tx, err := types.NewTransaction(types.NewTransactionParam{
		Message: types.NewMessage(types.NewMessageParam{
			FeePayer:        common.PublicKeyFromString(DestinationWallet),
			RecentBlockhash: res.Blockhash,
			Instructions: []types.Instruction{
				assotokenprog.CreateAssociatedTokenAccount(assotokenprog.CreateAssociatedTokenAccountParam{
					Funder:                 common.PublicKeyFromString(DestinationWallet),
					Owner:                  common.PublicKeyFromString(DestinationWallet),
					Mint:                   common.PublicKeyFromString(MintKey),
					AssociatedTokenAccount: destination,
				}),
				// 				// sysprog.CreateAccount(sysprog.CreateAccountParam{
				// 				// 	From:     AliceDestSK.PublicKey,
				// 				// 	New:      destinationWallet.PublicKey,
				// 				// 	Owner:    common.TokenProgramID,
				// 				// 	Lamports: rentExemptionBalance,
				// 				// 	Space:    tokenprog.TokenAccountSize,
				// 				// }),
				tokenprog.InitializeAccount(tokenprog.InitializeAccountParam{
					Account: destination,
					Mint:    common.PublicKeyFromString(MintKey),
					Owner:   common.PublicKeyFromString(DestinationWallet),
				}),
				// tokenprog.Transfer(tokenprog.TransferParam{
				// 	From:    common.PublicKeyFromString(DestinationWallet),
				// 	To:      destination,
				// 	Auth:    common.PublicKeyFromString(DestinationWallet),
				// 	Signers: []common.PublicKey{},
				// 	Amount:  1e9,
				// }),
			},
		}),
		Signers: []types.Account{DestSign},
	})
	if err != nil {
		log.Fatalf("generate tx error, err: %v\n", err)
	}
	log.Println("------------------------------------------check point 13-------------------------------")
	txhash, err := p.SendTransaction(context.Background(), tx)
	log.Println("------------------------------------------check point 14-------------------------------", destination)
	if err != nil {
		log.Println("send tx error, ", err)
	}
	log.Println("------------------------------------------check point 15-------------------------------")
	log.Println("txhash for token account:", txhash)
	return destTokenAddress, nil
}
