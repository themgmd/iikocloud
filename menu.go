package iikocloud

import (
	"context"
	"fmt"
	"github.com/themgmd/iikocloud/internal/constants"
	"github.com/themgmd/iikocloud/internal/http"
	"net/url"
)

// MenuRequest .
type MenuRequest struct {
}

// MenuResponse .
type MenuResponse struct {
	CorrelationId   string          `json:"correlationId"`
	ExternalMenus   []ExternalMenu  `json:"externalMenus"`
	PriceCategories []PriceCategory `json:"priceCategories"`
}

// V2Menu .
func (c *Client) V2Menu(ctx context.Context, req MenuRequest, opts ...http.Option) (*MenuResponse, error) {
	path, err := url.JoinPath(c.baseURL, constants.V2Menu)
	if err != nil {
		return nil, fmt.Errorf("url.JoinPath: %w", err)
	}

	// позволяет перезаписать таймаут если передать опцию снаружи
	opts = append([]http.Option{http.WithCustomTimeout(c.timeout)}, opts...)
	// запрос требует авторизации
	opts = append(opts, http.WithAuthorization(c.token))

	resp, err := http.DoRequest[MenuRequest, MenuResponse](ctx, c.httpClient, path, req, opts...)
	if err != nil {
		return nil, fmt.Errorf("http.DoRequest: %w", err)
	}
	return resp, nil
}
