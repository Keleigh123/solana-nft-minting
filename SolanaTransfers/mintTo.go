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

var Ata = common.PublicKeyFromString("Axpnz51M9b5iYVNcTwLJpLqWrQPQ3LxBEhi9kNTFb3PH")
var RandomTokenAccount = common.PublicKeyFromString("3mjHybVMEAyaqruekTKW9juVXjmqXZCPkBAyBWmaY4N5")

func MintTo() {
	c := client.NewClient(rpc.DevnetRPCEndpoint)

	res, err := c.GetRecentBlockhash(context.Background())
	if err != nil {
		log.Fatalf("get recent block hash error, err: %v\n", err)
	}
	tx, err := types.NewTransaction(types.NewTransactionParam{
		Message: types.NewMessage(types.NewMessageParam{
			FeePayer:        AliceSK.PublicKey,
			RecentBlockhash: res.Blockhash,
			Instructions: []types.Instruction{
				tokenprog.MintToChecked(tokenprog.MintToCheckedParam{
					Mint:     MintPubkey,
					Auth:     AliceSK.PublicKey,
					Signers:  []common.PublicKey{},
					To:       RandomTokenAccount,
					Amount:   1e8,
					Decimals: 8,
				}),
			},
		}),
		Signers: []types.Account{AliceSK, AliceSK},
	})
	if err != nil {
		log.Fatalf("generate tx error, err: %v\n", err)
	}

	txhash, err := c.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Fatalf("send raw tx error, err: %v\n", err)
	}

	log.Println("txhash:", txhash)
}
