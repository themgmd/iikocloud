package iikocloud

import (
	"context"
	"fmt"
	"net/url"

	"github.com/themgmd/iikocloud/internal/constants"
	"github.com/themgmd/iikocloud/internal/http"
)

type CustomerCreateOrUpdateRequest struct {
	OrganizationId string   `json:"organizationId"`
	Customer       Customer `json:"customer"`
}

type CustomerCreateOrUpdateResponse struct {
	CorrelationId string `json:"correlationId"`
	Id            string `json:"id"`
}

// CustomerCreateOrUpdate создать или обновить клиента
func (c *Client) CustomerCreateOrUpdate(ctx context.Context, req CustomerCreateOrUpdateRequest, opts ...http.Option) (*CustomerCreateOrUpdateResponse, error) {
	path, err := url.JoinPath(c.baseURL, constants.V1CustomerCreateOrUpdate)
	if err != nil {
		return nil, fmt.Errorf("url.JoinPath: %w", err)
	}

	opts = append([]http.Option{http.WithCustomTimeout(c.timeout)}, opts...)
	opts = append(opts, http.WithAuthorization(c.token))

	resp, err := http.DoRequest[CustomerCreateOrUpdateRequest, CustomerCreateOrUpdateResponse](ctx, c.httpClient, path, req, opts...)
	if err != nil {
		return nil, fmt.Errorf("http.DoRequest: %w", err)
	}

	return resp, nil
}

