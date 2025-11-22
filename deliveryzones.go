package iikocloud

import (
	"context"
	"fmt"
	"net/url"

	"github.com/themgmd/iikocloud/internal/constants"
	"github.com/themgmd/iikocloud/internal/http"
)

type DeliveryZonesRequest struct {
	OrganizationIds []string `json:"organizationIds"`
}

type DeliveryZone struct {
	Id              string `json:"id"`
	Name            string `json:"name"`
	Coordinates     []struct {
		Latitude  float64 `json:"latitude"`
		Longitude float64 `json:"longitude"`
	} `json:"coordinates"`
	IsDeleted bool `json:"isDeleted"`
}

type DeliveryZonesResponse struct {
	CorrelationId  string          `json:"correlationId"`
	DeliveryZones  []DeliveryZone `json:"deliveryZones"`
}

// DeliveryZones получить зоны доставки
func (c *Client) DeliveryZones(ctx context.Context, req DeliveryZonesRequest, opts ...http.Option) (*DeliveryZonesResponse, error) {
	path, err := url.JoinPath(c.baseURL, constants.V1DeliveryZones)
	if err != nil {
		return nil, fmt.Errorf("url.JoinPath: %w", err)
	}

	opts = append([]http.Option{http.WithCustomTimeout(c.timeout)}, opts...)
	opts = append(opts, http.WithAuthorization(c.token))

	resp, err := http.DoRequest[DeliveryZonesRequest, DeliveryZonesResponse](ctx, c.httpClient, path, req, opts...)
	if err != nil {
		return nil, fmt.Errorf("http.DoRequest: %w", err)
	}

	return resp, nil
}

