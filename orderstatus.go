package iikocloud

import (
	"context"
	"fmt"
	"net/url"

	"github.com/themgmd/iikocloud/internal/constants"
	"github.com/themgmd/iikocloud/internal/http"
)

type OrderStatusRequest struct {
	OrganizationId string   `json:"organizationId"`
	OrderIds       []string `json:"orderIds"`
}

type OrderStatusResponse struct {
	CorrelationId string      `json:"correlationId"`
	Orders        []OrderInfo `json:"orders"`
}

// OrderStatus получить статус заказов
func (c *Client) OrderStatus(ctx context.Context, req OrderStatusRequest, opts ...http.Option) (*OrderStatusResponse, error) {
	path, err := url.JoinPath(c.baseURL, constants.V1OrderStatus)
	if err != nil {
		return nil, fmt.Errorf("url.JoinPath: %w", err)
	}

	opts = append([]http.Option{http.WithCustomTimeout(c.timeout)}, opts...)
	opts = append(opts, http.WithAuthorization(c.token))

	resp, err := http.DoRequest[OrderStatusRequest, OrderStatusResponse](ctx, c.httpClient, path, req, opts...)
	if err != nil {
		return nil, fmt.Errorf("http.DoRequest: %w", err)
	}

	return resp, nil
}

