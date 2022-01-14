package order

import "github.com/polymarket/order-utils/pkg/signature"

type MarketOrder struct {
	Salt         int                     `json:"salt"`
	Signer       string                  `json:"signer"`
	Maker        string                  `json:"maker"`
	MakerAsset   string                  `json:"makerAsset"`
	MakerAmount  string                  `json:"makerAmount"`
	MakerAssetID int                     `json:"makerAssetID"`
	TakerAsset   string                  `json:"takerAsset"`
	TakerAssetID int                     `json:"takerAssetID"`
	SigType      signature.SignatureType `json:"sigType"`
}
