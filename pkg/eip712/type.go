package eip712

import "github.com/ethereum/go-ethereum/accounts/abi"

// Types solidity=>go types
// Src: https://github.com/ethereum/go-ethereum/blob/master/accounts/abi/abi_test.go
var (
	Uint256, _    = abi.NewType("uint256", "", nil)
	Uint32, _     = abi.NewType("uint32", "", nil)
	Uint16, _     = abi.NewType("uint16", "", nil)
	Uint8, _      = abi.NewType("uint8", "", nil)
	String, _     = abi.NewType("string", "", nil)
	Bool, _       = abi.NewType("bool", "", nil)
	Bytes, _      = abi.NewType("bytes", "", nil)
	Bytes32, _    = abi.NewType("bytes32", "", nil)
	Address, _    = abi.NewType("address", "", nil)
	Uint64Arr, _  = abi.NewType("uint64[]", "", nil)
	AddressArr, _ = abi.NewType("address[]", "", nil)
	Int8, _       = abi.NewType("int8", "", nil)
	Int16T, _     = abi.NewType("int16", "", nil)
	Bytes1T, _    = abi.NewType("bytes1", "", nil)
)
