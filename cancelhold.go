package iikocloud

import (
	"context"
	"fmt"
	"github.com/themgmd/iikocloud/internal/constants"
	"github.com/themgmd/iikocloud/internal/http"
	"net/url"
)

type CancelHoldRequest struct {
	TransactionId  string `json:"transactionId"`
	OrganizationId string `json:"organizationId"`
}

type CancelHoldResponse struct{}

// CancelHold .
func (c *Client) CancelHold(ctx context.Context, req CancelHoldRequest, opts ...http.Option) (*CancelHoldResponse, error) {
	path, err := url.JoinPath(c.baseURL, constants.V1CancelHoldMoney)
	if err != nil {
		return nil, fmt.Errorf("url.JoinPath: %w", err)
	}

	// позволяет перезаписать таймаут если передать опцию снаружи
	opts = append([]http.Option{http.WithCustomTimeout(c.timeout)}, opts...)
	// запрос требует авторизации
	opts = append(opts, http.WithAuthorization(c.token))

	resp, err := http.DoRequest[CancelHoldRequest, CancelHoldResponse](ctx, c.httpClient, path, req, opts...)
	if err != nil {
		return nil, fmt.Errorf("http.DoRequest: %w", err)
	}

	return resp, nil
}
