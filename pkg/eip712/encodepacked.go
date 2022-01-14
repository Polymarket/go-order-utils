package eip712

import (
	"encoding/binary"
	"errors"
	"fmt"
	"math/big"
	"reflect"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

//Implementation of solidity abi.encodePacked in golang, not in core geth
//Taken from: https://github.com/ethereum/go-ethereum/pull/22273, with minor modifications
//As our usecase will not have any ambiguous types, we can safely use this implementation

// EncodePacked implements the non-standard encoding available in solidity.
// Since encoding is ambigious there is no decoding function.

func EncodePacked(args []abi.Type, values []interface{}) ([]byte, error) {
	enc := make([]byte, 0)
	for idx, arg := range args {
		switch arg.T {
		case abi.TupleTy:
			return []byte{}, errors.New("Not implemented")
		case abi.ArrayTy, abi.SliceTy:
		default:
			packed, err := encodePackElement(arg, reflect.ValueOf(values[idx]))
			if err != nil {
				return []byte{}, err
			}
			enc = append(enc, packed...)
		}
	}
	return enc, nil
}

// indirect recursively dereferences the value until it either gets the value
// or finds a big.Int
//Src: https://github.com/ethereum/go-ethereum/blob/master/accounts/abi/reflect.go#L52
func indirect(v reflect.Value) reflect.Value {
	if v.Kind() == reflect.Ptr && v.Elem().Type() != reflect.TypeOf(big.Int{}) {
		return indirect(v.Elem())
	}
	return v
}

// mustArrayToByteSlice creates a new byte slice with the exact same size as value
// and copies the bytes in value to the new slice.
//Src: https://github.com/ethereum/go-ethereum/blob/master/accounts/abi/reflect.go#L89
func mustArrayToByteSlice(value reflect.Value) reflect.Value {
	slice := reflect.MakeSlice(reflect.TypeOf([]byte{}), value.Len(), value.Len())
	reflect.Copy(slice, value)
	return slice
}

func encodePackElement(t abi.Type, value reflect.Value) ([]byte, error) {
	value = indirect(value)

	switch t.T {
	case abi.IntTy, abi.UintTy:
		return encodePackedNum(t, value), nil
	case abi.StringTy:
		return encodePackedByteSlice(t, []byte(value.String())), nil
	case abi.AddressTy, abi.FixedBytesTy:
		if value.Kind() == reflect.Array {
			value = mustArrayToByteSlice(value)
		}

		return encodePackedByteSlice(t, value.Bytes()), nil
	case abi.BoolTy:
		if value.Bool() {
			return []byte{1}, nil
		}
		return []byte{0}, nil
	case abi.BytesTy:
		if value.Kind() == reflect.Array {
			value = mustArrayToByteSlice(value)
		}
		if value.Type() != reflect.TypeOf([]byte{}) {
			return []byte{}, errors.New("Bytes type is neither slice nor array")
		}
		return encodePackedByteSlice(t, value.Bytes()), nil
	default:
		return []byte{}, fmt.Errorf("Could not encode pack element, unknown type: %v", t.T)
	}
}

func encodePackedNum(t abi.Type, value reflect.Value) []byte {
	bytes := make([]byte, 8)
	switch kind := value.Kind(); kind {
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		binary.BigEndian.PutUint64(bytes, value.Uint())
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		binary.BigEndian.PutUint64(bytes, uint64(value.Int()))
	case reflect.Ptr:
		big := new(big.Int).Set(value.Interface().(*big.Int))
		bytes = big.Bytes()
	default:
		panic("abi: fatal error")
	}
	return encodePackedByteSlice(t, bytes)
}

func encodePackedByteSlice(t abi.Type, value []byte) []byte {
	size := t.Size / 8
	// If size is not set in the type, use the length of the value to pad
	if size == 0 {
		size = len(value)
	}
	padded := common.LeftPadBytes(value, size)
	return padded[len(padded)-size:]
}
