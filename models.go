package iikocloud

type Group struct {
	ImageLinks       []string `json:"imageLinks"`
	ParentGroup      string   `json:"parentGroup"`
	Order            int      `json:"order"`
	IsIncludedInMenu bool     `json:"isIncludedInMenu"`
	IsGroupModifier  bool     `json:"isGroupModifier"`
	ID               string   `json:"id"`
	Code             string   `json:"code"`
	Name             string   `json:"name"`
	Description      string   `json:"description"`
	AdditionalInfo   string   `json:"additionalInfo"`
	Tags             []string `json:"tags"`
	IsDeleted        bool     `json:"isDeleted"`
	SeoDescription   string   `json:"seoDescription"`
	SeoText          string   `json:"seoText"`
	SeoKeywords      string   `json:"seoKeywords"`
	SeoTitle         string   `json:"seoTitle"`
}

type ProductCategory struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	IsDeleted bool   `json:"isDeleted"`
}

type Price struct {
	CurrentPrice       float64 `json:"currentPrice"`
	IsIncludedInMenu   bool    `json:"isIncludedInMenu"`
	NextPrice          float64 `json:"nextPrice"`
	NextIncludedInMenu bool    `json:"nextIncludedInMenu"`
	NextDatePrice      string  `json:"nextDatePrice"`
}

type SizePrices struct {
	SizeID string `json:"sizeId"`
	Price  Price  `json:"price"`
}

type Modifiers struct {
	ID                  string  `json:"id"`
	DefaultAmount       float64 `json:"defaultAmount"`
	MinAmount           float64 `json:"minAmount"`
	MaxAmount           float64 `json:"maxAmount"`
	Required            bool    `json:"required"`
	HideIfDefaultAmount bool    `json:"hideIfDefaultAmount"`
	Splittable          bool    `json:"splittable"`
	FreeOfChargeAmount  float64 `json:"freeOfChargeAmount"`
}

type ChildModifiers struct {
	ID                  string  `json:"id"`
	DefaultAmount       float64 `json:"defaultAmount"`
	MinAmount           float64 `json:"minAmount"`
	MaxAmount           float64 `json:"maxAmount"`
	Required            bool    `json:"required"`
	HideIfDefaultAmount bool    `json:"hideIfDefaultAmount"`
	Splittable          bool    `json:"splittable"`
	FreeOfChargeAmount  float64 `json:"freeOfChargeAmount"`
}

type GroupModifiers struct {
	ID                                   string           `json:"id"`
	MinAmount                            float64          `json:"minAmount"`
	MaxAmount                            float64          `json:"maxAmount"`
	Required                             bool             `json:"required"`
	ChildModifiersHaveMinMaxRestrictions bool             `json:"childModifiersHaveMinMaxRestrictions"`
	ChildModifiers                       []ChildModifiers `json:"childModifiers"`
	HideIfDefaultAmount                  bool             `json:"hideIfDefaultAmount"`
	DefaultAmount                        float64          `json:"defaultAmount"`
	Splittable                           bool             `json:"splittable"`
	FreeOfChargeAmount                   float64          `json:"freeOfChargeAmount"`
}

type Product struct {
	FatAmount               float64          `json:"fatAmount"`
	ProteinsAmount          float64          `json:"proteinsAmount"`
	CarbohydratesAmount     float64          `json:"carbohydratesAmount"`
	EnergyAmount            float64          `json:"energyAmount"`
	FatFullAmount           float64          `json:"fatFullAmount"`
	ProteinsFullAmount      float64          `json:"proteinsFullAmount"`
	CarbohydratesFullAmount float64          `json:"carbohydratesFullAmount"`
	EnergyFullAmount        float64          `json:"energyFullAmount"`
	Weight                  float64          `json:"weight"`
	GroupID                 string           `json:"groupId"`
	ProductCategoryID       string           `json:"productCategoryId"`
	Type                    string           `json:"type"`
	OrderItemType           string           `json:"orderItemType"`
	ModifierSchemaID        string           `json:"modifierSchemaId"`
	ModifierSchemaName      string           `json:"modifierSchemaName"`
	Splittable              bool             `json:"splittable"`
	MeasureUnit             string           `json:"measureUnit"`
	SizePrices              []SizePrices     `json:"sizePrices"`
	Modifiers               []Modifiers      `json:"modifiers"`
	GroupModifiers          []GroupModifiers `json:"groupModifiers"`
	ImageLinks              []string         `json:"imageLinks"`
	DoNotPrintInCheque      bool             `json:"doNotPrintInCheque"`
	ParentGroup             string           `json:"parentGroup"`
	Order                   int              `json:"order"`
	FullNameEnglish         string           `json:"fullNameEnglish"`
	UseBalanceForSell       bool             `json:"useBalanceForSell"`
	CanSetOpenPrice         bool             `json:"canSetOpenPrice"`
	ID                      string           `json:"id"`
	Code                    string           `json:"code"`
	Name                    string           `json:"name"`
	Description             string           `json:"description"`
	AdditionalInfo          string           `json:"additionalInfo"`
	Tags                    []string         `json:"tags"`
	IsDeleted               bool             `json:"isDeleted"`
	SeoDescription          string           `json:"seoDescription"`
	SeoText                 string           `json:"seoText"`
	SeoKeywords             string           `json:"seoKeywords"`
	SeoTitle                string           `json:"seoTitle"`
}

type Size struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Priority  int    `json:"priority"`
	IsDefault bool   `json:"isDefault"`
}

// ExternalMenu .
type ExternalMenu struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// PriceCategory .
type PriceCategory struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

type Organization struct {
	Country                           string   `json:"country"`
	RestaurantAddress                 string   `json:"restaurantAddress"`
	Latitude                          float64  `json:"latitude"`
	Longitude                         float64  `json:"longitude"`
	UseUAEAddressingSystem            bool     `json:"useUaeAddressingSystem"`
	Version                           string   `json:"version"`
	CurrencyIsoName                   string   `json:"currencyIsoName"`
	CurrencyMinimumDenomination       float64  `json:"currencyMinimumDenomination"`
	CountryPhoneCode                  string   `json:"countryPhoneCode"`
	MarketingSourceRequiredInDelivery bool     `json:"marketingSourceRequiredInDelivery"`
	DefaultDeliveryCityID             string   `json:"defaultDeliveryCityId"`
	DeliveryCityIDs                   []string `json:"deliveryCityIds"`
	DeliveryServiceType               string   `json:"deliveryServiceType"`
	DefaultCallCenterPaymentTypeID    string   `json:"defaultCallCenterPaymentTypeId"`
	OrderItemCommentEnabled           bool     `json:"orderItemCommentEnabled"`
	INN                               string   `json:"inn"`
	ResponseType                      string   `json:"responseType"`
	ID                                string   `json:"id"`
	Name                              string   `json:"name"`
}
