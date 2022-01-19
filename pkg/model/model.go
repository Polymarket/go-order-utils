package model

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type LimitOrder struct {
	Salt           *big.Int       `json:"salt"`
	MakerAsset     common.Address `json:"makerAsset"`
	TakerAsset     common.Address `json:"takerAsset"`
	MakerAssetData []byte         `json:"makerAssetData"`
	TakerAssetData []byte         `json:"takerAssetData"`
	GetMakerAmount []byte         `json:"getMakerAmount"`
	GetTakerAmount []byte         `json:"getTakerAmount"`
	Predicate      []byte         `json:"predicate"`
	Permit         []byte         `json:"permit"`
	Interaction    []byte         `json:"interaction"`
	Signer         common.Address `json:"signer"`
	SigType        *big.Int       `json:"sigType"`
}

type LimitOrderAndSignature struct {
	Order     *LimitOrder `json:"order"`
	Signature []byte      `json:"signature"`
	OrderType string      `json:"orderType"`
}

type MarketOrder struct {
	Salt         *big.Int       `json:"salt"`
	Signer       common.Address `json:"signer"`
	Maker        common.Address `json:"maker"`
	MakerAsset   common.Address `json:"makerAsset"`
	MakerAmount  *big.Int       `json:"makerAmount"`
	MakerAssetID *big.Int       `json:"makerAssetID"`
	TakerAsset   common.Address `json:"takerAsset"`
	TakerAssetID *big.Int       `json:"takerAssetID"`
	SigType      *big.Int       `json:"sigType"`
}

type MarketOrderAndSignature struct {
	Order     *MarketOrder `json:"order"`
	Signature []byte       `json:"signature"`
	OrderType string       `json:"orderType"`
}

type SignatureType = int

const (
	EOA        SignatureType = iota // EIP712 signatures signed by EOAs
	CONTRACT                        // EIP1271 signatures signed by smart contracts
	POLY_PROXY                      // EIP712 signatures signed by polymarket proxy wallets
)
