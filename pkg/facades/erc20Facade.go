package facades

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/polymarket/order-utils/pkg/contracts/erc20"
)

const (
	ERC20_TRANSFER_FROM = "transferFrom"
	ERC20_BALANCE_OF    = "balanceOf"
)

type ERC20Facade interface {
	TransferFrom(fromAddress, toAddress common.Address, value *big.Int) ([]byte, error)
	BalanceOf(address common.Address) ([]byte, error)
}

type ERC20FacadeImpl struct {
	methods map[string]abi.Method
}

func NewERC20FacadeImpl() (*ERC20FacadeImpl, error) {
	erc20Abi, err := abi.JSON(strings.NewReader(erc20.Erc20ABI))
	if err != nil {
		return nil, err
	}

	return &ERC20FacadeImpl{
		methods: erc20Abi.Methods,
	}, nil
}

func (e *ERC20FacadeImpl) TransferFrom(fromAddress, toAddress common.Address, value *big.Int) ([]byte, error) {
	method := e.methods[ERC20_TRANSFER_FROM]
	inputs, err := method.Inputs.Pack(fromAddress, toAddress, value)
	if err != nil {
		return nil, err
	}
	return append(method.ID[:], inputs[:]...), nil
}

func (e *ERC20FacadeImpl) BalanceOf(address common.Address) ([]byte, error) {
	method := e.methods[ERC20_BALANCE_OF]
	inputs, err := method.Inputs.Pack(address)
	if err != nil {
		return nil, err
	}
	return append(method.ID[:], inputs[:]...), nil
}
