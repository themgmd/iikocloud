package iikocloud

import (
	"context"
	"fmt"
	"github.com/themgmd/iikocloud/internal/constants"
	"github.com/themgmd/iikocloud/internal/http"
	"net/url"
)

// CreateOrder .
type CreateOrder struct {
	Id                   string               `json:"id,omitempty"`
	ExternalNumber       string               `json:"externalNumber,omitempty"`
	TableIds             []string             `json:"tableIds,omitempty"`
	Customer             CreateOrderCustomer  `json:"customer,omitempty"`
	Phone                string               `json:"phone,omitempty"`
	PhoneExtension       string               `json:"phoneExtension,omitempty"`
	GuestCount           int                  `json:"guestCount,omitempty"`
	Guests               GuestsInfo           `json:"guests,omitempty"`
	TabName              string               `json:"tabName,omitempty"`
	MenuId               interface{}          `json:"menuId,omitempty"`
	PriceCategoryId      string               `json:"priceCategoryId,omitempty"`
	OrderTypeId          string               `json:"orderTypeId,omitempty"`
	OrderServiceType     string               `json:"orderServiceType,omitempty"`
	DeliveryPoint        DeliveryPoint        `json:"deliveryPoint,omitempty"`
	Comment              string               `json:"comment,omitempty"`
	MarketingSourceId    string               `json:"marketingSourceId,omitempty"`
	OperatorId           string               `json:"operatorId,omitempty"`
	DeliveryDuration     int                  `json:"deliveryDuration,omitempty"`
	DeliveryZone         string               `json:"deliveryZone,omitempty"`
	Items                []OrderItem          `json:"items,omitempty"`
	Combos               []OrderCombo         `json:"combos,omitempty"`
	Payments             []Payment            `json:"payments,omitempty"`
	Tips                 []Tip                `json:"tips,omitempty"`
	SourceKey            string               `json:"sourceKey,omitempty"`
	DiscountsInfo        DiscountsInfo        `json:"discountsInfo,omitempty"`
	LoyaltyInfo          LoyaltyInfo          `json:"loyaltyInfo,omitempty"`
	ChequeAdditionalInfo ChequeAdditionalInfo `json:"chequeAdditionalInfo,omitempty"`
	ExternalData         []ExternalData       `json:"externalData,omitempty"`
	CompleteBefore       string               `json:"completeBefore,omitempty"`
}

// CreateOrderCustomer расширенная структура клиента для создания заказа
type CreateOrderCustomer struct {
	OrderCustomer
	Email                                 string `json:"email,omitempty"`
	ShouldReceivePromoActionsInfo         bool   `json:"shouldReceivePromoActionsInfo,omitempty"`
	ShouldReceiveOrderStatusNotifications bool   `json:"shouldReceiveOrderStatusNotifications,omitempty"`
}

// CreateOrderRequest структура запроса для создания заказа
type CreateOrderRequest struct {
	OrganizationId      string              `json:"organizationId"`
	TerminalGroupId     string              `json:"terminalGroupId"`
	Order               CreateOrder         `json:"order"`
	CreateOrderSettings CreateOrderSettings `json:"createOrderSettings,omitempty"`
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
