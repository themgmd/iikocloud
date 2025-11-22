package iikocloud

import (
	"context"
	"fmt"
	"net/url"

	"github.com/themgmd/iikocloud/internal/constants"
	"github.com/themgmd/iikocloud/internal/http"
)

type OrderTypesRequest struct {
	OrganizationIds []string `json:"organizationIds"`
}

type OrderType struct {
	Id               string `json:"id"`
	Name             string `json:"name"`
	OrderServiceType string `json:"orderServiceType"`
}

type OrderTypesResponse struct {
	CorrelationId string      `json:"correlationId"`
	OrderTypes    []OrderType `json:"orderTypes"`
}

// OrderTypes получить типы заказов
func (c *Client) OrderTypes(ctx context.Context, req OrderTypesRequest, opts ...http.Option) (*OrderTypesResponse, error) {
	path, err := url.JoinPath(c.baseURL, constants.V1OrderTypes)
	if err != nil {
		return nil, fmt.Errorf("url.JoinPath: %w", err)
	}

	opts = append([]http.Option{http.WithCustomTimeout(c.timeout)}, opts...)
	opts = append(opts, http.WithAuthorization(c.token))

	resp, err := http.DoRequest[OrderTypesRequest, OrderTypesResponse](ctx, c.httpClient, path, req, opts...)
	if err != nil {
		return nil, fmt.Errorf("http.DoRequest: %w", err)
	}

	return resp, nil
}

