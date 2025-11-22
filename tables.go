package iikocloud

import (
	"context"
	"fmt"
	"net/url"

	"github.com/themgmd/iikocloud/internal/constants"
	"github.com/themgmd/iikocloud/internal/http"
)

type TablesRequest struct {
	OrganizationIds []string `json:"organizationIds"`
	TerminalGroupIds []string `json:"terminalGroupIds,omitempty"`
}

type TableInfo struct {
	Id              string `json:"id"`
	Number          int    `json:"number"`
	Name            string `json:"name"`
	SeatingCapacity int    `json:"seatingCapacity"`
	Revision        int    `json:"revision"`
	IsDeleted       bool   `json:"isDeleted"`
	RestaurantSectionId string `json:"restaurantSectionId,omitempty"`
	TerminalGroupId string `json:"terminalGroupId,omitempty"`
}

type TablesResponse struct {
	CorrelationId string      `json:"correlationId"`
	Tables        []TableInfo `json:"tables"`
}

// Tables получить список столов
func (c *Client) Tables(ctx context.Context, req TablesRequest, opts ...http.Option) (*TablesResponse, error) {
	path, err := url.JoinPath(c.baseURL, constants.V1Tables)
	if err != nil {
		return nil, fmt.Errorf("url.JoinPath: %w", err)
	}

	opts = append([]http.Option{http.WithCustomTimeout(c.timeout)}, opts...)
	opts = append(opts, http.WithAuthorization(c.token))

	resp, err := http.DoRequest[TablesRequest, TablesResponse](ctx, c.httpClient, path, req, opts...)
	if err != nil {
		return nil, fmt.Errorf("http.DoRequest: %w", err)
	}

	return resp, nil
}

