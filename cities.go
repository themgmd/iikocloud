package iikocloud

import (
	"context"
	"fmt"
	"net/url"

	"github.com/themgmd/iikocloud/internal/constants"
	"github.com/themgmd/iikocloud/internal/http"
)

type CitiesRequest struct {
	OrganizationIds []string `json:"organizationIds"`
}

type City struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	RegionId  string `json:"regionId,omitempty"`
	RegionName string `json:"regionName,omitempty"`
	IsDeleted bool   `json:"isDeleted"`
}

type CitiesResponse struct {
	CorrelationId string  `json:"correlationId"`
	Cities        []City  `json:"cities"`
}

// Cities получить список городов
func (c *Client) Cities(ctx context.Context, req CitiesRequest, opts ...http.Option) (*CitiesResponse, error) {
	path, err := url.JoinPath(c.baseURL, constants.V1Cities)
	if err != nil {
		return nil, fmt.Errorf("url.JoinPath: %w", err)
	}

	opts = append([]http.Option{http.WithCustomTimeout(c.timeout)}, opts...)
	opts = append(opts, http.WithAuthorization(c.token))

	resp, err := http.DoRequest[CitiesRequest, CitiesResponse](ctx, c.httpClient, path, req, opts...)
	if err != nil {
		return nil, fmt.Errorf("http.DoRequest: %w", err)
	}

	return resp, nil
}

