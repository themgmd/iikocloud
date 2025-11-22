package iikocloud

import (
	"context"
	"fmt"
	"net/url"

	"github.com/themgmd/iikocloud/internal/constants"
	"github.com/themgmd/iikocloud/internal/http"
)

type StreetsRequest struct {
	OrganizationIds []string `json:"organizationIds"`
	CityId          string   `json:"cityId,omitempty"`
}

type Street struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	CityId    string `json:"cityId"`
	CityName  string `json:"cityName,omitempty"`
	IsDeleted bool   `json:"isDeleted"`
}

type StreetsResponse struct {
	CorrelationId string   `json:"correlationId"`
	Streets       []Street `json:"streets"`
}

// Streets получить список улиц
func (c *Client) Streets(ctx context.Context, req StreetsRequest, opts ...http.Option) (*StreetsResponse, error) {
	path, err := url.JoinPath(c.baseURL, constants.V1Streets)
	if err != nil {
		return nil, fmt.Errorf("url.JoinPath: %w", err)
	}

	opts = append([]http.Option{http.WithCustomTimeout(c.timeout)}, opts...)
	opts = append(opts, http.WithAuthorization(c.token))

	resp, err := http.DoRequest[StreetsRequest, StreetsResponse](ctx, c.httpClient, path, req, opts...)
	if err != nil {
		return nil, fmt.Errorf("http.DoRequest: %w", err)
	}

	return resp, nil
}

