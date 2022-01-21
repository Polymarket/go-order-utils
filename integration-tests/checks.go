package main

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
)

func marketOrderCheck(ctx context.Context, httpClient *http.Client) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		panic(err)
	}

	credentials, err := getCredentials(ctx, httpClient, privateKey)
	if err != nil {
		panic(err)
	}

	marketOrder, err := buildMarketOrderAndSignature(privateKey)
	if err != nil {
		panic(err)
	}

	marketOrderJson, err := json.Marshal(marketOrder)
	if err != nil {
		panic(err)
	}

	l2Headers, err := createL2Headers(
		privateKey,
		&ApiKeyCreds{
			key:        credentials["apiKey"].(string),
			secret:     credentials["secret"].(string),
			passphrase: credentials["passphrase"].(string),
		},
		&L2HeaderArgs{method: "POST", requestPath: "/order", body: string(marketOrderJson)},
	)
	if err != nil {
		panic(err)
	}

	//payload :=
	createOrderReq, err := generateHTTPRequest(
		ctx,
		"POST",
		fmt.Sprintf("%s/order", TRACKER_URL),
		l2Headers,
		strings.NewReader(string(marketOrderJson)),
	)
	if err != nil {
		panic(err)
	}

	createOrderResp, err := httpClient.Do(createOrderReq)
	if err != nil {
		panic(err)
	}

	if createOrderResp.StatusCode != 200 {
		panic(parseHttpResponseWithError("create market order response status != 200", createOrderResp))
	}
}

func limitOrderCheck(ctx context.Context, httpClient *http.Client) {
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		panic(err)
	}

	credentials, err := getCredentials(ctx, httpClient, privateKey)
	if err != nil {
		panic(err)
	}

	limitOrder, err := buildLimitOrderAndSignature(privateKey)
	if err != nil {
		panic(err)
	}

	limitOrderJson, err := json.Marshal(limitOrder)
	if err != nil {
		panic(err)
	}

	l2Headers, err := createL2Headers(
		privateKey,
		&ApiKeyCreds{
			key:        credentials["apiKey"].(string),
			secret:     credentials["secret"].(string),
			passphrase: credentials["passphrase"].(string),
		},
		&L2HeaderArgs{method: "POST", requestPath: "/order", body: string(limitOrderJson)},
	)
	if err != nil {
		panic(err)
	}

	//payload :=
	createOrderReq, err := generateHTTPRequest(
		ctx,
		"POST",
		fmt.Sprintf("%s/order", TRACKER_URL),
		l2Headers,
		strings.NewReader(string(limitOrderJson)),
	)
	if err != nil {
		panic(err)
	}

	createOrderResp, err := httpClient.Do(createOrderReq)
	if err != nil {
		panic(err)
	}

	if createOrderResp.StatusCode != 200 {
		panic(parseHttpResponseWithError("create limit order response status != 200", createOrderResp))
	}
}

func getCredentials(ctx context.Context, httpClient *http.Client, privateKey *ecdsa.PrivateKey) (map[string]interface{}, error) {
	l1Headers, err := createL1Headers(privateKey)
	if err != nil {
		return nil, err
	}

	createApiKeyReq, err := generateHTTPRequest(
		ctx,
		"POST",
		fmt.Sprintf("%s/create-api-key", TRACKER_URL),
		l1Headers,
		nil,
	)
	if err != nil {
		return nil, err
	}

	createApiKeyResp, err := httpClient.Do(createApiKeyReq)
	if err != nil {
		return nil, err
	}

	if createApiKeyResp.StatusCode != 200 {
		panic(parseHttpResponseWithError("create api key response status != 200", createApiKeyResp))
	}

	credentials, err := readBody(createApiKeyResp)
	if err != nil {
		return nil, err
	}

	return credentials, nil
}

func healthCheck(ctx context.Context, httpClient *http.Client) {
	healthReq, err := generateHTTPRequest(
		ctx,
		"GET",
		fmt.Sprintf("%s/", TRACKER_URL),
		map[string]string{},
		nil,
	)
	if err != nil {
		panic(err)
	}

	healthResp, err := httpClient.Do(healthReq)
	if err != nil {
		panic(err)
	}

	if healthResp.StatusCode != 200 {
		panic(parseHttpResponseWithError("health response status != 200", healthResp))
	}
}
