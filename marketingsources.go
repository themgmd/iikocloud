package iikocloud

import (
	"context"
	"fmt"
	"net/url"

	"github.com/themgmd/iikocloud/internal/constants"
	"github.com/themgmd/iikocloud/internal/http"
)

type MarketingSourcesRequest struct {
	OrganizationIds []string `json:"organizationIds"`
}

type MarketingSource struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	IsDeleted bool   `json:"isDeleted"`
}

type MarketingSourcesResponse struct {
	CorrelationId    string            `json:"correlationId"`
	MarketingSources []MarketingSource `json:"marketingSources"`
}

// MarketingSources получить источники маркетинга
func (c *Client) MarketingSources(ctx context.Context, req MarketingSourcesRequest, opts ...http.Option) (*MarketingSourcesResponse, error) {
	path, err := url.JoinPath(c.baseURL, constants.V1MarketingSources)
	if err != nil {
		return nil, fmt.Errorf("url.JoinPath: %w", err)
	}

	opts = append([]http.Option{http.WithCustomTimeout(c.timeout)}, opts...)
	opts = append(opts, http.WithAuthorization(c.token))

	resp, err := http.DoRequest[MarketingSourcesRequest, MarketingSourcesResponse](ctx, c.httpClient, path, req, opts...)
	if err != nil {
		return nil, fmt.Errorf("http.DoRequest: %w", err)
	}

	return resp, nil
}

