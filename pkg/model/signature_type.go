package model

type SignatureType = int

const (
	EOA        SignatureType = iota // EIP712 signatures signed by EOAs
	CONTRACT                        // EIP1271 signatures signed by smart contracts
	POLY_PROXY                      // EIP712 signatures signed by polymarket proxy wallets
)
