package iikocloud

import (
	"context"
	"fmt"
	"github.com/themgmd/iikocloud/internal/constants"
	"github.com/themgmd/iikocloud/internal/http"
	"net/url"
)

type PaymentTypesRequest struct {
	OrganizationIDs []string `json:"organizationIds"`
}

type PaymentTypesResponse struct {
	CorrelationId string        `json:"correlationId"`
	PaymentTypes  []PaymentType `json:"paymentTypes"`
}

// PaymentTypes .
func (c *Client) PaymentTypes(ctx context.Context, req PaymentTypesRequest, opts ...http.Option) (*PaymentTypesResponse, error) {
	path, err := url.JoinPath(c.baseURL, constants.V1PaymentTypes)
	if err != nil {
		return nil, fmt.Errorf("url.JoinPath: %w", err)
	}

	// позволяет перезаписать таймаут если передать опцию снаружи
	opts = append([]http.Option{http.WithCustomTimeout(c.timeout)}, opts...)
	// запрос требует авторизации
	opts = append(opts, http.WithAuthorization(c.token))

	resp, err := http.DoRequest[PaymentTypesRequest, PaymentTypesResponse](ctx, c.httpClient, path, req, opts...)
	if err != nil {
		return nil, fmt.Errorf("http.DoRequest: %w", err)
	}

	return resp, nil
}
