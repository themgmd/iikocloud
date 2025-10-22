package iikocloud

import (
	"context"
	"fmt"
	"net/url"

	"github.com/themgmd/iikocloud/internal/constants"
	"github.com/themgmd/iikocloud/internal/http"
)

type NomenclatureRequest struct {
	OrganizationID string `json:"organizationId"`
	StartRevision  int64  `json:"startRevision"`
}

type NomenclatureResponse struct {
	CorrelationID     string            `json:"correlationId"`
	Groups            []Group           `json:"groups"`
	ProductCategories []ProductCategory `json:"productCategories"`
	Products          []Product         `json:"products"`
	Sizes             []Size            `json:"sizes"`
	Revision          int64             `json:"revision"`
}

// Nomenclature .
func (c *Client) Nomenclature(ctx context.Context, req NomenclatureRequest, opts ...http.Option) (*NomenclatureResponse, error) {
	path, err := url.JoinPath(c.baseURL, constants.V1Nomenclature)
	if err != nil {
		return nil, fmt.Errorf("url.JoinPath: %w", err)
	}

	// позволяет перезаписать таймаут если передать опцию снаружи
	opts = append([]http.Option{http.WithCustomTimeout(c.timeout)}, opts...)
	// запрос требует авторизации
	opts = append(opts, http.WithAuthorization(c.token))

	resp, err := http.DoRequest[NomenclatureRequest, NomenclatureResponse](ctx, c.httpClient, path, req, opts...)
	if err != nil {
		return nil, fmt.Errorf("http.DoRequest: %w", err)
	}

	return resp, nil
}
