package SolanaTransfers

import (
	"context"
	"log"

	"github.com/portto/solana-go-sdk/client"
	"github.com/portto/solana-go-sdk/common"

	//"github.com/portto/solana-go-sdk/program/assotokenprog"
	"github.com/portto/solana-go-sdk/program/tokenprog"
	"github.com/portto/solana-go-sdk/rpc"
	"github.com/portto/solana-go-sdk/types"
)

//var AliceTokenRandomTokenPubkey = common.PublicKeyFromString("BV5EnZ3wFAhRt4Kff6VXRMjebEn8Jg9zQaGLby52Aro")
//var AliceTokenATAPubkey = common.PublicKeyFromString("2oNTPr1zd4AwDKxrAfsUp7Qt6cCtd88v5tdAVBa4qeqy")
//var Ata = common.PublicKeyFromString("Axpnz51M9b5iYVNcTwLJpLqWrQPQ3LxBEhi9kNTFb3PH")
//var AliceDestSK, _ = types.AccountFromBytes([]byte{209, 1, 1, 254, 161, 211, 239, 134, 128, 149, 40, 25, 210, 57, 242, 208, 231, 173, 239, 30, 242, 209, 87, 194, 31, 71, 104, 245, 207, 251, 160, 174, 189, 69, 62, 103, 242, 33, 220, 216, 20, 212, 111, 4, 187, 205, 33, 83, 5, 176, 60, 199, 249, 65, 154, 110, 203, 81, 225, 125, 167, 202, 129, 99})
func TransferTokens(mintKey string, source string, destination string) error {
	c := client.NewClient(rpc.TestnetRPCEndpoint)
	log.Println("-----------------------------------checkpoint 1 ------------------------------------")
	// destinationWallet := types.NewAccount()
	// fmt.Println("Wallet Address:", destinationWallet.PublicKey.ToBase58())
	// fmt.Println("Private Key:", destinationWallet.PrivateKey)

	// //fund new account----------------------------------------------------------------------------------------
	// sig, err := c.RequestAirdrop(
	// 	context.TODO(),
	// 	destinationWallet.PublicKey.ToBase58(), // address
	// 	1e9,                                    // lamports (1 SOL = 10^9 lamports)
	// )
	// if err != nil {
	// 	log.Fatalf("failed to request airdrop, err: %v", err)
	// }
	// fmt.Println("-----------------------------------------air drop succeeded for wallet 2-------------------------------")
	// fmt.Println(sig)
	// // // // // display the wallet public and private keys
	// fmt.Println("Wallet Address:", destination.PublicKey.ToBase58())
	// fmt.Println("Private Key:", destination.PrivateKey)
	log.Println("-----------------------------------checkpoint 2 ------------------------------------")
	// dest := destination.PublicKey.ToBase58()
	res, err := c.GetRecentBlockhash(context.Background())
	if err != nil {
		log.Fatalf("get recent block hash error, err: %v\n", err)
	}
	tx, err := types.NewTransaction(types.NewTransactionParam{
		Message: types.NewMessage(types.NewMessageParam{
			FeePayer:        feePayer.PublicKey,
			RecentBlockhash: res.Blockhash,
			Instructions: []types.Instruction{
				tokenprog.Transfer(tokenprog.TransferParam{
					From:    common.PublicKeyFromString(source),
					To:      common.PublicKeyFromString(destination),
					Auth:    feePayer.PublicKey,
					Signers: []common.PublicKey{},
					Amount:  1,
				}),
				// tokenprog.TransferChecked(tokenprog.TransferCheckedParam{
				// 	From:     common.PublicKeyFromString(source),
				// 	To:       destinationWallet.PublicKey,
				// 	Mint:     common.PublicKeyFromString(mintKey),
				// 	Auth:     feePayer.PublicKey,
				// 	Signers:  []common.PublicKey{},
				// 	Amount:   1e0,
				// 	Decimals: 0,
				// }),
			},
		}),
		Signers: []types.Account{feePayer},
	})
	log.Println("-----------------------------------checkpoint 3 ------------------------------------")
	if err != nil {
		log.Println("failed to new tx", err)
	}
	log.Println("-----------------------------------checkpoint 4 ------------------------------------")
	txhash, err := c.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Println("send raw tx error", err)
	}
	log.Println("-----------------------------------checkpoint 5 ------------------------------------")
	log.Println("txhash for transferring tokens:", txhash)
	return nil
}
