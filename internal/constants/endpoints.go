package constants

type Endpoint = string

const (
	V1AccessTokenEndpoint Endpoint = "/api/1/access_token"
	V1Nomenclature        Endpoint = "/api/1/nomenclature"
	V2Menu                Endpoint = "/api/2/menu"
	V2MenuByID            Endpoint = "/api/2/menu/by_id"
)
