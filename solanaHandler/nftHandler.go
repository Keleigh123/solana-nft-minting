package solanaHandler

import (
	"fmt"
	"log"
	"solana-wallet/SolanaTransfers"
)

func HandleCalls() {
	firstAction, err := SolanaTransfers.Mint()
	if err != nil {
		log.Println("Error when creating NFT")
		fmt.Println(err)
	} else {
		fmt.Println(firstAction)
		secondAction, mintKey, err := SolanaTransfers.GenerateRandomTokenAccount(firstAction)
		if err != nil {
			log.Println("Error when creating an Account to hold NFT")
			fmt.Println(err)
		} else {
			fmt.Println(secondAction, mintKey)
			err := SolanaTransfers.MintTo(secondAction, mintKey)
			if err != nil {
				log.Println("Error when creating an Account to hold NFT")
				fmt.Println(err)
			} else {
				//fmt.Println(thirdAction, mintKey)
				thirdAction, err := SolanaTransfers.GenerateTokenAccount(mintKey)
				if err != nil {
					log.Println("Error when creating an Account to hold NFT")
					fmt.Println(err)
				} else {
					fmt.Println(thirdAction)
					err := SolanaTransfers.TransferTokens(mintKey, secondAction, thirdAction)
					if err != nil {
						log.Println("Error when creating an Account to hold NFT")
						fmt.Println(err)
					}
				}
			}
		}

	}

	//SolanaTransfers.GenerateTokenAccount()
	//SolanaTransfers.TransferTokens()

}
