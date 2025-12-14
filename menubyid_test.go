package iikocloud

import (
	"encoding/json"
	"testing"
)

func TestMenuByIDUnmarshal(t *testing.T) {
	var menuResponse MenuByIDResponse
	err := json.Unmarshal([]byte(response), &menuResponse)
	if err != nil {
		t.Fatal(err)
	}
}

const response = `
{
    "taxCategories": [],
    "productCategories": [],
    "allergenGroups": [],
    "customerTagGroups": [],
    "overrideTaxCategories": [],
    "revision": 1765531984,
    "formatVersion": 3,
    "id": 66534,
    "name": "CRMProsto",
    "description": "",
    "buttonImageUrl": null,
    "intervals": [],
    "comboCategories": [],
    "itemGroups": [
        {
            "id": "0b33ee83-81fe-4471-8ee1-356fa64756f1",
            "name": "Десерты",
            "description": "",
            "buttonImageUrl": null,
            "iikoGroupId": "cebf6322-1cc7-4401-957a-f92da7e6cc94",
            "items": [
                {
                    "sku": "00656",
                    "name": "Подарок Сан Себастьян",
                    "description": "",
                    "itemSizes": [
                        {
                            "sku": "00656",
                            "sizeCode": null,
                            "sizeName": "",
                            "isDefault": true,
                            "itemModifierGroups": [],
                            "prices": [
                                {
                                    "organizations": [
                                        "d7794d38-cbb9-4ea3-ae55-93708cbc2234"
                                    ],
                                    "price": 0.0,
                                    "taxCategoryId": null
                                }
                            ],
                            "nutritions": [
                                {
                                    "fats": 0.0,
                                    "proteins": 0.0,
                                    "carbs": 0.0,
                                    "energy": 0.0,
                                    "organizations": [
                                        "d7794d38-cbb9-4ea3-ae55-93708cbc2234"
                                    ],
                                    "saturatedFattyAcid": null,
                                    "salt": null,
                                    "sugar": null
                                }
                            ],
                            "isHidden": false,
                            "measureUnitType": "GRAM",
                            "buttonImageUrl": null,
                            "weight": 120.0,
                            "id": null
                        }
                    ],
                    "modifierSchemaId": null,
                    "modifierSchemaName": "",
                    "type": "DISH",
                    "canSetOpenPrice": false,
                    "useBalanceForSell": false,
                    "measureUnit": "",
                    "productCategoryId": "d23e5b41-53ef-1999-0196-c7c4b942f96a",
                    "customerTagGroups": [],
                    "paymentSubject": "COMMODITY",
                    "paymentSubjectCode": "1",
                    "outerEanCode": null,
                    "isMarked": false,
                    "isHidden": false,
                    "barcodes": [],
                    "orderItemType": "Product",
                    "taxCategoryId": null,
                    "allergenGroupIds": [],
                    "labels": [],
                    "tags": [],
                    "id": "d0bf60de-77e1-4dc2-8b27-2179d074e765",
                    "splittable": false
                }
            ],
            "scheduleId": null,
            "scheduleName": null,
            "schedules": [],
            "isHidden": false,
            "labels": [],
            "tags": []
        }
    ]
}`
