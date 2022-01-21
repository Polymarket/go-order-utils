package main

import (
	"crypto/ecdsa"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/polymarket/go-order-utils/pkg/builders"
	"github.com/polymarket/go-order-utils/pkg/config"
	"github.com/polymarket/go-order-utils/pkg/model"
	"github.com/polymarket/go-order-utils/pkg/sign"
)

func buildMarketOrderAndSignature(privateKey *ecdsa.PrivateKey) (*model.MarketOrderAndSignature, error) {
	chainId := 42
	contract, err := config.GetContracts(chainId)
	if err != nil {
		return nil, err
	}

	marketOrderBuilder := builders.NewMarketOrderBuilderImpl(
		common.HexToAddress(contract.Exchange.Address),
		chainId,
		nil,
	)

	signer, err := sign.GetPublicAddress(privateKey)
	if err != nil {
		return nil, err
	}

	/*
		makerAssetAddress common.Address,
		takerAssetAddress common.Address,
		makerAddress common.Address,
		signer common.Address,
		makerAmount *big.Int,
		takerAssetID *big.Int,
		makerAssetID *big.Int,
		sigType int
	*/
	marketOrder := marketOrderBuilder.BuildMarketOrder(
		common.HexToAddress(contract.Collateral),                          // makerAssetAddress common.Address
		common.HexToAddress("0xE7819d9745e64c14541732ca07CC3898670b7650"), // takerAssetAddress common.Address
		signer,                       // makerAddress common.Address
		signer,                       // signer common.Address
		big.NewInt(int64(100000000)), // makerAmount *big.Int
		big.NewInt(int64(0)),         // makerAssetID *big.Int
		big.NewInt(int64(1)),         // takerAssetID *big.Int
		model.EOA,                    // sigType int,
	)

	orderHash, err := marketOrderBuilder.BuildMarketOrderHash(marketOrder)
	if err != nil {
		return nil, err
	}

	signature, err := marketOrderBuilder.BuildSignature(privateKey, orderHash)
	if err != nil {
		return nil, err
	}

	return marketOrderBuilder.BuildMarketOrderAndSignature(marketOrder, orderHash, signature)
}

func buildLimitOrderAndSignature(privateKey *ecdsa.PrivateKey) (*model.LimitOrderAndSignature, error) {
	chainId := 42
	contract, err := config.GetContracts(chainId)
	if err != nil {
		return nil, err
	}

	limitOrderBuilder, err := builders.NewLimitOrderBuilderImpl(
		common.HexToAddress(contract.Exchange.Address),
		chainId,
		nil,
	)
	if err != nil {
		return nil, err
	}

	signer, err := sign.GetPublicAddress(privateKey)
	if err != nil {
		return nil, err
	}

	/*
		(*builders.LimitOrderBuilderImpl).BuildLimitOrder(
			exchangeAddress common.Address,
			makerAssetAddress common.Address,
			makerAddress common.Address,
			takerAssetAddress common.Address,
			takerAddress common.Address,
			signer common.Address,
			permit []byte,
			interaction []byte,
			predicate []byte,
			makerAmount *big.Int,
			takerAmount *big.Int,
			makerAssetID *big.Int,
			takerAssetID *big.Int,
			expiry *big.Int,
			nonce *big.Int,
			sigType int,
		) (*model.LimitOrder, error)
	*/
	limitOrder, err := limitOrderBuilder.BuildLimitOrder(
		common.HexToAddress(contract.Exchange.Address), // exchangeAddress common.Address
		common.HexToAddress(contract.Collateral),       // makerAssetAddress common.Address
		signer,                                         // makerAddress common.Address
		common.HexToAddress("0xE7819d9745e64c14541732ca07CC3898670b7650"), // takerAssetAddress common.Address
		common.HexToAddress("0x0000000000000000000000000000000000000000"), // takerAddress common.Address
		signer,                       // signer common.Address
		[]byte(""),                   // permit []byte
		[]byte(""),                   // interaction []byte,
		[]byte(""),                   //	predicate []byte,
		big.NewInt(int64(100000000)), // makerAmount *big.Int,
		big.NewInt(int64(200000000)), // takerAmount *big.Int,
		big.NewInt(int64(0)),         // makerAssetID *big.Int,
		big.NewInt(int64(1)),         // takerAssetID *big.Int,
		big.NewInt(int64(time.Now().Unix()+int64(60000))), // expiry *big.Int,
		big.NewInt(int64(0)), // nonce *big.Int,
		model.EOA,            // sigType int,
	)
	if err != nil {
		return nil, err
	}

	orderHash, err := limitOrderBuilder.BuildLimitOrderHash(limitOrder)
	if err != nil {
		return nil, err
	}

	signature, err := limitOrderBuilder.BuildSignature(privateKey, orderHash)
	if err != nil {
		return nil, err
	}

	return limitOrderBuilder.BuildLimitOrderAndSignature(limitOrder, orderHash, signature)
}
