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

func GenerateRandomTokenAccount() {
	c := client.NewClient(rpc.DevnetRPCEndpoint)

	aliceRandomTokenAccount := types.NewAccount()
	fmt.Println("alice random token account:", aliceRandomTokenAccount.PublicKey.ToBase58())

	//----------------------------fund account --------------------------
	sig, err := c.RequestAirdrop(
		context.TODO(),
		aliceRandomTokenAccount.PublicKey.ToBase58(), // address
		1e9, // lamports (1 SOL = 10^9 lamports)
	)
	if err != nil {
		log.Fatalf("failed to request airdrop, err: %v", err)
	}
	fmt.Println("-----------------------------------------air drop succeeded-------------------------------")
	fmt.Println(sig)

	rentExemptionBalance, err := c.GetMinimumBalanceForRentExemption(context.Background(), tokenprog.TokenAccountSize)
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
					New:      aliceRandomTokenAccount.PublicKey,
					Owner:    common.TokenProgramID,
					Lamports: rentExemptionBalance,
					Space:    tokenprog.TokenAccountSize,
				}),
				tokenprog.InitializeAccount(tokenprog.InitializeAccountParam{
					Account: aliceRandomTokenAccount.PublicKey,
					Mint:    MintPubkey,
					Owner:   AliceSK.PublicKey,
				}),
			},
		}),
		Signers: []types.Account{AliceSK, aliceRandomTokenAccount},
	})
	if err != nil {
		log.Println("generate tx error, err: %v\n", err)
	}

	txhash, err := c.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Println("send tx error, err: %v\n", err)
	}

	log.Println("txhash:", txhash)
}
