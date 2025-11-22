package iikocloud

import (
	"context"
	"fmt"
	"net/url"

	"github.com/themgmd/iikocloud/internal/constants"
	"github.com/themgmd/iikocloud/internal/http"
)

type CancelOrderRequest struct {
	OrganizationId string `json:"organizationId"`
	OrderId        string `json:"orderId"`
	Comment        string `json:"comment,omitempty"`
}

type CancelOrderResponse struct {
	CorrelationId string `json:"correlationId"`
}

// CancelOrder отменить заказ
func (c *Client) CancelOrder(ctx context.Context, req CancelOrderRequest, opts ...http.Option) (*CancelOrderResponse, error) {
	path, err := url.JoinPath(c.baseURL, constants.V1CancelOrder)
	if err != nil {
		return nil, fmt.Errorf("url.JoinPath: %w", err)
	}

	opts = append([]http.Option{http.WithCustomTimeout(c.timeout)}, opts...)
	opts = append(opts, http.WithAuthorization(c.token))

	resp, err := http.DoRequest[CancelOrderRequest, CancelOrderResponse](ctx, c.httpClient, path, req, opts...)
	if err != nil {
		return nil, fmt.Errorf("http.DoRequest: %w", err)
	}

	return resp, nil
}

