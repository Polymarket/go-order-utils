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
type CannonicalLimitOrder struct {
	Salt           int    `json:"salt"`
	MakerAsset     string `json:"makerAsset"`
	TakerAsset     string `json:"takerAsset"`
	MakerAssetData string `json:"makerAssetData"`
	TakerAssetData string `json:"takerAssetData"`
	GetMakerAmount string `json:"getMakerAmount"`
	GetTakerAmount string `json:"getTakerAmount"`
	Predicate      string `json:"predicate"`
	Permit         string `json:"permit"`
	Interaction    string `json:"interaction"`
	Signer         string `json:"signer"`
	SigType        int    `json:"sigType"`
}

type LimitOrderAndSignature struct {
	Order     *CannonicalLimitOrder `json:"order"`
	Signature string                `json:"signature"`
	OrderType string                `json:"orderType"`
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

type CannonicalMarketOrder struct {
	Salt         int    `json:"salt"`
	Signer       string `json:"signer"`
	Maker        string `json:"maker"`
	MakerAsset   string `json:"makerAsset"`
	MakerAmount  string `json:"makerAmount"`
	MakerAssetID string `json:"makerAssetID"`
	TakerAsset   string `json:"takerAsset"`
	TakerAssetID string `json:"takerAssetID"`
	SigType      int    `json:"sigType"`
}

type MarketOrderAndSignature struct {
	Order     *CannonicalMarketOrder `json:"order"`
	Signature string                 `json:"signature"`
	OrderType string                 `json:"orderType"`
}

type SignatureType = int

const (
	EOA        SignatureType = iota // EIP712 signatures signed by EOAs
	CONTRACT                        // EIP1271 signatures signed by smart contracts
	POLY_PROXY                      // EIP712 signatures signed by polymarket proxy wallets
)
