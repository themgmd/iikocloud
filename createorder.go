package iikocloud

import (
	"context"
	"fmt"
	"github.com/themgmd/iikocloud/internal/constants"
	"github.com/themgmd/iikocloud/internal/http"
	"net/url"
)

type CreateOrderRequest struct {
	OrganizationId  string `json:"organizationId"`
	TerminalGroupId string `json:"terminalGroupId"`
	Order           struct {
		Id             string   `json:"id"`
		ExternalNumber string   `json:"externalNumber"`
		TableIds       []string `json:"tableIds"`
		Customer       struct {
			Id                                    string `json:"id"`
			Name                                  string `json:"name"`
			Surname                               string `json:"surname"`
			Comment                               string `json:"comment"`
			Birthdate                             string `json:"birthdate"`
			Email                                 string `json:"email"`
			ShouldReceivePromoActionsInfo         bool   `json:"shouldReceivePromoActionsInfo"`
			ShouldReceiveOrderStatusNotifications bool   `json:"shouldReceiveOrderStatusNotifications"`
			Gender                                string `json:"gender"`
			Type                                  string `json:"type"`
		} `json:"customer"`
		Phone                string               `json:"phone"`
		GuestCount           int                  `json:"guestCount"`
		Guests               GuestsInfo           `json:"guests"`
		TabName              string               `json:"tabName"`
		MenuId               interface{}          `json:"menuId"`
		PriceCategoryId      string               `json:"priceCategoryId"`
		Items                []OrderItem          `json:"items"`
		Combos               []OrderCombo         `json:"combos"`
		Payments             []Payment            `json:"payments"`
		Tips                 []Tip                `json:"tips"`
		SourceKey            string               `json:"sourceKey"`
		DiscountsInfo        DiscountsInfo        `json:"discountsInfo"`
		LoyaltyInfo          LoyaltyInfo          `json:"loyaltyInfo"`
		OrderTypeId          string               `json:"orderTypeId"`
		ChequeAdditionalInfo ChequeAdditionalInfo `json:"chequeAdditionalInfo"`
		ExternalData         []ExternalData       `json:"externalData"`
	} `json:"order"`
	CreateOrderSettings CreateOrderSettings `json:"createOrderSettings"`
}

type CreateOrderResponse struct {
	CorrelationId string    `json:"correlationId"`
	OrderInfo     OrderInfo `json:"orderInfo"`
}

// CreateOrder .
func (c *Client) CreateOrder(ctx context.Context, req CreateOrderRequest, opts ...http.Option) (*CreateOrderResponse, error) {
	path, err := url.JoinPath(c.baseURL, constants.V1CreateOrder)
	if err != nil {
		return nil, fmt.Errorf("url.JoinPath: %w", err)
	}

	// позволяет перезаписать таймаут если передать опцию снаружи
	opts = append([]http.Option{http.WithCustomTimeout(c.timeout)}, opts...)
	// запрос требует авторизации
	opts = append(opts, http.WithAuthorization(c.token))

	resp, err := http.DoRequest[CreateOrderRequest, CreateOrderResponse](ctx, c.httpClient, path, req, opts...)
	if err != nil {
		return nil, fmt.Errorf("http.DoRequest: %w", err)
	}

	return resp, nil
}
