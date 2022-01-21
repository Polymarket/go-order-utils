package main

import (
	"context"
	"fmt"
	"net/http"
	"os"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/joho/godotenv"
)

const (
	TRACKER_URL = "http://localhost:8080"
)

func main() {
	fmt.Println("running...")
	setEnvVars()

	ctx := context.Background()
	env := getEnvVars()
	httpClient := &http.Client{}

	privateKey, err := crypto.HexToECDSA(env["privateKey"])
	if err != nil {
		panic(err)
	}

	credentials := map[string]interface{}{
		"apiKey":     env["apiKey"],
		"passphrase": env["passphrase"],
		"secret":     env["secret"],
	}

	healthCheck(ctx, httpClient)
	limitOrderCheck(ctx, httpClient, credentials, privateKey)
	marketOrderCheck(ctx, httpClient, credentials, privateKey)
}

func setEnvVars() {
	err := godotenv.Load()
	if err != nil {
		panic(fmt.Errorf("error loading .env file!: %s", err.Error()))
	}
}

func getEnvVars() map[string]string {
	privateKey, ok := os.LookupEnv("GO_ORDER_UTILS_WALLET_PRIVATE_KEY")
	if !ok {
		panic("GO_ORDER_UTILS_WALLET_PRIVATE_KEY not set")
	}

	apiKey, ok := os.LookupEnv("GO_ORDER_UTILS_API_KEY")
	if !ok {
		panic("GO_ORDER_UTILS_API_KEY not set")
	}

	passphrase, ok := os.LookupEnv("GO_ORDER_UTILS_PASSPHRASE")
	if !ok {
		panic("GO_ORDER_UTILS_PASSPHRASE not set")
	}

	secret, ok := os.LookupEnv("GO_ORDER_UTILS_SECRET")
	if !ok {
		panic("GO_ORDER_UTILS_SECRET not set")
	}

	return map[string]string{
		"privateKey": privateKey,
		"apiKey":     apiKey,
		"passphrase": passphrase,
		"secret":     secret,
	}
}
