package solanaHandler

import (
	"fmt"
	"log"
	"solana-wallet/SolanaTransfers"
)

func HandleCalls() {
	mintKey, source, err := SolanaTransfers.MainCode()
	if err != nil {
		log.Println("Error when minting NFT")
		fmt.Println(err)
	} else {
		fmt.Println(mintKey, source)
		destination, err := SolanaTransfers.GenerateTokenAccount(mintKey)
		if err != nil {
			log.Println("Error when creating destination account")
			fmt.Println(err)
		} else {
			fmt.Println(destination)
			err := SolanaTransfers.TransferTokens(mintKey, source, destination)
			if err != nil {
				log.Println("Error when transferring NFTs")
				fmt.Println(err)
			}
			fmt.Println("Token Transfer Complete")

			// fmt.Println(destination)

		}
	}
}

// 	firstAction, err := SolanaTransfers.Mint()
// 	if err != nil {
// 		log.Println("Error when creating NFT")
// 		fmt.Println(err)
// 	} else {
// 		fmt.Println(firstAction)
// 		secondAction, mintKey, err := SolanaTransfers.GenerateRandomTokenAccount(firstAction)
// 		if err != nil {
// 			log.Println("Error when creating an Account to hold NFT")
// 			fmt.Println(err)
// 		} else {
// 			fmt.Println(secondAction, mintKey)
// 			err := SolanaTransfers.MintTo(secondAction, mintKey)
// 			if err != nil {
// 				log.Println("Error when creating an Account to hold NFT")
// 				fmt.Println(err)
// 			} else {
// 				//fmt.Println(thirdAction, mintKey)
// 				thirdAction, err := SolanaTransfers.GenerateTokenAccount(mintKey)
// 				if err != nil {
// 					log.Println("Error when creating an Account to hold NFT")
// 					fmt.Println(err)
// 				} else {
// 					fmt.Println(thirdAction)
// 					err := SolanaTransfers.TransferTokens(mintKey, secondAction, thirdAction)
// 					if err != nil {
// 						log.Println("Error when creating an Account to hold NFT")
// 						fmt.Println(err)
// 					}
// 				}
// 			}
// 		}

// 	}

// 	//SolanaTransfers.GenerateTokenAccount()
// 	//SolanaTransfers.TransferTokens()
