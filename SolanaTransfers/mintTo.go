package SolanaTransfers

import (
	"context"
	"log"

	"github.com/portto/solana-go-sdk/client"
	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/program/tokenprog"
	"github.com/portto/solana-go-sdk/rpc"
	"github.com/portto/solana-go-sdk/types"
)

//var Ata = common.PublicKeyFromString("Axpnz51M9b5iYVNcTwLJpLqWrQPQ3LxBEhi9kNTFb3PH")
//var RandomTokenAccount = common.PublicKeyFromString("3mjHybVMEAyaqruekTKW9juVXjmqXZCPkBAyBWmaY4N5")

func MintTo(RandomTokenAccount string, MintKey string) error {
	c := client.NewClient(rpc.TestnetRPCEndpoint)
	log.Println("-----------------------------------checkpoint 6 ------------------------------------")
	res, err := c.GetRecentBlockhash(context.Background())
	if err != nil {
		log.Fatalf("get recent block hash error, err: %v\n", err)
	}
	log.Println("-----------------------------------checkpoint 7 ------------------------------------")
	tx, err := types.NewTransaction(types.NewTransactionParam{
		Message: types.NewMessage(types.NewMessageParam{
			FeePayer:        AliceSK.PublicKey,
			RecentBlockhash: res.Blockhash,
			Instructions: []types.Instruction{
				tokenprog.MintToChecked(tokenprog.MintToCheckedParam{
					Mint:     common.PublicKeyFromString(MintKey),
					Auth:     AliceSK.PublicKey,
					Signers:  []common.PublicKey{},
					To:       common.PublicKeyFromString(RandomTokenAccount),
					Amount:   1e8,
					Decimals: 8,
				}),
			},
		}),
		Signers: []types.Account{AliceSK},
	})
	log.Println("-----------------------------------checkpoint 8 ------------------------------------")
	if err != nil {
		log.Println("generate tx error, err: %v\n", err)
	}
	log.Println("-----------------------------------checkpoint 9 ------------------------------------")
	txhash, err := c.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Println("send raw tx error, err: %v\n", err)
	}
	log.Println("-----------------------------------checkpoint 10 ------------------------------------")
	log.Println("txhash:", txhash)
	return nil
}
