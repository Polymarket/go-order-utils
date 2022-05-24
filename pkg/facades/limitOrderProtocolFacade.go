package facades

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/polymarket/go-order-utils/pkg/contracts/limitOrder"
	"github.com/polymarket/go-order-utils/pkg/model"
)

const (
	LIMIT_ORDER_PROTOCOL_GET_MAKER_AMOUNT  = "getMakerAmount"
	LIMIT_ORDER_PROTOCOL_GET_TAKER_AMOUNT  = "getTakerAmount"
	LIMIT_ORDER_PROTOCOL_FILL_ORDER        = "fillOrder"
	LIMIT_ORDER_PROTOCOL_CANCEL_ORDER      = "cancelOrder"
	LIMIT_ORDER_PROTOCOL_NONCE             = "nonce"
	LIMIT_ORDER_PROTOCOL_ADVANCE_NONCE     = "advanceNonce"
	LIMIT_ORDER_PROTOCOL_INCREASE_NONCE    = "increaseNonce"
	LIMIT_ORDER_PROTOCOL_AND               = "and"
	LIMIT_ORDER_PROTOCOL_OR                = "or"
	LIMIT_ORDER_PROTOCOL_EQ                = "eq"
	LIMIT_ORDER_PROTOCOL_LT                = "lt"
	LIMIT_ORDER_PROTOCOL_GT                = "gt"
	LIMIT_ORDER_PROTOCOL_TIMESTAMP_BELOW   = "timestampBelow"
	LIMIT_ORDER_PROTOCOL_NONCE_EQUALS      = "nonceEquals"
	LIMIT_ORDER_PROTOCOL_REMAINING         = "remaining"
	LIMIT_ORDER_PROTOCOL_CHECK_PREDICATE   = "checkPredicate"
	LIMIT_ORDER_PROTOCOL_REMAINING_RAW     = "remainingsRaw"
	LIMIT_ORDER_PROTOCOL_DOMAIN_SEPARATOR  = "domainSeparator"
	LIMIT_ORDER_PROTOCOL_BATCH_FILL_ORDERS = "batchFillOrders"
)

type LimitOrderProtocolFacade interface {
	GetMakerAmount(orderMakerAmount, orderTakerAmount, swapTakerAmount *big.Int) ([]byte, error)
	GetTakerAmount(orderMakerAmount, orderTakerAmount, swapMakerAmount *big.Int) ([]byte, error)
	FillOrder(order model.LimitOrder, signature []byte, makingAmount, takingAmount, thresholdAmount *big.Int) ([]byte, error)
	CancelOrder(order model.LimitOrder) ([]byte, error)
	Nonce(address common.Address) ([]byte, error)
	AdvanceNonce(amount uint8) ([]byte, error)
	IncreaseNonce() ([]byte, error)
	And(targets []common.Address, data [][]byte) ([]byte, error)
	Or(targets []common.Address, data [][]byte) ([]byte, error)
	Eq(value *big.Int, target common.Address, data []byte) ([]byte, error)
	Lt(value *big.Int, target common.Address, data []byte) ([]byte, error)
	Gt(value *big.Int, target common.Address, data []byte) ([]byte, error)
	TimestampBelow(time *big.Int) ([]byte, error)
	NonceEquals(makerAddress common.Address, makerNonce *big.Int) ([]byte, error)
	Remaining(orderHash [32]byte) ([]byte, error)
	CheckPredicate(order model.LimitOrder) ([]byte, error)
	RemainingsRaw(orderHashes [][32]byte) ([]byte, error)
	DomainSeparator() ([]byte, error)
	BatchFillOrders(fillData *limitOrder.OrdersLimitOrderFillData) ([]byte, error)
}

type LimitOrderProtocolFacadeImpl struct {
	methods map[string]abi.Method
}

func NewLimitOrderProtocolFacadeImpl() (*LimitOrderProtocolFacadeImpl, error) {
	limitOrderProtocolAbi, err := abi.JSON(strings.NewReader(limitOrder.LimitOrderABI))
	if err != nil {
		return nil, err
	}
	return &LimitOrderProtocolFacadeImpl{
		methods: limitOrderProtocolAbi.Methods,
	}, nil
}

func (l *LimitOrderProtocolFacadeImpl) GetMakerAmount(orderMakerAmount, orderTakerAmount, swapTakerAmount *big.Int) ([]byte, error) {
	method := l.methods[LIMIT_ORDER_PROTOCOL_GET_MAKER_AMOUNT]
	inputs, err := method.Inputs.Pack(orderMakerAmount, orderTakerAmount, swapTakerAmount)
	if err != nil {
		return nil, err
	}
	return append(method.ID[:], inputs[:]...), nil
}

func (l *LimitOrderProtocolFacadeImpl) GetTakerAmount(orderMakerAmount, orderTakerAmount, swapMakerAmount *big.Int) ([]byte, error) {
	method := l.methods[LIMIT_ORDER_PROTOCOL_GET_TAKER_AMOUNT]
	inputs, err := method.Inputs.Pack(orderMakerAmount, orderTakerAmount, swapMakerAmount)
	if err != nil {
		return nil, err
	}
	return append(method.ID[:], inputs[:]...), nil
}

func (l *LimitOrderProtocolFacadeImpl) FillOrder(order model.LimitOrder, signature []byte, makingAmount, takingAmount, thresholdAmount *big.Int) ([]byte, error) {
	method := l.methods[LIMIT_ORDER_PROTOCOL_FILL_ORDER]
	inputs, err := method.Inputs.Pack(order, signature, makingAmount, takingAmount, thresholdAmount)
	if err != nil {
		return nil, err
	}
	return append(method.ID[:], inputs[:]...), nil
}

func (l *LimitOrderProtocolFacadeImpl) CancelOrder(order model.LimitOrder) ([]byte, error) {
	method := l.methods[LIMIT_ORDER_PROTOCOL_CANCEL_ORDER]
	inputs, err := method.Inputs.Pack(order)
	if err != nil {
		return nil, err
	}
	return append(method.ID[:], inputs[:]...), nil
}

func (l *LimitOrderProtocolFacadeImpl) Nonce(address common.Address) ([]byte, error) {
	method := l.methods[LIMIT_ORDER_PROTOCOL_NONCE]
	inputs, err := method.Inputs.Pack(address)
	if err != nil {
		return nil, err
	}
	return append(method.ID[:], inputs[:]...), nil
}

func (l *LimitOrderProtocolFacadeImpl) AdvanceNonce(amount uint8) ([]byte, error) {
	method := l.methods[LIMIT_ORDER_PROTOCOL_ADVANCE_NONCE]
	inputs, err := method.Inputs.Pack(amount)
	if err != nil {
		return nil, err
	}
	return append(method.ID[:], inputs[:]...), nil
}

func (l *LimitOrderProtocolFacadeImpl) IncreaseNonce() ([]byte, error) {
	method := l.methods[LIMIT_ORDER_PROTOCOL_INCREASE_NONCE]
	inputs, err := method.Inputs.Pack()
	if err != nil {
		return nil, err
	}
	return append(method.ID[:], inputs[:]...), nil
}

func (l *LimitOrderProtocolFacadeImpl) And(targets []common.Address, data [][]byte) ([]byte, error) {
	method := l.methods[LIMIT_ORDER_PROTOCOL_AND]
	inputs, err := method.Inputs.Pack(targets, data)
	if err != nil {
		return nil, err
	}
	return append(method.ID[:], inputs[:]...), nil
}

func (l *LimitOrderProtocolFacadeImpl) Or(targets []common.Address, data [][]byte) ([]byte, error) {
	method := l.methods[LIMIT_ORDER_PROTOCOL_OR]
	inputs, err := method.Inputs.Pack(targets, data)
	if err != nil {
		return nil, err
	}
	return append(method.ID[:], inputs[:]...), nil
}

func (l *LimitOrderProtocolFacadeImpl) Eq(value *big.Int, target common.Address, data []byte) ([]byte, error) {
	method := l.methods[LIMIT_ORDER_PROTOCOL_EQ]
	inputs, err := method.Inputs.Pack(value, target, data)
	if err != nil {
		return nil, err
	}
	return append(method.ID[:], inputs[:]...), nil
}

func (l *LimitOrderProtocolFacadeImpl) Lt(value *big.Int, target common.Address, data []byte) ([]byte, error) {
	method := l.methods[LIMIT_ORDER_PROTOCOL_LT]
	inputs, err := method.Inputs.Pack(value, target, data)
	if err != nil {
		return nil, err
	}
	return append(method.ID[:], inputs[:]...), nil
}

func (l *LimitOrderProtocolFacadeImpl) Gt(value *big.Int, target common.Address, data []byte) ([]byte, error) {
	method := l.methods[LIMIT_ORDER_PROTOCOL_GT]
	inputs, err := method.Inputs.Pack(value, target, data)
	if err != nil {
		return nil, err
	}
	return append(method.ID[:], inputs[:]...), nil
}

func (l *LimitOrderProtocolFacadeImpl) TimestampBelow(time *big.Int) ([]byte, error) {
	method := l.methods[LIMIT_ORDER_PROTOCOL_TIMESTAMP_BELOW]
	inputs, err := method.Inputs.Pack(time)
	if err != nil {
		return nil, err
	}
	return append(method.ID[:], inputs[:]...), nil
}

func (l *LimitOrderProtocolFacadeImpl) NonceEquals(makerAddress common.Address, makerNonce *big.Int) ([]byte, error) {
	method := l.methods[LIMIT_ORDER_PROTOCOL_NONCE_EQUALS]
	inputs, err := method.Inputs.Pack(makerAddress, makerNonce)
	if err != nil {
		return nil, err
	}
	return append(method.ID[:], inputs[:]...), nil
}

func (l *LimitOrderProtocolFacadeImpl) Remaining(orderHash [32]byte) ([]byte, error) {
	method := l.methods[LIMIT_ORDER_PROTOCOL_REMAINING]
	inputs, err := method.Inputs.Pack(orderHash)
	if err != nil {
		return nil, err
	}
	return append(method.ID[:], inputs[:]...), nil
}

func (l *LimitOrderProtocolFacadeImpl) CheckPredicate(order model.LimitOrder) ([]byte, error) {
	method := l.methods[LIMIT_ORDER_PROTOCOL_CHECK_PREDICATE]
	inputs, err := method.Inputs.Pack(order)
	if err != nil {
		return nil, err
	}
	return append(method.ID[:], inputs[:]...), nil
}

func (l *LimitOrderProtocolFacadeImpl) RemainingsRaw(orderHashes [][32]byte) ([]byte, error) {
	method := l.methods[LIMIT_ORDER_PROTOCOL_REMAINING_RAW]
	inputs, err := method.Inputs.Pack(orderHashes)
	if err != nil {
		return nil, err
	}
	return append(method.ID[:], inputs[:]...), nil
}

func (l *LimitOrderProtocolFacadeImpl) DomainSeparator() ([]byte, error) {
	method := l.methods[LIMIT_ORDER_PROTOCOL_DOMAIN_SEPARATOR]
	inputs, err := method.Inputs.Pack()
	if err != nil {
		return nil, err
	}
	return append(method.ID[:], inputs[:]...), nil
}

func (l *LimitOrderProtocolFacadeImpl) BatchFillOrders(fillData *limitOrder.OrdersLimitOrderFillData) ([]byte, error) {
	method := l.methods[LIMIT_ORDER_PROTOCOL_BATCH_FILL_ORDERS]
	inputs, err := method.Inputs.Pack(fillData)
	if err != nil {
		return nil, err
	}
	return append(method.ID[:], inputs[:]...), nil
}
