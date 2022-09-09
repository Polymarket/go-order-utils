package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func generateHTTPRequest(ctx context.Context, method, url string,
	headers map[string]string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequestWithContext(ctx, method, url, body)
	if err != nil {
		fmt.Printf("Error creating %s request for url: %s, err: %s", method, url, err.Error())
		return nil, err
	}

	// encode URL.
	req.URL.RawQuery = req.URL.Query().Encode()

	for key, value := range headers {
		req.Header.Add(key, value)
	}

	return req, nil
}

func parseHttpResponseWithError(msg string, resp *http.Response) string {
	data := map[string]string{
		"where":  msg,
		"status": resp.Status,
	}

	if resp.Body != nil {
		defer resp.Body.Close()

		bytes, err := io.ReadAll(resp.Body)
		if err == nil {
			data["body"] = string(bytes)
		}
	}

	dataB, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	return string(dataB)
}

func readBody(httpResp *http.Response) (map[string]interface{}, error) {
	defer httpResp.Body.Close()

	body, err := io.ReadAll(httpResp.Body)
	if err != nil {
		return nil, err
	}

	resp := map[string]interface{}{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
