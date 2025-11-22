package iikocloud

import (
	"context"
	"fmt"
	"net/url"

	"github.com/themgmd/iikocloud/internal/constants"
	"github.com/themgmd/iikocloud/internal/http"
)

type CustomerSearchRequest struct {
	OrganizationId string `json:"organizationId"`
	Type           string `json:"type,omitempty"`
	Phone          string `json:"phone,omitempty"`
	Email          string `json:"email,omitempty"`
	CardTrack      string `json:"cardTrack,omitempty"`
	CardNumber     string `json:"cardNumber,omitempty"`
	Name           string `json:"name,omitempty"`
}

type CustomerSearchResponse struct {
	CorrelationId string     `json:"correlationId"`
	Customers     []Customer `json:"customers"`
}

// CustomerSearch поиск клиентов
func (c *Client) CustomerSearch(ctx context.Context, req CustomerSearchRequest, opts ...http.Option) (*CustomerSearchResponse, error) {
	path, err := url.JoinPath(c.baseURL, constants.V1CustomerSearch)
	if err != nil {
		return nil, fmt.Errorf("url.JoinPath: %w", err)
	}

	opts = append([]http.Option{http.WithCustomTimeout(c.timeout)}, opts...)
	opts = append(opts, http.WithAuthorization(c.token))

	resp, err := http.DoRequest[CustomerSearchRequest, CustomerSearchResponse](ctx, c.httpClient, path, req, opts...)
	if err != nil {
		return nil, fmt.Errorf("http.DoRequest: %w", err)
	}

	return resp, nil
}

