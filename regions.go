package iikocloud

import (
	"context"
	"fmt"
	"net/url"

	"github.com/themgmd/iikocloud/internal/constants"
	"github.com/themgmd/iikocloud/internal/http"
)

type RegionsRequest struct {
	OrganizationIds []string `json:"organizationIds"`
}

type RegionsResponse struct {
	CorrelationId string   `json:"correlationId"`
	Regions       []Region `json:"regions"`
}

// Regions получить список регионов
func (c *Client) Regions(ctx context.Context, req RegionsRequest, opts ...http.Option) (*RegionsResponse, error) {
	path, err := url.JoinPath(c.baseURL, constants.V1Regions)
	if err != nil {
		return nil, fmt.Errorf("url.JoinPath: %w", err)
	}

	opts = append([]http.Option{http.WithCustomTimeout(c.timeout)}, opts...)
	opts = append(opts, http.WithAuthorization(c.token))

	resp, err := http.DoRequest[RegionsRequest, RegionsResponse](ctx, c.httpClient, path, req, opts...)
	if err != nil {
		return nil, fmt.Errorf("http.DoRequest: %w", err)
	}

	return resp, nil
}
