package iikocloud

import (
	"context"
	"fmt"
	"net/url"

	"github.com/themgmd/iikocloud/internal/constants"
	"github.com/themgmd/iikocloud/internal/http"
)

type CustomerWalletRefundRequest struct {
	OrganizationId string `json:"organizationId"`
	CustomerId     string `json:"customerId"`
	WalletId       string `json:"walletId"`
	Sum            int    `json:"sum"`
	Comment        string `json:"comment,omitempty"`
}

type CustomerWalletRefundResponse struct {
	CorrelationId string `json:"correlationId"`
	TransactionId string `json:"transactionId"`
}

// CustomerWalletRefund вернуть средства на кошелек клиента
func (c *Client) CustomerWalletRefund(ctx context.Context, req CustomerWalletRefundRequest, opts ...http.Option) (*CustomerWalletRefundResponse, error) {
	path, err := url.JoinPath(c.baseURL, constants.V1CustomerWalletRefund)
	if err != nil {
		return nil, fmt.Errorf("url.JoinPath: %w", err)
	}

	opts = append([]http.Option{http.WithCustomTimeout(c.timeout)}, opts...)
	opts = append(opts, http.WithAuthorization(c.token))

	resp, err := http.DoRequest[CustomerWalletRefundRequest, CustomerWalletRefundResponse](ctx, c.httpClient, path, req, opts...)
	if err != nil {
		return nil, fmt.Errorf("http.DoRequest: %w", err)
	}

	return resp, nil
}

