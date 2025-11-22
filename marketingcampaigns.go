package iikocloud

import (
	"context"
	"fmt"
	"net/url"

	"github.com/themgmd/iikocloud/internal/constants"
	"github.com/themgmd/iikocloud/internal/http"
)

type MarketingCampaignsRequest struct {
	OrganizationIds []string `json:"organizationIds"`
}

type MarketingCampaign struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	IsDeleted   bool   `json:"isDeleted"`
}

type MarketingCampaignsResponse struct {
	CorrelationId      string             `json:"correlationId"`
	MarketingCampaigns []MarketingCampaign `json:"marketingCampaigns"`
}

// MarketingCampaigns получить список маркетинговых кампаний
func (c *Client) MarketingCampaigns(ctx context.Context, req MarketingCampaignsRequest, opts ...http.Option) (*MarketingCampaignsResponse, error) {
	path, err := url.JoinPath(c.baseURL, constants.V1MarketingCampaigns)
	if err != nil {
		return nil, fmt.Errorf("url.JoinPath: %w", err)
	}

	opts = append([]http.Option{http.WithCustomTimeout(c.timeout)}, opts...)
	opts = append(opts, http.WithAuthorization(c.token))

	resp, err := http.DoRequest[MarketingCampaignsRequest, MarketingCampaignsResponse](ctx, c.httpClient, path, req, opts...)
	if err != nil {
		return nil, fmt.Errorf("http.DoRequest: %w", err)
	}

	return resp, nil
}

