package iikocloud

import (
	"context"
	"fmt"
	"github.com/themgmd/iikocloud/internal/constants"
	"github.com/themgmd/iikocloud/internal/http"
	"net/url"
)

type AvailableRestaurantSectionsRequest struct {
	TerminalGroupIds []string `json:"terminalGroupIds"`
	ReturnSchema     bool     `json:"returnSchema"`
	Revision         int      `json:"revision"`
}

type AvailableRestaurantSectionsResponse struct {
	CorrelationId      string              `json:"correlationId"`
	RestaurantSections []RestaurantSection `json:"restaurantSections"`
	Revision           int                 `json:"revision"`
}

// AvailableRestaurantSections .
func (c *Client) AvailableRestaurantSections(ctx context.Context, req AvailableRestaurantSectionsRequest, opts ...http.Option) (*AvailableRestaurantSectionsRequest, error) {
	path, err := url.JoinPath(c.baseURL, constants.V1AvailableRestaurantSection)
	if err != nil {
		return nil, fmt.Errorf("url.JoinPath: %w", err)
	}

	// позволяет перезаписать таймаут если передать опцию снаружи
	opts = append([]http.Option{http.WithCustomTimeout(c.timeout)}, opts...)
	// запрос требует авторизации
	opts = append(opts, http.WithAuthorization(c.token))

	resp, err := http.DoRequest[AvailableRestaurantSectionsRequest, AvailableRestaurantSectionsRequest](ctx, c.httpClient, path, req, opts...)
	if err != nil {
		return nil, fmt.Errorf("http.DoRequest: %w", err)
	}
	return resp, nil
}
