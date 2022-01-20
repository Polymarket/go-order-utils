package main

import (
	"context"
	"fmt"
	"net/http"
)

const (
	TRACKER_URL = "http://localhost:8080"
)

func main() {
	fmt.Println("running...")
	ctx := context.Background()
	httpClient := &http.Client{}

	healthCheck(ctx, httpClient)
	limitOrderCheck(ctx, httpClient)
}
