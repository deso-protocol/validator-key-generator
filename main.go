package main

import (
	"flag"
	"fmt"
	"github.com/deso-protocol/core/bls"
	"github.com/deso-protocol/core/lib"
)

func getBLSVotingAuthorizationAndPublicKey(blsKeyStore *lib.BLSKeystore, transactorPublicKey *lib.PublicKey) (
	*bls.PublicKey, *bls.Signature,
) {
	votingAuthPayload := lib.CreateValidatorVotingAuthorizationPayload(transactorPublicKey.ToBytes())
	votingAuthorization, err := blsKeyStore.GetSigner().Sign(votingAuthPayload)
	if err != nil {
		panic(err)
	}
	return blsKeyStore.GetSigner().GetPublicKey(), votingAuthorization
}

func main() {

	flagBLSSeedPhrase := flag.String("bls-seed-phrase",
		"", "A 12-word BIP39 seed phrase to use to generate a BLS key pair for your validator.")
	flagDeSoPublicKey := flag.String("deso-public-key",
		"", "The DeSo public key you will be using for your validator.")
	flag.Parse()

	if flagBLSSeedPhrase == nil || *flagBLSSeedPhrase == "" {
		panic("Please provide a BLS seed phrase.")
	}

	if flagDeSoPublicKey == nil || *flagDeSoPublicKey == "" {
		panic("Please provide a DeSo public key.")
	}

	// Parse BIP39 Validator seed phrase
	// E.g. "category ignore around vibrant delay cargo apart truly rabbit blue master cash"
	keystore, err := lib.NewBLSKeystore(*flagBLSSeedPhrase)
	if err != nil {
		panic(err)
	}
	// Parse your DeSo Public Key
	// E.g. BC1YLhS6ruuvtGX58AG8gAvEjhsBR2xdZB54AaGUZ43MnS3nWcm5RYx
	pubKeyBytes, _, err := lib.Base58CheckDecode(*flagDeSoPublicKey)
	if err != nil {
		panic(err)
	}
	publicKey, votingAuthorization := getBLSVotingAuthorizationAndPublicKey(keystore, lib.NewPublicKey(pubKeyBytes))
	fmt.Println("Validator BLS PublicKey: ", publicKey.ToString())
	fmt.Println("Validator Voting Authorization: ", votingAuthorization.ToString())
}
