package iikocloud

import (
	"context"
	"fmt"
	"github.com/themgmd/iikocloud/internal/constants"
	"github.com/themgmd/iikocloud/internal/http"
	"net/url"
)

type CalculateCheckinRequest struct {
	Order                                Order             `json:"order"`
	Coupon                               string            `json:"coupon"`
	ReferrerId                           string            `json:"referrerId"`
	TerminalGroupId                      string            `json:"terminalGroupId"`
	AvailablePaymentMarketingCampaignIds []string          `json:"availablePaymentMarketingCampaignIds"`
	ApplicableManualConditions           []string          `json:"applicableManualConditions"`
	DynamicDiscounts                     []DynamicDiscount `json:"dynamicDiscounts"`
	IsLoyaltyTraceEnabled                bool              `json:"isLoyaltyTraceEnabled"`
	OrganizationId                       string            `json:"organizationId"`
}

type CalculateCheckinResponse struct {
	LoyaltyProgramResults []LoyaltyProgramResult `json:"loyaltyProgramResults"`
	AvailablePayments     []AvailablePayment     `json:"availablePayments"`
	ValidationWarnings    []string               `json:"validationWarnings"`
	Warnings              []Warning              `json:"Warnings"`
	LoyaltyTrace          string                 `json:"loyaltyTrace"`
}

// CalculateCheckin .
func (c *Client) CalculateCheckin(ctx context.Context, req CalculateCheckinRequest, opts ...http.Option) (*CalculateCheckinResponse, error) {
	path, err := url.JoinPath(c.baseURL, constants.V1CalculateCheckin)
	if err != nil {
		return nil, fmt.Errorf("url.JoinPath: %w", err)
	}

	// позволяет перезаписать таймаут если передать опцию снаружи
	opts = append([]http.Option{http.WithCustomTimeout(c.timeout)}, opts...)
	// запрос требует авторизации
	opts = append(opts, http.WithAuthorization(c.token))

	resp, err := http.DoRequest[CalculateCheckinRequest, CalculateCheckinResponse](ctx, c.httpClient, path, req, opts...)
	if err != nil {
		return nil, fmt.Errorf("http.DoRequest: %w", err)
	}

	return resp, nil
}
