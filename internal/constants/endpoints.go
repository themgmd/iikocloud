package constants

type Endpoint = string

const (
	V1AccessTokenEndpoint        Endpoint = "/api/1/access_token"
	V1Nomenclature               Endpoint = "/api/1/nomenclature"
	V2Menu                       Endpoint = "/api/2/menu"
	V2MenuByID                   Endpoint = "/api/2/menu/by_id"
	V1Organizations              Endpoint = "/api/1/organizations"
	V1PaymentTypes               Endpoint = "/api/1/payment_types"
	V1CalculateCheckin           Endpoint = "/api/1/loyalty/iiko/calculate"
	V1CancelHoldMoney            Endpoint = "/api/1/loyalty/iiko/customer/wallet/cancel_hold"
	V1CreateOrder                Endpoint = "/api/1//order/create"
	V1AvailableRestaurantSection Endpoint = "/api/1/reserve/available_restaurant_sections"
	V1TerminalGroups             Endpoint = "/api/1/terminal_groups"
)
