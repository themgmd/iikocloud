package http

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"github.com/themgmd/iikocloud/internal/constants"
	"net/http"
)

// DoRequest make http request with options
func DoRequest[B, R any](ctx context.Context, client *http.Client, endpoint constants.Endpoint, body B, opts ...Option) (*R, error) {
	jsonBody, err := json.Marshal(body)
	if err != nil {
		return nil, fmt.Errorf(" json.Marshal: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, endpoint, bytes.NewBuffer(jsonBody))
	if err != nil {
		return nil, fmt.Errorf("http.NewRequest: %w", err)
	}

	req.Header.Set("Content-Type", "application/json; charset=UTF-8")

	for _, opt := range opts {
		opt(req)
	}

	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("client.Do: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode >= 400 {
		var errorResponse ErrorResponse
		if err = json.NewDecoder(resp.Body).Decode(&errorResponse); err != nil {
			return nil, err
		}

		errorResponse.StatusCode = resp.StatusCode
		return nil, errorResponse
	}

	var response R
	if err = json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, fmt.Errorf(" json.Decode: %w", err)
	}

	return &response, nil
}
