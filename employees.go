package iikocloud

import (
	"context"
	"fmt"
	"net/url"

	"github.com/themgmd/iikocloud/internal/constants"
	"github.com/themgmd/iikocloud/internal/http"
)

type EmployeesRequest struct {
	OrganizationIds []string `json:"organizationIds"`
}

type Employee struct {
	Id              string   `json:"id"`
	Name            string   `json:"name"`
	Phone           string   `json:"phone,omitempty"`
	Email           string   `json:"email,omitempty"`
	Position        string   `json:"position,omitempty"`
	TerminalGroupIds []string `json:"terminalGroupIds,omitempty"`
	IsDeleted       bool     `json:"isDeleted"`
}

type EmployeesResponse struct {
	CorrelationId string     `json:"correlationId"`
	Employees     []Employee `json:"employees"`
}

// Employees получить сотрудников
func (c *Client) Employees(ctx context.Context, req EmployeesRequest, opts ...http.Option) (*EmployeesResponse, error) {
	path, err := url.JoinPath(c.baseURL, constants.V1Employees)
	if err != nil {
		return nil, fmt.Errorf("url.JoinPath: %w", err)
	}

	opts = append([]http.Option{http.WithCustomTimeout(c.timeout)}, opts...)
	opts = append(opts, http.WithAuthorization(c.token))

	resp, err := http.DoRequest[EmployeesRequest, EmployeesResponse](ctx, c.httpClient, path, req, opts...)
	if err != nil {
		return nil, fmt.Errorf("http.DoRequest: %w", err)
	}

	return resp, nil
}

