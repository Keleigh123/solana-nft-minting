package SolanaTransfers

import (
	"context"
	"fmt"
	"log"

	"github.com/portto/solana-go-sdk/client"
	"github.com/portto/solana-go-sdk/common"
	"github.com/portto/solana-go-sdk/program/sysprog"
	"github.com/portto/solana-go-sdk/program/tokenprog"
	"github.com/portto/solana-go-sdk/rpc"
	"github.com/portto/solana-go-sdk/types"
)

var AliceSK, _ = types.AccountFromBytes([]byte{162, 128, 223, 203, 33, 217, 35, 50, 114, 79, 106, 50, 93, 174, 66, 2, 47, 22, 191, 158, 233, 41, 109, 52, 49, 255, 214, 3, 3, 182, 50, 185, 2, 212, 203, 14, 73, 174, 65, 37, 136, 138, 5, 84, 53, 62, 136, 198, 69, 3, 211, 0, 20, 214, 9, 140, 211, 24, 14, 197, 109, 104, 35, 65})

func Mint() string {
	log.Println("-----------------------testing--------------")

	c := client.NewClient(rpc.TestnetRPCEndpoint)

	// create an mint account
	mint := types.NewAccount()
	fmt.Println("mint:", mint.PublicKey.ToBase58())

	// get rent
	rentExemptionBalance, err := c.GetMinimumBalanceForRentExemption(
		context.Background(),
		tokenprog.MintAccountSize,
	)
	if err != nil {
		log.Fatalf("get min balacne for rent exemption, err: %v", err)
	}

	res, err := c.GetRecentBlockhash(context.Background())
	if err != nil {
		log.Fatalf("get recent block hash error, err: %v\n", err)
	}
	tx, err := types.NewTransaction(types.NewTransactionParam{
		Message: types.NewMessage(types.NewMessageParam{
			FeePayer:        AliceSK.PublicKey,
			RecentBlockhash: res.Blockhash,
			Instructions: []types.Instruction{
				sysprog.CreateAccount(sysprog.CreateAccountParam{
					From:     AliceSK.PublicKey,
					New:      mint.PublicKey,
					Owner:    common.TokenProgramID,
					Lamports: rentExemptionBalance,
					Space:    tokenprog.MintAccountSize,
				}),
				tokenprog.InitializeMint(tokenprog.InitializeMintParam{
					Decimals:   8,
					Mint:       mint.PublicKey,
					MintAuth:   AliceSK.PublicKey,
					FreezeAuth: nil,
				}),
			},
		}),
		Signers: []types.Account{AliceSK, mint},
	})

	var MintKey = mint.PublicKey
	if err != nil {
		log.Fatalf("generate tx error, err: %v\n", err)
	}

	txhash, err := c.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Fatalf("send tx error, err: %v\n", err)
	}

	log.Println("txhash for minting account:", txhash)
	return MintKey.String()

}
