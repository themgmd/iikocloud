package iikocloud

import (
	"context"
	"fmt"
	"net/url"

	"github.com/themgmd/iikocloud/internal/constants"
	"github.com/themgmd/iikocloud/internal/http"
)

type CustomerByIDRequest struct {
	OrganizationId string `json:"organizationId"`
	Id             string `json:"id"`
	Type           string `json:"type,omitempty"`
}

type Customer struct {
	Id                            string   `json:"id"`
	Name                          string   `json:"name,omitempty"`
	Surname                       string   `json:"surname,omitempty"`
	Comment                       string   `json:"comment,omitempty"`
	Gender                        string   `json:"gender,omitempty"`
	InBlacklist                   bool     `json:"inBlacklist,omitempty"`
	BlacklistReason               string   `json:"blacklistReason,omitempty"`
	Birthdate                     string   `json:"birthdate,omitempty"`
	Type                          string   `json:"type,omitempty"`
	Email                         string   `json:"email,omitempty"`
	Phone                         string   `json:"phone,omitempty"`
	ShouldReceivePromoActionsInfo bool     `json:"shouldReceivePromoActionsInfo,omitempty"`
	ShouldReceiveOrderStatusNotifications bool `json:"shouldReceiveOrderStatusNotifications,omitempty"`
}

type CustomerByIDResponse struct {
	CorrelationId string   `json:"correlationId"`
	Customers     []Customer `json:"customers"`
}

// CustomerByID получить информацию о клиенте по его ID
func (c *Client) CustomerByID(ctx context.Context, req CustomerByIDRequest, opts ...http.Option) (*CustomerByIDResponse, error) {
	path, err := url.JoinPath(c.baseURL, constants.V1CustomerByID)
	if err != nil {
		return nil, fmt.Errorf("url.JoinPath: %w", err)
	}

	opts = append([]http.Option{http.WithCustomTimeout(c.timeout)}, opts...)
	opts = append(opts, http.WithAuthorization(c.token))

	resp, err := http.DoRequest[CustomerByIDRequest, CustomerByIDResponse](ctx, c.httpClient, path, req, opts...)
	if err != nil {
		return nil, fmt.Errorf("http.DoRequest: %w", err)
	}

	return resp, nil
}

