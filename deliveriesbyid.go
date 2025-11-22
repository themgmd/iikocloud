package iikocloud

import (
	"context"
	"fmt"
	"net/url"

	"github.com/themgmd/iikocloud/internal/constants"
	"github.com/themgmd/iikocloud/internal/http"
)

type DeliveriesByIDRequest struct {
	OrganizationId string   `json:"organizationId"`
	OrderIds       []string `json:"orderIds"`
}

type DeliveriesByIDResponse struct {
	CorrelationId string      `json:"correlationId"`
	Orders        []OrderInfo `json:"orders"`
}

// DeliveriesByID получить информацию о доставках по их ID
func (c *Client) DeliveriesByID(ctx context.Context, req DeliveriesByIDRequest, opts ...http.Option) (*DeliveriesByIDResponse, error) {
	path, err := url.JoinPath(c.baseURL, constants.V1DeliveriesByID)
	if err != nil {
		return nil, fmt.Errorf("url.JoinPath: %w", err)
	}

	opts = append([]http.Option{http.WithCustomTimeout(c.timeout)}, opts...)
	opts = append(opts, http.WithAuthorization(c.token))

	resp, err := http.DoRequest[DeliveriesByIDRequest, DeliveriesByIDResponse](ctx, c.httpClient, path, req, opts...)
	if err != nil {
		return nil, fmt.Errorf("http.DoRequest: %w", err)
	}

	return resp, nil
}

