package SolanaTransfers

import (
	"context"
	"fmt"
	"log"

	"github.com/portto/solana-go-sdk/client"
	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/program/assotokenprog"

	//"github.com/portto/solana-go-sdk/program/sysprog"
	//"github.com/portto/solana-go-sdk/program/tokenprog"
	"github.com/portto/solana-go-sdk/rpc"
	"github.com/portto/solana-go-sdk/types"
)

//var AliceSK, _ = types.AccountFromBytes([]byte{162, 128, 223, 203, 33, 217, 35, 50, 114, 79, 106, 50, 93, 174, 66, 2, 47, 22, 191, 158, 233, 41, 109, 52, 49, 255, 214, 3, 3, 182, 50, 185, 2, 212, 203, 14, 73, 174, 65, 37, 136, 138, 5, 84, 53, 62, 136, 198, 69, 3, 211, 0, 20, 214, 9, 140, 211, 24, 14, 197, 109, 104, 35, 65})
//var MintPubkey = common.PublicKeyFromString("5BKPuBAAPyLf8WwM43tuAdkt4oEbZn5LJrALt4pAN9V2")

func GenerateTokenAccount(MintKey string) (string, error) {
	log.Println("------------------------------------------check point 11-------------------------------")
	c := client.NewClient(rpc.TestnetRPCEndpoint)

	// AliceRandomTokenAccount := types.NewAccount()
	// fmt.Println("alice token account:", AliceRandomTokenAccount.PublicKey.ToBase58())
	var mint = common.PublicKeyFromString(MintKey)
	ata, _, err := common.FindAssociatedTokenAddress(AliceSK.PublicKey, mint)
	if err != nil {
		log.Fatalf("find ata error, err: %v", err)
	}
	fmt.Println("ata:", ata.ToBase58())
	var ataKey = ata.ToBase58()
	// rentExemptionBalance, err := c.GetMinimumBalanceForRentExemption(context.Background(), tokenprog.TokenAccountSize)
	// if err != nil {
	// 	log.Fatalf("get min balacne for rent exemption, err: %v", err)
	// }

	res, err := c.GetRecentBlockhash(context.Background())
	if err != nil {
		log.Fatalf("get recent block hash error, err: %v\n", err)
	}
	log.Println("------------------------------------------check point 12-------------------------------")
	tx, err := types.NewTransaction(types.NewTransactionParam{
		Message: types.NewMessage(types.NewMessageParam{
			FeePayer:        AliceSK.PublicKey,
			RecentBlockhash: res.Blockhash,
			Instructions: []types.Instruction{
				assotokenprog.CreateAssociatedTokenAccount(assotokenprog.CreateAssociatedTokenAccountParam{
					Funder:                 AliceSK.PublicKey,
					Owner:                  AliceSK.PublicKey,
					Mint:                   common.PublicKeyFromString(MintKey),
					AssociatedTokenAccount: ata,
				}),
			},
		}),
		Signers: []types.Account{AliceSK},
	})
	if err != nil {
		log.Fatalf("generate tx error, err: %v\n", err)
	}
	log.Println("------------------------------------------check point 13-------------------------------")
	txhash, err := c.SendTransaction(context.Background(), tx)
	log.Println("------------------------------------------check point 14-------------------------------")
	if err != nil {
		log.Println("send tx error, err: %v\n", err)
	}
	log.Println("------------------------------------------check point 15-------------------------------")
	log.Println("txhash for token account:", txhash)
	return ataKey, nil
}
