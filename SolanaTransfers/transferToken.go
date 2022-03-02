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

//var AliceTokenRandomTokenPubkey = common.PublicKeyFromString("BV5EnZ3wFAhRt4Kff6VXRMjebEn8Jg9zQaGLby52Aro")
//var AliceTokenATAPubkey = common.PublicKeyFromString("2oNTPr1zd4AwDKxrAfsUp7Qt6cCtd88v5tdAVBa4qeqy")
var Ata = common.PublicKeyFromString("Axpnz51M9b5iYVNcTwLJpLqWrQPQ3LxBEhi9kNTFb3PH")

func TransferTokens(mintKey string, randomAccount string, destination string) error {
	c := client.NewClient(rpc.TestnetRPCEndpoint)

	res, err := c.GetRecentBlockhash(context.Background())
	if err != nil {
		log.Fatalf("get recent block hash error, err: %v\n", err)
	}
	tx, err := types.NewTransaction(types.NewTransactionParam{
		Message: types.NewMessage(types.NewMessageParam{
			FeePayer:        AliceSK.PublicKey,
			RecentBlockhash: res.Blockhash,
			Instructions: []types.Instruction{
				tokenprog.TransferChecked(tokenprog.TransferCheckedParam{
					From:     common.PublicKeyFromString(randomAccount),
					To:       common.PublicKeyFromString(destination),
					Mint:     common.PublicKeyFromString(mintKey),
					Auth:     AliceSK.PublicKey,
					Signers:  []common.PublicKey{},
					Amount:   1e8,
					Decimals: 8,
				}),
			},
		}),
		Signers: []types.Account{AliceSK, AliceSK},
	})
	if err != nil {
		log.Fatalf("failed to new tx, err: %v", err)
	}

	txhash, err := c.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Fatalf("send raw tx error, err: %v\n", err)
	}

	log.Println("txhash for transferring tokens:", txhash)
	return nil
}
