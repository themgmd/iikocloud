package iikocloud

import (
	"context"
	"fmt"
	"net/url"

	"github.com/themgmd/iikocloud/internal/constants"
	"github.com/themgmd/iikocloud/internal/http"
)

// AccessTokenRequest .
type AccessTokenRequest struct {
	ApiLogin string `json:"apiLogin"`
}

// AccessTokenResponse .
type AccessTokenResponse struct {
	CorrelationID string `json:"correlationId"`
	Token         string `json:"token"`
}

func (c *Client) accessToken(req AccessTokenRequest, opts ...http.Option) (*AccessTokenResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), c.timeout)
	defer cancel()

	path, err := url.JoinPath(c.baseURL, constants.V1AccessTokenEndpoint)
	if err != nil {
		return nil, fmt.Errorf("url.JoinPath: %w", err)
	}

	// позволяет перезаписать таймаут если передать опцию снаружи
	opts = append([]http.Option{http.WithCustomTimeout(c.timeout)}, opts...)

	resp, err := http.DoRequest[AccessTokenRequest, AccessTokenResponse](ctx, c.httpClient, path, req, opts...)
	if err != nil {
		return nil, fmt.Errorf("DoRequest: %w", err)
	}

	return resp, err
}
