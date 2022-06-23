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
	makerAmount, _ := big.NewInt(0).SetString("20000000000000000000", 10)
	makerAssetID, _ := big.NewInt(0).SetString("26319442990640912640927597566187363896165058256070126954393241152928258958718", 10)
	marketOrder := marketOrderBuilder.BuildMarketOrder(
		common.HexToAddress("0x3E89215CC47670F084cB9f31BBAA6B5E99e45d0d"), // makerAssetAddress common.Address
		common.HexToAddress("0xe22da380ee6b445bb8273c81944adeb6e8450422"), // takerAssetAddress common.Address
		signer,                // makerAddress common.Address
		signer,                // signer common.Address
		makerAmount,           //big.NewInt(int64(10000000)), // makerAmount *big.Int
		big.NewInt(int64(-1)), // takerAssetID *big.Int
		makerAssetID,          // makerAssetID *big.Int
		model.EOA,             // sigType int,
	)

	orderHash, err := marketOrderBuilder.BuildMarketOrderHash(marketOrder)
	if err != nil {
		return nil, err
	}

	signature, err := marketOrderBuilder.BuildSignature(privateKey, orderHash)
	if err != nil {
		return nil, err
	}

	return marketOrderBuilder.BuildMarketOrderAndSignature(marketOrder, orderHash, signature, "0", "FOK")
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
			exchangeAddress,
			makerAssetAddress,
			makerAddress,
			takerAssetAddress,
			takerAddress,
			signer common.Address,
			permit,
			interaction,
			predicate []byte,
			makerAmount,
			takerAmount,
			makerAssetID,
			takerAssetID,
			expiry,
			nonce *big.Int,
			sigType model.SignatureType,
		) (*model.LimitOrder, error)
	*/
	takerAmount, _ := big.NewInt(0).SetString("20000000000000000000", 10)
	takerAssetID, _ := big.NewInt(0).SetString("26319442990640912640927597566187363896165058256070126954393241152928258958718", 10)
	limitOrder, err := limitOrderBuilder.BuildLimitOrder(
		common.HexToAddress(contract.Exchange.Address),                    // exchangeAddress common.Address
		common.HexToAddress("0xe22da380ee6b445bb8273c81944adeb6e8450422"), // makerAssetAddress common.Address
		signer, // makerAddress common.Address
		common.HexToAddress("0x3E89215CC47670F084cB9f31BBAA6B5E99e45d0d"), // takerAssetAddress common.Address
		common.HexToAddress("0x0000000000000000000000000000000000000000"), // takerAddress common.Address
		signer,                      // signer common.Address
		[]byte(""),                  //	predicate []byte,
		big.NewInt(int64(10000000)), // makerAmount *big.Int,
		takerAmount,                 // takerAmount *big.Int,
		big.NewInt(int64(-1)),       // makerAssetID *big.Int,
		takerAssetID,                // takerAssetID *big.Int,
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
