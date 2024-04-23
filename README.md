# About this Repo

This repo contains a simple tool used to generate a validator's BLS public key and
voting authorization given a BIP 39 seed phrase and a DeSo public key. 

The BIP 39 seed phrase is a seed phrase that will be used by a validator node to 
sign blocks for PoS consensus. This does not need to be the same seed phrase used
for your DeSo public key. You will supply this value as the `pos-validator-seed` flag 
when running your node.

The DeSo public key is the public key that will be used when registering your validator
in the validator hub.

You can register your validator at the validator hub.
- [Mainnet](https://explorer.deso.com/validator-settings)
- [Testnet](https://explorer-testnet.deso.com/validator-settings)

# Run
Use the command below and replace `<SEED_PHRASE>` and `<DESO_PUBLIC_KEY>` with your values.
```bash
go run main.go --bls-seed-phrase="<SEED_PHRASE>" --deso-public-key="<DESO_PUBLIC_KEY>"
```

### Example
```bash
go run main.go --bls-seed-phrase="category ignore around vibrant delay cargo apart truly rabbit blue master cash" --deso-public-key="BC1YLhS6ruuvtGX58AG8gAvEjhsBR2xdZB54AaGUZ43MnS3nWcm5RYx"
```