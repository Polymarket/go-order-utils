package signature

type SignatureType = int

const (
	EOA        SignatureType = iota // EIP712 signatures signed by EOAs
	CONTRACT                        // EIP1271 signatures signed by smart contracts
	POLY_PROXY                      // EIP712 signatures signed by polymarket proxy wallets
)

// Proxy Wallet bytecode to be formatted with the factory and implementation addresses
var PROXY_BYTECODE_HEX = "0x3d3d606380380380913d393d73%s5af4602a57600080fd5b602d8060366000396000f3363d3d373d3d3d363d73%s5af43d82803e903d91602b57fd5bf352e831dd00000000000000000000000000000000000000000000000000000000000000200000000000000000000000000000000000000000000000000000000000000000"
