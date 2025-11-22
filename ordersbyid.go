package iikocloud

import (
	"context"
	"fmt"
	"net/url"

	"github.com/themgmd/iikocloud/internal/constants"
	"github.com/themgmd/iikocloud/internal/http"
)

type OrdersByIDRequest struct {
	OrganizationId string   `json:"organizationId"`
	OrderIds       []string `json:"orderIds"`
}

type OrdersByIDResponse struct {
	CorrelationId string      `json:"correlationId"`
	Orders        []OrderInfo `json:"orders"`
}

// OrdersByID получить информацию о заказах по их ID
func (c *Client) OrdersByID(ctx context.Context, req OrdersByIDRequest, opts ...http.Option) (*OrdersByIDResponse, error) {
	path, err := url.JoinPath(c.baseURL, constants.V1OrdersByID)
	if err != nil {
		return nil, fmt.Errorf("url.JoinPath: %w", err)
	}

	opts = append([]http.Option{http.WithCustomTimeout(c.timeout)}, opts...)
	opts = append(opts, http.WithAuthorization(c.token))

	resp, err := http.DoRequest[OrdersByIDRequest, OrdersByIDResponse](ctx, c.httpClient, path, req, opts...)
	if err != nil {
		return nil, fmt.Errorf("http.DoRequest: %w", err)
	}

	return resp, nil
}

