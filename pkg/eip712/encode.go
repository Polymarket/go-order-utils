package eip712

import (
	"github.com/ethereum/go-ethereum/accounts/abi"
)

func Encode(args []abi.Type, values []interface{}) ([]byte, error) {
	arguments := make([]abi.Argument, 0)
	for _, t := range args {
		argument := abi.Argument{Type: t}
		arguments = append(arguments, argument)
	}

	return abi.Arguments(arguments).Pack(values...)
}
