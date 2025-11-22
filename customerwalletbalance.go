package iikocloud

import (
	"context"
	"fmt"
	"net/url"

	"github.com/themgmd/iikocloud/internal/constants"
	"github.com/themgmd/iikocloud/internal/http"
)

type CustomerWalletBalanceRequest struct {
	OrganizationId string `json:"organizationId"`
	CustomerId     string `json:"customerId"`
}

type WalletBalance struct {
	WalletId string `json:"walletId"`
	Balance  int    `json:"balance"`
}

type CustomerWalletBalanceResponse struct {
	CorrelationId string          `json:"correlationId"`
	Balances      []WalletBalance `json:"balances"`
}

// CustomerWalletBalance получить баланс кошелька клиента
func (c *Client) CustomerWalletBalance(ctx context.Context, req CustomerWalletBalanceRequest, opts ...http.Option) (*CustomerWalletBalanceResponse, error) {
	path, err := url.JoinPath(c.baseURL, constants.V1CustomerWalletBalance)
	if err != nil {
		return nil, fmt.Errorf("url.JoinPath: %w", err)
	}

	opts = append([]http.Option{http.WithCustomTimeout(c.timeout)}, opts...)
	opts = append(opts, http.WithAuthorization(c.token))

	resp, err := http.DoRequest[CustomerWalletBalanceRequest, CustomerWalletBalanceResponse](ctx, c.httpClient, path, req, opts...)
	if err != nil {
		return nil, fmt.Errorf("http.DoRequest: %w", err)
	}

	return resp, nil
}

