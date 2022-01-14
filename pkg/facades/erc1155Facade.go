package facades

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/polymarket/order-utils/pkg/contracts/limitOrder"
)

const (
	ERC1155_TRANSFER_FROM = "func_733NCGU"
	ERC1155_BALANCE_OF    = "balanceOf"
)

type ERC1155Facade interface {
	TransferFrom(fromAddress, toAddress, tokenAddress common.Address, value, tokenId *big.Int) ([]byte, error)
	BalanceOf(address common.Address, id *big.Int) ([]byte, error)
}

type ERC1155FacadeImpl struct {
	methods map[string]abi.Method
}

func NewERC1155FacadeImpl() (*ERC1155FacadeImpl, error) {
	limitOrderProtocolAbi, err := abi.JSON(strings.NewReader(limitOrder.LimitOrderABI))
	if err != nil {
		return nil, err
	}
	return &ERC1155FacadeImpl{
		methods: limitOrderProtocolAbi.Methods,
	}, nil
}

func (e *ERC1155FacadeImpl) TransferFrom(fromAddress, toAddress, tokenAddress common.Address, value, tokenId *big.Int) ([]byte, error) {
	method := e.methods[ERC1155_TRANSFER_FROM]
	inputs, err := method.Inputs.Pack(fromAddress, toAddress, value, tokenAddress, tokenId, 0x0)
	if err != nil {
		return nil, err
	}
	return append(method.ID[:], inputs[:]...), nil
}

func (e *ERC1155FacadeImpl) BalanceOf(address common.Address, id *big.Int) ([]byte, error) {
	method := e.methods[ERC1155_BALANCE_OF]
	inputs, err := method.Inputs.Pack(address, id)
	if err != nil {
		return nil, err
	}
	return append(method.ID[:], inputs[:]...), nil
}
