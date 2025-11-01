package iikocloud

import (
	"context"
	"fmt"
	"github.com/themgmd/iikocloud/internal/constants"
	"github.com/themgmd/iikocloud/internal/http"
	"net/url"
)

type TerminalGroupsRequest struct {
	OrganizationIds    []string `json:"organizationIds"`
	IncludeDisabled    bool     `json:"includeDisabled"`
	ReturnExternalData []string `json:"returnExternalData"`
}

type TerminalGroupsResponse struct {
	CorrelationId         string          `json:"correlationId"`
	TerminalGroups        []TerminalGroup `json:"terminalGroups"`
	TerminalGroupsInSleep []TerminalGroup `json:"terminalGroupsInSleep"`
}

// TerminalGroups .
func (c *Client) TerminalGroups(ctx context.Context, req TerminalGroupsRequest, opts ...http.Option) (*TerminalGroupsResponse, error) {
	path, err := url.JoinPath(c.baseURL, constants.V1TerminalGroups)
	if err != nil {
		return nil, fmt.Errorf("url.JoinPath: %w", err)
	}

	// позволяет перезаписать таймаут если передать опцию снаружи
	opts = append([]http.Option{http.WithCustomTimeout(c.timeout)}, opts...)
	// запрос требует авторизации
	opts = append(opts, http.WithAuthorization(c.token))

	resp, err := http.DoRequest[TerminalGroupsRequest, TerminalGroupsResponse](ctx, c.httpClient, path, req, opts...)
	if err != nil {
		return nil, fmt.Errorf("http.DoRequest: %w", err)
	}
	return resp, nil
}
