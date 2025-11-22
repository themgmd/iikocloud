package iikocloud

import (
	"context"
	"fmt"
	"net/url"

	"github.com/themgmd/iikocloud/internal/constants"
	"github.com/themgmd/iikocloud/internal/http"
)

type CouponsByIDRequest struct {
	OrganizationId string   `json:"organizationId"`
	CouponIds      []string `json:"couponIds"`
}

type CouponsByIDResponse struct {
	CorrelationId string       `json:"correlationId"`
	Coupons       []CouponInfo `json:"coupons"`
}

// CouponsByID получить информацию о купонах по их ID
func (c *Client) CouponsByID(ctx context.Context, req CouponsByIDRequest, opts ...http.Option) (*CouponsByIDResponse, error) {
	path, err := url.JoinPath(c.baseURL, constants.V1CouponsByID)
	if err != nil {
		return nil, fmt.Errorf("url.JoinPath: %w", err)
	}

	opts = append([]http.Option{http.WithCustomTimeout(c.timeout)}, opts...)
	opts = append(opts, http.WithAuthorization(c.token))

	resp, err := http.DoRequest[CouponsByIDRequest, CouponsByIDResponse](ctx, c.httpClient, path, req, opts...)
	if err != nil {
		return nil, fmt.Errorf("http.DoRequest: %w", err)
	}

	return resp, nil
}

