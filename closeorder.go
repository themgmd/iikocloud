package iikocloud

import (
	"context"
	"fmt"
	"net/url"

	"github.com/themgmd/iikocloud/internal/constants"
	"github.com/themgmd/iikocloud/internal/http"
)

type CloseOrderRequest struct {
	ChequeAdditionalInfo ChequeAdditionalInfo `json:"chequeAdditionalInfo"`
	OrganizationId       string               `json:"organizationId"`
	OrderId              string               `json:"orderId"`
}

type CloseOrderResponse struct {
	CorrelationId string `json:"correlationId"`
}

// CloseOrder закрыть заказ
func (c *Client) CloseOrder(ctx context.Context, req CloseOrderRequest, opts ...http.Option) (*CloseOrderResponse, error) {
	path, err := url.JoinPath(c.baseURL, constants.V1CloseOrder)
	if err != nil {
		return nil, fmt.Errorf("url.JoinPath: %w", err)
	}

	opts = append([]http.Option{http.WithCustomTimeout(c.timeout)}, opts...)
	opts = append(opts, http.WithAuthorization(c.token))

	resp, err := http.DoRequest[CloseOrderRequest, CloseOrderResponse](ctx, c.httpClient, path, req, opts...)
	if err != nil {
		return nil, fmt.Errorf("http.DoRequest: %w", err)
	}

	return resp, nil
}
