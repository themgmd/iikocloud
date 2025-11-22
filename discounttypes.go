package iikocloud

import (
	"context"
	"fmt"
	"net/url"

	"github.com/themgmd/iikocloud/internal/constants"
	"github.com/themgmd/iikocloud/internal/http"
)

type DiscountTypesRequest struct {
	OrganizationIds []string `json:"organizationIds"`
}

type DiscountType struct {
	Id                           string   `json:"id"`
	Name                         string   `json:"name"`
	Comment                      string   `json:"comment,omitempty"`
	ApplicableMarketingCampaigns []string `json:"applicableMarketingCampaigns,omitempty"`
	IsDeleted                    bool     `json:"isDeleted"`
}

type DiscountTypesResponse struct {
	CorrelationId string         `json:"correlationId"`
	DiscountTypes []DiscountType `json:"discountTypes"`
}

// DiscountTypes получить типы скидок
func (c *Client) DiscountTypes(ctx context.Context, req DiscountTypesRequest, opts ...http.Option) (*DiscountTypesResponse, error) {
	path, err := url.JoinPath(c.baseURL, constants.V1DiscountTypes)
	if err != nil {
		return nil, fmt.Errorf("url.JoinPath: %w", err)
	}

	opts = append([]http.Option{http.WithCustomTimeout(c.timeout)}, opts...)
	opts = append(opts, http.WithAuthorization(c.token))

	resp, err := http.DoRequest[DiscountTypesRequest, DiscountTypesResponse](ctx, c.httpClient, path, req, opts...)
	if err != nil {
		return nil, fmt.Errorf("http.DoRequest: %w", err)
	}

	return resp, nil
}

