package iikocloud

import (
	"context"
	"fmt"
	"net/url"

	"github.com/themgmd/iikocloud/internal/constants"
	"github.com/themgmd/iikocloud/internal/http"
)

type TipsTypesRequest struct {
	OrganizationIds []string `json:"organizationIds"`
}

type TipsType struct {
	Id                           string   `json:"id"`
	Name                         string   `json:"name"`
	Comment                      string   `json:"comment,omitempty"`
	ApplicableMarketingCampaigns []string `json:"applicableMarketingCampaigns,omitempty"`
	IsDeleted                    bool     `json:"isDeleted"`
}

type TipsTypesResponse struct {
	CorrelationId string     `json:"correlationId"`
	TipsTypes     []TipsType `json:"tipsTypes"`
}

// TipsTypes получить типы чаевых
func (c *Client) TipsTypes(ctx context.Context, req TipsTypesRequest, opts ...http.Option) (*TipsTypesResponse, error) {
	path, err := url.JoinPath(c.baseURL, constants.V1TipsTypes)
	if err != nil {
		return nil, fmt.Errorf("url.JoinPath: %w", err)
	}

	opts = append([]http.Option{http.WithCustomTimeout(c.timeout)}, opts...)
	opts = append(opts, http.WithAuthorization(c.token))

	resp, err := http.DoRequest[TipsTypesRequest, TipsTypesResponse](ctx, c.httpClient, path, req, opts...)
	if err != nil {
		return nil, fmt.Errorf("http.DoRequest: %w", err)
	}

	return resp, nil
}

