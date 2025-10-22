package iikocloud

import (
	"context"
	"fmt"
	"net/url"

	"github.com/themgmd/iikocloud/internal/constants"
	"github.com/themgmd/iikocloud/internal/http"
)

type OrganizationsRequest struct {
	OrganizationIDs      []string `json:"organizationIds"`
	ReturnAdditionalInfo bool     `json:"returnAdditionalInfo"`
	IncludeDisabled      bool     `json:"includeDisabled"`
}

type OrganizationsResponse struct {
	CorrelationID string         `json:"correlationId"`
	Organizations []Organization `json:"organizations"`
}

// Organizations .
func (c *Client) Organizations(ctx context.Context, req OrganizationsRequest, opts ...http.Option) (*OrganizationsResponse, error) {
	path, err := url.JoinPath(c.baseURL, constants.V1Organizations)
	if err != nil {
		return nil, fmt.Errorf("url.JoinPath: %w", err)
	}

	// позволяет перезаписать таймаут если передать опцию снаружи
	opts = append([]http.Option{http.WithCustomTimeout(c.timeout)}, opts...)
	// запрос требует авторизации
	opts = append(opts, http.WithAuthorization(c.token))

	resp, err := http.DoRequest[OrganizationsRequest, OrganizationsResponse](ctx, c.httpClient, path, req, opts...)
	if err != nil {
		return nil, fmt.Errorf("http.DoRequest: %w", err)
	}

	return resp, nil
}
