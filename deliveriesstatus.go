package iikocloud

import (
	"context"
	"fmt"
	"net/url"

	"github.com/themgmd/iikocloud/internal/constants"
	"github.com/themgmd/iikocloud/internal/http"
)

type DeliveriesStatusRequest struct {
	OrganizationId string   `json:"organizationId"`
	OrderIds       []string `json:"orderIds"`
}

type DeliveriesStatusResponse struct {
	CorrelationId string      `json:"correlationId"`
	Orders        []OrderInfo `json:"orders"`
}

// DeliveriesStatus получить статус доставок
func (c *Client) DeliveriesStatus(ctx context.Context, req DeliveriesStatusRequest, opts ...http.Option) (*DeliveriesStatusResponse, error) {
	path, err := url.JoinPath(c.baseURL, constants.V1DeliveriesStatus)
	if err != nil {
		return nil, fmt.Errorf("url.JoinPath: %w", err)
	}

	opts = append([]http.Option{http.WithCustomTimeout(c.timeout)}, opts...)
	opts = append(opts, http.WithAuthorization(c.token))

	resp, err := http.DoRequest[DeliveriesStatusRequest, DeliveriesStatusResponse](ctx, c.httpClient, path, req, opts...)
	if err != nil {
		return nil, fmt.Errorf("http.DoRequest: %w", err)
	}

	return resp, nil
}

