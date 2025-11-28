package iikocloud

import (
	"context"
	"fmt"
	"net/url"

	"github.com/themgmd/iikocloud/internal/constants"
	"github.com/themgmd/iikocloud/internal/http"
)

type CustomerWalletChargeOffRequest struct {
	OrganizationId string `json:"organizationId"`
	CustomerId     string `json:"customerId"`
	WalletId       string `json:"walletId"`
	Sum            int    `json:"sum"`
	Comment        string `json:"comment,omitempty"`
}

type CustomerWalletChargeOffResponse struct {
	CorrelationId string `json:"correlationId"`
	TransactionId string `json:"transactionId"`
}

// CustomerWalletChargeOff списать средства с кошелька клиента
func (c *Client) CustomerWalletChargeOff(ctx context.Context, req CustomerWalletChargeOffRequest, opts ...http.Option) (*CustomerWalletChargeOffResponse, error) {
	path, err := url.JoinPath(c.baseURL, constants.V1CustomerWalletChargeOff)
	if err != nil {
		return nil, fmt.Errorf("url.JoinPath: %w", err)
	}

	opts = append([]http.Option{http.WithCustomTimeout(c.timeout)}, opts...)
	opts = append(opts, http.WithAuthorization(c.token))

	resp, err := http.DoRequest[CustomerWalletChargeOffRequest, CustomerWalletChargeOffResponse](ctx, c.httpClient, path, req, opts...)
	if err != nil {
		return nil, fmt.Errorf("http.DoRequest: %w", err)
	}

	return resp, nil
}

