package iikocloud

import (
	"context"
	"fmt"
	"net/url"

	"github.com/themgmd/iikocloud/internal/constants"
	"github.com/themgmd/iikocloud/internal/http"
)

// MenuByIDRequest .
type MenuByIDRequest struct {
	ExternalMenuId  string   `json:"externalMenuId"`
	OrganizationIds []string `json:"organizationIds"`
	Version         int      `json:"version"`
}

// TaxCategory .
type TaxCategory struct {
	Id         string `json:"id"`
	Name       string `json:"name"`
	Percentage int    `json:"percentage"`
}

// AllergenGroup .
type AllergenGroup struct {
	Id        string `json:"id"`
	Code      string `json:"code"`
	Name      string `json:"name"`
	IsDeleted string `json:"isDeleted"`
}

// CustomerTagGroup .
type CustomerTagGroup struct {
	Id                string                 `json:"id"`
	Name              string                 `json:"name"`
	SelectSeveralTags bool                   `json:"selectSeveralTags"`
	Items             []CustomerTagGroupItem `json:"items"`
}

// CustomerTagGroupItem .
type CustomerTagGroupItem struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// OverrideTaxCategory .
type OverrideTaxCategory struct {
	OrderType       string `json:"orderType"`
	BaseTaxCategory string `json:"baseTaxCategory"`
	NewTaxCategory  string `json:"newTaxCategory"`
}

// Interval .
type Interval struct {
	OrganizationId string `json:"organizationId"`
	FromTime       string `json:"fromTime"`
	ToTime         string `json:"toTime"`
}

// ComboSizePrice .
type ComboSizePrice struct {
	Organizations []string `json:"organizations"`
	Price         float64  `json:"price"`
	TaxCategoryId string   `json:"taxCategoryId"`
}

// ComboGroupItemSize .
type ComboGroupItemSize struct {
	ComboSizeId string           `json:"comboSizeId"`
	SizeId      string           `json:"sizeId"`
	Name        string           `json:"name"`
	ShortName   string           `json:"shortName"`
	Prices      []ComboSizePrice `json:"prices"`
}

// ComboGroupItem .
type ComboGroupItem struct {
	ItemId                  string               `json:"itemId"`
	ForbiddenModifiers      []interface{}        `json:"forbiddenModifiers"`
	SizeId                  string               `json:"sizeId"`
	PriceModificationAmount int                  `json:"priceModificationAmount"`
	Sizes                   []ComboGroupItemSize `json:"sizes"`
}

// ComboGroup .
type ComboGroup struct {
	Id          interface{}      `json:"id"`
	Name        string           `json:"name"`
	IsMainGroup bool             `json:"isMainGroup"`
	Items       []ComboGroupItem `json:"items"`
	SkipStep    bool             `json:"skipStep"`
}

// Image .
type Image struct {
	Url  string `json:"url"`
	Hash string `json:"hash"`
}

type ComboSize struct {
	Name        string `json:"name"`
	Id          string `json:"id"`
	ButtonImage Image  `json:"buttonImage"`
	ShortName   string `json:"shortName"`
}

// ComboExtended .
type ComboExtended struct {
	Name           string       `json:"name"`
	Price          string       `json:"price"`
	Groups         []ComboGroup `json:"groups"`
	Image          []Image      `json:"image"`
	Description    string       `json:"description"`
	Sizes          []ComboSize  `json:"sizes"`
	PriceStrategy  string       `json:"priceStrategy"`
	StartDate      string       `json:"startDate"`
	ExpirationDate string       `json:"expirationDate"`
	Id             string       `json:"id"`
}

// ComboCategoryExtended .
type ComboCategoryExtended struct {
	Id     interface{}     `json:"id"`
	Name   string          `json:"name"`
	Combos []ComboExtended `json:"combos"`
}

// ItemGroupCustomerTagGroup .
type ItemGroupCustomerTagGroup struct {
	CustomerTagGroupId string        `json:"customerTagGroupId"`
	SelectedTagIds     []interface{} `json:"selectedTagIds"`
}

// Barcode .
type Barcode struct {
	Barcode   string `json:"barcode"`
	Container string `json:"container"`
}

// ItemGroupNutrition .
type ItemGroupNutrition struct {
	Fats               float64  `json:"fats"`
	Proteins           float64  `json:"proteins"`
	Carbs              float64  `json:"carbs"`
	Energy             float64  `json:"energy"`
	Organizations      []string `json:"organizations"`
	SaturatedFattyAcid float64  `json:"saturatedFattyAcid"`
	Salt               float64  `json:"salt"`
	Sugar              float64  `json:"sugar"`
}

// ItemGroupPrice .
type ItemGroupPrice struct {
	Organizations []string `json:"organizations"`
	Price         float64  `json:"price"`
	TaxCategoryId string   `json:"taxCategoryId"`
}

// Restrictions .
type Restrictions struct {
	MinQuantity           int         `json:"minQuantity"`
	MaxQuantity           int         `json:"maxQuantity"`
	FreeQuantity          int         `json:"freeQuantity"`
	HideIfDefaultQuantity bool        `json:"hideIfDefaultQuantity"`
	DefaultQuantity       interface{} `json:"defaultQuantity"`
}

// ItemModifierGroupItem .
type ItemModifierGroupItem struct {
	Sku                 string                      `json:"sku"`
	Name                string                      `json:"name"`
	Description         string                      `json:"description"`
	Restrictions        Restrictions                `json:"restrictions"`
	IsHidden            bool                        `json:"isHidden"`
	Prices              []ItemGroupPrice            `json:"prices"`
	Nutritions          []ItemGroupNutrition        `json:"nutritions"`
	TaxCategoryId       string                      `json:"taxCategoryId"`
	IndependentQuantity bool                        `json:"independentQuantity"`
	ProductCategoryId   string                      `json:"productCategoryId"`
	CustomerTagGroups   []ItemGroupCustomerTagGroup `json:"customerTagGroups"`
	PaymentSubject      string                      `json:"paymentSubject"`
	OuterEanCode        string                      `json:"outerEanCode"`
	IsMarked            bool                        `json:"isMarked"`
	MeasureUnitType     string                      `json:"measureUnitType"`
	PaymentSubjectCode  string                      `json:"paymentSubjectCode"`
	Barcodes            []Barcode                   `json:"barcodes"`
	ButtonImageUrl      string                      `json:"buttonImageUrl"`
	AllergenGroupIds    []interface{}               `json:"allergenGroupIds"`
	Labels              []interface{}               `json:"labels"`
	Tags                []interface{}               `json:"tags"`
	Id                  string                      `json:"id"`
	Weight              interface{}                 `json:"weight"`
}

// ItemModifierGroup .
type ItemModifierGroup struct {
	Name                                 string                  `json:"name"`
	Description                          string                  `json:"description"`
	Restrictions                         Restrictions            `json:"restrictions"`
	Items                                []ItemModifierGroupItem `json:"items"`
	IsHidden                             bool                    `json:"isHidden"`
	ChildModifiersHaveMinMaxRestrictions bool                    `json:"childModifiersHaveMinMaxRestrictions"`
	Sku                                  string                  `json:"sku"`
	Splittable                           interface{}             `json:"splittable"`
	Id                                   string                  `json:"id"`
}

// ItemGroupItemSize .
type ItemGroupItemSize struct {
	Sku                string               `json:"sku"`
	SizeCode           string               `json:"sizeCode"`
	SizeName           string               `json:"sizeName"`
	IsDefault          bool                 `json:"isDefault"`
	ItemModifierGroups []ItemModifierGroup  `json:"itemModifierGroups"`
	Prices             []ItemGroupPrice     `json:"prices"`
	Nutritions         []ItemGroupNutrition `json:"nutritions"`
	IsHidden           bool                 `json:"isHidden"`
	MeasureUnitType    string               `json:"measureUnitType"`
	ButtonImageUrl     string               `json:"buttonImageUrl"`
	Weight             interface{}          `json:"weight"`
	Id                 string               `json:"id"`
}

// ItemGroupItem .
type ItemGroupItem struct {
	Sku                string                      `json:"sku"`
	Name               string                      `json:"name"`
	Description        string                      `json:"description"`
	ItemSizes          []ItemGroupItemSize         `json:"itemSizes"`
	ModifierSchemaId   string                      `json:"modifierSchemaId"`
	ModifierSchemaName string                      `json:"modifierSchemaName"`
	Type               string                      `json:"type"`
	CanSetOpenPrice    bool                        `json:"canSetOpenPrice"`
	UseBalanceForSell  bool                        `json:"useBalanceForSell"`
	MeasureUnit        string                      `json:"measureUnit"`
	ProductCategoryId  string                      `json:"productCategoryId"`
	CustomerTagGroups  []ItemGroupCustomerTagGroup `json:"customerTagGroups"`
	PaymentSubject     string                      `json:"paymentSubject"`
	PaymentSubjectCode string                      `json:"paymentSubjectCode"`
	OuterEanCode       string                      `json:"outerEanCode"`
	IsMarked           bool                        `json:"isMarked"`
	IsHidden           bool                        `json:"isHidden"`
	Barcodes           []Barcode                   `json:"barcodes"`
	OrderItemType      string                      `json:"orderItemType"`
	TaxCategoryId      string                      `json:"taxCategoryId"`
	AllergenGroupIds   []interface{}               `json:"allergenGroupIds"`
	Labels             []string                    `json:"labels"`
	Tags               []string                    `json:"tags"`
	Id                 string                      `json:"id"`
	Splittable         interface{}                 `json:"splittable"`
}

// Schedule .
type Schedule struct {
	Begin      string   `json:"begin"`
	End        string   `json:"end"`
	DaysOfWeek []string `json:"daysOfWeek"`
}

// ItemGroup .
type ItemGroup struct {
	Id             string          `json:"id"`
	Name           string          `json:"name"`
	Description    string          `json:"description"`
	ButtonImageUrl string          `json:"buttonImageUrl"`
	IikoGroupId    string          `json:"iikoGroupId"`
	Items          []ItemGroupItem `json:"items"`
	ScheduleId     string          `json:"scheduleId"`
	ScheduleName   string          `json:"scheduleName"`
	Schedules      []Schedule      `json:"schedules"`
	IsHidden       bool            `json:"isHidden"`
	Labels         []string        `json:"labels"`
	Tags           []string        `json:"tags"`
}

// MenuByIDResponse .
type MenuByIDResponse struct {
	TaxCategories         []TaxCategory           `json:"taxCategories"`
	ProductCategories     []ProductCategory       `json:"productCategories"`
	AllergenGroups        []AllergenGroup         `json:"allergenGroups"`
	CustomerTagGroups     []CustomerTagGroup      `json:"customerTagGroups"`
	OverrideTaxCategories []OverrideTaxCategory   `json:"overrideTaxCategories"`
	Revision              int                     `json:"revision"`
	FormatVersion         int                     `json:"formatVersion"`
	Id                    int                     `json:"id"`
	Name                  string                  `json:"name"`
	Description           string                  `json:"description"`
	ButtonImageUrl        string                  `json:"buttonImageUrl"`
	Intervals             []Interval              `json:"intervals"`
	ComboCategories       []ComboCategoryExtended `json:"comboCategories"`
	ItemGroups            []ItemGroup             `json:"itemGroups"`
}

// V2MenuByID .
func (c *Client) V2MenuByID(ctx context.Context, req MenuByIDRequest, opts ...http.Option) (*MenuByIDResponse, error) {
	path, err := url.JoinPath(c.baseURL, constants.V2MenuByID)
	if err != nil {
		return nil, fmt.Errorf("url.JoinPath: %w", err)
	}

	// позволяет перезаписать таймаут если передать опцию снаружи
	opts = append([]http.Option{http.WithCustomTimeout(c.timeout)}, opts...)
	// запрос требует авторизации
	opts = append(opts, http.WithAuthorization(c.token))

	resp, err := http.DoRequest[MenuByIDRequest, MenuByIDResponse](ctx, c.httpClient, path, req, opts...)
	if err != nil {
		return nil, fmt.Errorf("http.DoRequest: %w", err)
	}
	return resp, nil
}
