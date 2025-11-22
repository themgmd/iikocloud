package iikocloud

import (
	"context"
	"fmt"
	"net/url"

	"github.com/themgmd/iikocloud/internal/constants"
	"github.com/themgmd/iikocloud/internal/http"
)

type CustomerWalletHoldRequest struct {
	OrganizationId string `json:"organizationId"`
	CustomerId     string `json:"customerId"`
	WalletId       string `json:"walletId"`
	Sum            int    `json:"sum"`
	Comment        string `json:"comment,omitempty"`
}

type CustomerWalletHoldResponse struct {
	CorrelationId string `json:"correlationId"`
	TransactionId string `json:"transactionId"`
}

// CustomerWalletHold заблокировать средства на кошельке клиента
func (c *Client) CustomerWalletHold(ctx context.Context, req CustomerWalletHoldRequest, opts ...http.Option) (*CustomerWalletHoldResponse, error) {
	path, err := url.JoinPath(c.baseURL, constants.V1CustomerWalletHold)
	if err != nil {
		return nil, fmt.Errorf("url.JoinPath: %w", err)
	}

	opts = append([]http.Option{http.WithCustomTimeout(c.timeout)}, opts...)
	opts = append(opts, http.WithAuthorization(c.token))

	resp, err := http.DoRequest[CustomerWalletHoldRequest, CustomerWalletHoldResponse](ctx, c.httpClient, path, req, opts...)
	if err != nil {
		return nil, fmt.Errorf("http.DoRequest: %w", err)
	}

	return resp, nil
}

