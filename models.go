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

type TerminalGroupItem struct {
	Id             string         `json:"id"`
	OrganizationId string         `json:"organizationId"`
	Name           string         `json:"name"`
	Address        string         `json:"address"`
	TimeZone       string         `json:"timeZone"`
	ExternalData   []ExternalData `json:"externalData"`
}

type PaymentType struct {
	Id                           string              `json:"id,omitempty"`
	Code                         string              `json:"code,omitempty"`
	Name                         string              `json:"name,omitempty"`
	Comment                      string              `json:"comment,omitempty"`
	Combinable                   bool                `json:"combinable,omitempty"`
	ExternalRevision             int                 `json:"externalRevision,omitempty"`
	ApplicableMarketingCampaigns []string            `json:"applicableMarketingCampaigns,omitempty"`
	IsDeleted                    bool                `json:"isDeleted,omitempty"`
	PrintCheque                  bool                `json:"printCheque,omitempty"`
	PaymentProcessingType        string              `json:"paymentProcessingType,omitempty"`
	PaymentTypeKind              string              `json:"paymentTypeKind,omitempty"`
	TerminalGroups               []TerminalGroupItem `json:"terminalGroups,omitempty"`
}

type Coordinates struct {
	Latitude  int `json:"latitude"`
	Longitude int `json:"longitude"`
}

type DeliveryPointAddress struct {
	Type string `json:"type"`
}

type OrderCustomer struct {
	Id              string `json:"id,omitempty"`
	Name            string `json:"name,omitempty"`
	Surname         string `json:"surname,omitempty"`
	Comment         string `json:"comment,omitempty"`
	Gender          string `json:"gender,omitempty"`
	InBlacklist     bool   `json:"inBlacklist,omitempty"`
	BlacklistReason string `json:"blacklistReason,omitempty"`
	Birthdate       string `json:"birthdate,omitempty"`
	Type            string `json:"type,omitempty"`
}

type GuestsInfo struct {
	Count               int  `json:"count"`
	SplitBetweenPersons bool `json:"splitBetweenPersons"`
}

type DeliveryPoint struct {
	Coordinates           Coordinates          `json:"coordinates"`
	Address               DeliveryPointAddress `json:"address"`
	ExternalCartographyId string               `json:"externalCartographyId"`
	Comment               string               `json:"comment"`
}

type LoyaltyProgramResult struct {
	MarketingCampaignId          string           `json:"marketingCampaignId"`
	Name                         string           `json:"name"`
	Discounts                    []DiscountResult `json:"discounts"`
	Upsales                      []Upsale         `json:"upsales"`
	FreeProducts                 []FreeProduct    `json:"freeProducts"`
	AvailableComboSpecifications []string         `json:"availableComboSpecifications"`
	AvailableCombos              []AvailableCombo `json:"availableCombos"`
	NeedToActivateCertificate    bool             `json:"needToActivateCertificate"`
}

type DiscountResult struct {
	Code        int    `json:"code"`
	OrderItemId string `json:"orderItemId"`
	PositionId  string `json:"positionId"`
	DiscountSum int    `json:"discountSum"`
	Amount      int    `json:"amount"`
	Comment     string `json:"comment"`
}

type Upsale struct {
	SourceActionId     string    `json:"sourceActionId"`
	SuggestionText     string    `json:"suggestionText"`
	DescriptionForUser string    `json:"descriptionForUser"`
	ProductCodes       []string  `json:"productCodes"`
	Products           []Product `json:"products"`
}

type FreeProduct struct {
	SourceActionId     string         `json:"sourceActionId"`
	DescriptionForUser string         `json:"descriptionForUser"`
	Products           []OrderProduct `json:"products"`
}

type OrderProduct struct {
	Id    string   `json:"id"`
	Code  string   `json:"code"`
	Size  []string `json:"size"`
	Sizes []Size   `json:"sizes"`
}

type AvailableCombo struct {
	SpecificationId string         `json:"specificationId"`
	GroupMapping    []GroupMapping `json:"groupMapping"`
}

type GroupMapping struct {
	GroupId string `json:"groupId"`
	ItemId  string `json:"itemId"`
}

// AvailablePayment .
type AvailablePayment struct {
	Id          string       `json:"id"`
	MaxSum      int          `json:"maxSum"`
	Order       int          `json:"order"`
	WalletInfos []WalletInfo `json:"walletInfos"`
}

type WalletInfo struct {
	Id           string `json:"id"`
	MaxSum       int    `json:"maxSum"`
	CanHoldMoney bool   `json:"canHoldMoney"`
}

// Warning .
type Warning struct {
	Code      string `json:"Code"`
	ErrorCode string `json:"ErrorCode"`
	Message   string `json:"Message"`
}

type Order struct {
	Id                   string               `json:"id"`
	ExternalNumber       string               `json:"externalNumber"`
	CompleteBefore       string               `json:"completeBefore"`
	Phone                string               `json:"phone"`
	PhoneExtension       string               `json:"phoneExtension"`
	OrderTypeId          string               `json:"orderTypeId"`
	OrderServiceType     string               `json:"orderServiceType"`
	DeliveryPoint        DeliveryPoint        `json:"deliveryPoint"`
	Comment              string               `json:"comment"`
	Customer             OrderCustomer        `json:"customer"`
	Guests               GuestsInfo           `json:"guests"`
	MarketingSourceId    string               `json:"marketingSourceId"`
	OperatorId           string               `json:"operatorId"`
	DeliveryDuration     int                  `json:"deliveryDuration"`
	DeliveryZone         string               `json:"deliveryZone"`
	PriceCategoryId      string               `json:"priceCategoryId"`
	Items                []OrderItem          `json:"items"`
	Combos               []OrderCombo         `json:"combos"`
	Payments             []Payment            `json:"payments"`
	Tips                 []Tip                `json:"tips"`
	SourceKey            string               `json:"sourceKey"`
	DiscountsInfo        DiscountsInfo        `json:"discountsInfo"`
	LoyaltyInfo          LoyaltyInfo          `json:"loyaltyInfo"`
	ChequeAdditionalInfo ChequeAdditionalInfo `json:"chequeAdditionalInfo"`
	ExternalData         []ExternalData       `json:"externalData"`
}

type OrderItem struct {
	Type             string           `json:"type"`
	Amount           int              `json:"amount"`
	ProductSizeId    string           `json:"productSizeId"`
	ComboInformation ComboInformation `json:"comboInformation"`
	Comment          string           `json:"comment"`
}

type OrderCombo struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Amount    int    `json:"amount"`
	Price     int    `json:"price"`
	SourceId  string `json:"sourceId"`
	ProgramId string `json:"programId"`
	SizeId    string `json:"sizeId"`
	Size      Size   `json:"size"`
}

type ComboInformation struct {
	ComboId        string `json:"comboId"`
	ComboSourceId  string `json:"comboSourceId"`
	ComboGroupId   string `json:"comboGroupId"`
	ComboGroupName string `json:"comboGroupName"`
}

type Payment struct {
	PaymentTypeKind        string                `json:"paymentTypeKind"`
	Sum                    int                   `json:"sum"`
	PaymentTypeId          string                `json:"paymentTypeId"`
	IsProcessedExternally  bool                  `json:"isProcessedExternally"`
	PaymentAdditionalData  PaymentAdditionalData `json:"paymentAdditionalData"`
	IsFiscalizedExternally bool                  `json:"isFiscalizedExternally"`
	IsPrepay               bool                  `json:"isPrepay"`
}

type Tip = Payment

type PaymentAdditionalData struct {
	Type string `json:"type"`
}

type DiscountsInfo struct {
	Card                  Card       `json:"card"`
	Discounts             []Discount `json:"discounts"`
	FixedLoyaltyDiscounts bool       `json:"fixedLoyaltyDiscounts"`
}

type Card struct {
	Track string `json:"track"`
}

type Discount struct {
	Type string `json:"type"`
}

type LoyaltyInfo struct {
	Coupon                     string   `json:"coupon"`
	ApplicableManualConditions []string `json:"applicableManualConditions"`
}

type ChequeAdditionalInfo struct {
	NeedReceipt     bool   `json:"needReceipt"`
	Email           string `json:"email"`
	SettlementPlace string `json:"settlementPlace"`
	Phone           string `json:"phone"`
}

type ExternalData struct {
	Key      string `json:"key"`
	Value    string `json:"value"`
	IsPublic bool   `json:"isPublic"`
}

type DynamicDiscount struct {
	ManualConditionId string `json:"manualConditionId"`
	Sum               int    `json:"Sum"`
}

type ErrorInfo struct {
	Code           string      `json:"code"`
	Message        string      `json:"message"`
	Description    string      `json:"description"`
	AdditionalData interface{} `json:"additionalData"`
	ErrorReason    string      `json:"errorReason"`
}

type OrderInfo struct {
	Id             string    `json:"id"`
	PosId          string    `json:"posId"`
	ExternalNumber string    `json:"externalNumber"`
	OrganizationId string    `json:"organizationId"`
	Timestamp      int       `json:"timestamp"`
	CreationStatus string    `json:"creationStatus"`
	ErrorInfo      ErrorInfo `json:"errorInfo"`
	Order          struct {
		TableIds    []string      `json:"tableIds"`
		Customer    OrderCustomer `json:"customer"`
		Phone       string        `json:"phone"`
		Status      string        `json:"status"`
		WhenCreated string        `json:"whenCreated"`
		Waiter      struct {
			Id    string `json:"id"`
			Name  string `json:"name"`
			Phone string `json:"phone"`
		} `json:"waiter"`
		TabName                        string `json:"tabName"`
		SplitOrderBetweenCashRegisters string `json:"splitOrderBetweenCashRegisters"`
		MenuId                         string `json:"menuId"`
		PriceCategory                  struct {
			Id   string `json:"id"`
			Name string `json:"name"`
		} `json:"priceCategory"`
		Sum             int    `json:"sum"`
		Number          int    `json:"number"`
		SourceKey       string `json:"sourceKey"`
		WhenBillPrinted string `json:"whenBillPrinted"`
		WhenClosed      string `json:"whenClosed"`
		Conception      struct {
			Id   string `json:"id"`
			Name string `json:"name"`
			Code string `json:"code"`
		} `json:"conception"`
		GuestsInfo GuestsInfo `json:"guestsInfo"`
		Items      []struct {
			Type    string `json:"type"`
			Status  string `json:"status"`
			Deleted struct {
				DeletionMethod struct {
					Id          string `json:"id"`
					Comment     string `json:"comment"`
					RemovalType struct {
						Id   string `json:"id"`
						Name string `json:"name"`
					} `json:"removalType"`
				} `json:"deletionMethod"`
			} `json:"deleted"`
			Amount      int    `json:"amount"`
			Comment     string `json:"comment"`
			WhenPrinted string `json:"whenPrinted"`
			Size        struct {
				Id   string `json:"id"`
				Name string `json:"name"`
			} `json:"size"`
			ComboInformation struct {
				ComboId       string `json:"comboId"`
				ComboSourceId string `json:"comboSourceId"`
				GroupId       string `json:"groupId"`
				GroupName     string `json:"groupName"`
			} `json:"comboInformation"`
		} `json:"items"`
		Combos   []OrderCombo `json:"combos"`
		Payments []struct {
			PaymentType struct {
				Id   string `json:"id"`
				Name string `json:"name"`
				Kind string `json:"kind"`
			} `json:"paymentType"`
			Sum                    int  `json:"sum"`
			IsPreliminary          bool `json:"isPreliminary"`
			IsExternal             bool `json:"isExternal"`
			IsProcessedExternally  bool `json:"isProcessedExternally"`
			IsFiscalizedExternally bool `json:"isFiscalizedExternally"`
			IsPrepay               bool `json:"isPrepay"`
		} `json:"payments"`
		Tips []struct {
			TipsType struct {
				Id   string `json:"id"`
				Name string `json:"name"`
			} `json:"tipsType"`
			PaymentType struct {
				Id   string `json:"id"`
				Name string `json:"name"`
				Kind string `json:"kind"`
			} `json:"paymentType"`
			Sum                    int  `json:"sum"`
			IsPreliminary          bool `json:"isPreliminary"`
			IsExternal             bool `json:"isExternal"`
			IsProcessedExternally  bool `json:"isProcessedExternally"`
			IsFiscalizedExternally bool `json:"isFiscalizedExternally"`
			IsPrepay               bool `json:"isPrepay"`
		} `json:"tips"`
		Discounts []struct {
			DiscountType struct {
				Id   string `json:"id"`
				Name string `json:"name"`
			} `json:"discountType"`
			Sum                       int      `json:"sum"`
			SelectivePositions        []string `json:"selectivePositions"`
			SelectivePositionsWithSum []struct {
				PositionId string `json:"positionId"`
				Sum        int    `json:"sum"`
			} `json:"selectivePositionsWithSum"`
		} `json:"discounts"`
		OrderType struct {
			Id               string `json:"id"`
			Name             string `json:"name"`
			OrderServiceType string `json:"orderServiceType"`
		} `json:"orderType"`
		TerminalGroupId      string         `json:"terminalGroupId"`
		ProcessedPaymentsSum int            `json:"processedPaymentsSum"`
		LoyaltyInfo          LoyaltyInfo    `json:"loyaltyInfo"`
		ExternalData         []ExternalData `json:"externalData"`
	} `json:"order"`
}

type CreateOrderSettings struct {
	ServicePrint            bool `json:"servicePrint"`
	TransportToFrontTimeout int  `json:"transportToFrontTimeout"`
	CheckStopList           bool `json:"checkStopList"`
}

type Table struct {
	Id              string `json:"id"`
	Number          int    `json:"number"`
	Name            string `json:"name"`
	SeatingCapacity int    `json:"seatingCapacity"`
	Revision        int    `json:"revision"`
	IsDeleted       bool   `json:"isDeleted"`
	PosId           string `json:"posId"`
}

type RestaurantSection struct {
	Id              string  `json:"id"`
	TerminalGroupId string  `json:"terminalGroupId"`
	Name            string  `json:"name"`
	Tables          []Table `json:"tables"`
	Schema          struct {
		Width        int `json:"width"`
		Height       int `json:"height"`
		MarkElements []struct {
			Text string `json:"text"`
			Font struct {
				FontFamily string `json:"fontFamily"`
				Size       int    `json:"size"`
				FontStyle  string `json:"fontStyle"`
			} `json:"font"`
			Color struct {
				A int `json:"a"`
				R int `json:"r"`
				G int `json:"g"`
				B int `json:"b"`
			} `json:"color"`
			X      int `json:"x"`
			Y      int `json:"y"`
			Z      int `json:"z"`
			Angle  int `json:"angle"`
			Width  int `json:"width"`
			Height int `json:"height"`
		} `json:"markElements"`
		TableElements []struct {
			TableId string `json:"tableId"`
			X       int    `json:"x"`
			Y       int    `json:"y"`
			Z       int    `json:"z"`
			Angle   int    `json:"angle"`
			Width   int    `json:"width"`
			Height  int    `json:"height"`
		} `json:"tableElements"`
		RectangleElements []struct {
			Color struct {
				A int `json:"a"`
				R int `json:"r"`
				G int `json:"g"`
				B int `json:"b"`
			} `json:"color"`
			X      int `json:"x"`
			Y      int `json:"y"`
			Z      int `json:"z"`
			Angle  int `json:"angle"`
			Width  int `json:"width"`
			Height int `json:"height"`
		} `json:"rectangleElements"`
		EllipseElements []struct {
			Color struct {
				A int `json:"a"`
				R int `json:"r"`
				G int `json:"g"`
				B int `json:"b"`
			} `json:"color"`
			X      int `json:"x"`
			Y      int `json:"y"`
			Z      int `json:"z"`
			Angle  int `json:"angle"`
			Width  int `json:"width"`
			Height int `json:"height"`
		} `json:"ellipseElements"`
		Revision  int  `json:"revision"`
		IsDeleted bool `json:"isDeleted"`
	} `json:"schema"`
}

type TerminalGroup struct {
	OrganizationId string              `json:"organizationId"`
	Items          []TerminalGroupItem `json:"items"`
}
