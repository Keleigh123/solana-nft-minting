This is for nft development using Solana.

1. The main.go has a file with just some native asset testing and a call to the nft handler.
2. Firstly, the createMint.go file will create the token.
3. The randomTokenAccount.go file is the account generation for the account that will hold the nft.
4. mintTo.go is the file that will mint the nft ad send to the randomTokenAccount
5. associatedTokenAccount.go is another account creation .
6. transferToken.go is to send tokens from one destination to another, destination is from random token account to associated token account.