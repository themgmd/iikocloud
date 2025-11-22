package iikocloud

import (
	"context"
	"fmt"
	"net/url"

	"github.com/themgmd/iikocloud/internal/constants"
	"github.com/themgmd/iikocloud/internal/http"
)

type CouponsRequest struct {
	OrganizationId string `json:"organizationId"`
	Coupon         string `json:"coupon"`
}

type CouponInfo struct {
	Id                        string `json:"id"`
	Name                      string `json:"name"`
	IsActivated               bool   `json:"isActivated"`
	IsValid                   bool   `json:"isValid"`
	ActivatedDate             string `json:"activatedDate,omitempty"`
	ValidFrom                 string `json:"validFrom,omitempty"`
	ValidTo                   string `json:"validTo,omitempty"`
	MarketingCampaignId       string `json:"marketingCampaignId"`
	MarketingCampaignName     string `json:"marketingCampaignName"`
	ProgramId                 string `json:"programId"`
	ProgramName               string `json:"programName"`
	Description               string `json:"description,omitempty"`
}

type CouponsResponse struct {
	CorrelationId string       `json:"correlationId"`
	Coupons       []CouponInfo `json:"coupons"`
}

// Coupons получить информацию о купонах
func (c *Client) Coupons(ctx context.Context, req CouponsRequest, opts ...http.Option) (*CouponsResponse, error) {
	path, err := url.JoinPath(c.baseURL, constants.V1Coupons)
	if err != nil {
		return nil, fmt.Errorf("url.JoinPath: %w", err)
	}

	opts = append([]http.Option{http.WithCustomTimeout(c.timeout)}, opts...)
	opts = append(opts, http.WithAuthorization(c.token))

	resp, err := http.DoRequest[CouponsRequest, CouponsResponse](ctx, c.httpClient, path, req, opts...)
	if err != nil {
		return nil, fmt.Errorf("http.DoRequest: %w", err)
	}

	return resp, nil
}

