package iikocloud

import (
	"context"
	"fmt"
	"net/url"

	"github.com/themgmd/iikocloud/internal/constants"
	"github.com/themgmd/iikocloud/internal/http"
)

type StopListsRequest struct {
	OrganizationIds []string `json:"organizationIds"`
}

type StopListItem struct {
	ProductId    string `json:"productId"`
	ProductName  string `json:"productName"`
	Remainder    int    `json:"remainder"`
	LastPurchase string `json:"lastPurchase"`
	InUse        bool   `json:"inUse"`
}

type StopList struct {
	Id             string        `json:"id"`
	OrganizationId string        `json:"organizationId"`
	Items          []StopListItem `json:"items"`
}

type StopListsResponse struct {
	CorrelationId string     `json:"correlationId"`
	StopLists     []StopList `json:"stopLists"`
}

// StopLists получить стоп-листы организаций
func (c *Client) StopLists(ctx context.Context, req StopListsRequest, opts ...http.Option) (*StopListsResponse, error) {
	path, err := url.JoinPath(c.baseURL, constants.V1StopLists)
	if err != nil {
		return nil, fmt.Errorf("url.JoinPath: %w", err)
	}

	opts = append([]http.Option{http.WithCustomTimeout(c.timeout)}, opts...)
	opts = append(opts, http.WithAuthorization(c.token))

	resp, err := http.DoRequest[StopListsRequest, StopListsResponse](ctx, c.httpClient, path, req, opts...)
	if err != nil {
		return nil, fmt.Errorf("http.DoRequest: %w", err)
	}

	return resp, nil
}

