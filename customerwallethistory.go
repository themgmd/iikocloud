package iikocloud

import (
	"context"
	"fmt"
	"net/url"

	"github.com/themgmd/iikocloud/internal/constants"
	"github.com/themgmd/iikocloud/internal/http"
)

type CustomerWalletHistoryRequest struct {
	OrganizationId string `json:"organizationId"`
	CustomerId     string `json:"customerId"`
	WalletId       string `json:"walletId,omitempty"`
	DateFrom       string `json:"dateFrom,omitempty"`
	DateTo         string `json:"dateTo,omitempty"`
}

type WalletHistoryItem struct {
	TransactionId string `json:"transactionId"`
	Date          string `json:"date"`
	Type          string `json:"type"`
	Sum           int    `json:"sum"`
	Comment       string `json:"comment,omitempty"`
}

type CustomerWalletHistoryResponse struct {
	CorrelationId string             `json:"correlationId"`
	Items         []WalletHistoryItem `json:"items"`
}

// CustomerWalletHistory получить историю операций кошелька клиента
func (c *Client) CustomerWalletHistory(ctx context.Context, req CustomerWalletHistoryRequest, opts ...http.Option) (*CustomerWalletHistoryResponse, error) {
	path, err := url.JoinPath(c.baseURL, constants.V1CustomerWalletHistory)
	if err != nil {
		return nil, fmt.Errorf("url.JoinPath: %w", err)
	}

	opts = append([]http.Option{http.WithCustomTimeout(c.timeout)}, opts...)
	opts = append(opts, http.WithAuthorization(c.token))

	resp, err := http.DoRequest[CustomerWalletHistoryRequest, CustomerWalletHistoryResponse](ctx, c.httpClient, path, req, opts...)
	if err != nil {
		return nil, fmt.Errorf("http.DoRequest: %w", err)
	}

	return resp, nil
}

