package iikocloud

import (
	"context"
	"fmt"
	"net/url"

	"github.com/themgmd/iikocloud/internal/constants"
	"github.com/themgmd/iikocloud/internal/http"
)

type OperatorsRequest struct {
	OrganizationIds []string `json:"organizationIds"`
}

type Operator struct {
	Id              string   `json:"id"`
	Name            string   `json:"name"`
	Phone           string   `json:"phone,omitempty"`
	Email           string   `json:"email,omitempty"`
	TerminalGroupIds []string `json:"terminalGroupIds,omitempty"`
	IsDeleted       bool     `json:"isDeleted"`
}

type OperatorsResponse struct {
	CorrelationId string     `json:"correlationId"`
	Operators     []Operator `json:"operators"`
}

// Operators получить операторов
func (c *Client) Operators(ctx context.Context, req OperatorsRequest, opts ...http.Option) (*OperatorsResponse, error) {
	path, err := url.JoinPath(c.baseURL, constants.V1Operators)
	if err != nil {
		return nil, fmt.Errorf("url.JoinPath: %w", err)
	}

	opts = append([]http.Option{http.WithCustomTimeout(c.timeout)}, opts...)
	opts = append(opts, http.WithAuthorization(c.token))

	resp, err := http.DoRequest[OperatorsRequest, OperatorsResponse](ctx, c.httpClient, path, req, opts...)
	if err != nil {
		return nil, fmt.Errorf("http.DoRequest: %w", err)
	}

	return resp, nil
}

