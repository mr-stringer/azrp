package azrp

import "fmt"

// ApiResponse is a struct that encapsulates the response object from the Azure
// Retail Pricing API
type ApiResponse struct {
	BillingCurrency    string `json:"BillingCurrency"`
	CustomerEntityID   string `json:"CustomerEntityId"`
	CustomerEntityType string `json:"CustomerEntityType"`
	NextPageLink       string `json:"NextPageLink"`
	Count              uint   `json:"Count"`
	Items              []Item `json:"Items"`
}

// Item is a struct that encapsulates the Item object found in the Azure Retail
// Pricing API response
type Item struct {
	CurrencyCode         string        `json:"currencyCode"`
	TierMinimumUnits     float32       `json:"tierMinimumUnits"`
	RetailPrice          float32       `json:"retailPrice"`
	UnitPrice            float32       `json:"unitPrice"`
	ArmRegionName        string        `json:"armRegionName"`
	Location             string        `json:"location"`
	EffectiveStartDate   string        `json:"effectiveStartDate"`
	MeterID              string        `json:"meterId`
	MeterName            string        `json:"meterName"`
	ProductID            string        `json:"productId"`
	SkuID                string        `json:"skuId"`
	AvailabilityID       string        `json:"availabilityId"`
	ProductName          string        `json:"productName"`
	SkuName              string        `json:"skuName"`
	ServiceName          string        `json:"serviceName"`
	ServiceID            string        `json:"serviceId"`
	ServiceFamily        string        `json:"serviceFamily"`
	UnitOfMeasure        string        `json:"unitOfMeasure"`
	Type                 string        `json:"type"`
	IsPrimaryMeterRegion bool          `json:"isPrimaryMeterRegion"`
	ArmSkuName           string        `json:"armSkuName"`
	ReservationTerm      string        `json:"reservationTerm"`
	SavingsPlan          []SavingsPlan `json:"savingsPlan"`
}

// SavingsPlan is a struct that encapsulates the SavingsPlan object found in the
// Items object found in the Azure Retail Pricing API
type SavingsPlan struct {
	UnitPrice   float32 `json:"unitPrice"`
	RetailPrice float32 `json:"retailPrice"`
	Term        string  `json:"term"`
}

func (a *ApiResponse) GetCSV(header bool) []string {
	ret := []string{}
	if header {
		str := "\"CurrencyCode\"," +
			"\"TierMinimumUnits\"," +
			"\"RetailPrice\"," +
			"\"UnitPrice\"," +
			"\"ArmRegionName\"," +
			"\"Location\"," +
			"\"EffectiveStartDate\"," +
			"\"MeterID\"," +
			"\"MeterName\"," +
			"\"ProductID\"," +
			"\"SkuID\"," +
			"\"AvailabilityID\"," +
			"\"ProductName\"," +
			"\"SkuName\"," +
			"\"ServiceName\"," +
			"\"ServiceID\"," +
			"\"ServiceFamily\"," +
			"\"UnitOfMeasure\"," +
			"\"Type\"," +
			"\"IsPrimaryMeterRegion\"," +
			"\"ArmSkuName\"," +
			"\"ReservationTerm\""
		ret = append(ret, str)
	}
	for _, v := range a.Items {
		str := fmt.Sprint(
			"\"", v.CurrencyCode, "\",",
			"\"", v.TierMinimumUnits, "\",",
			"\"", v.RetailPrice, "\",",
			"\"", v.UnitPrice, "\",",
			"\"", v.ArmRegionName, "\",",
			"\"", v.Location, "\",",
			"\"", v.EffectiveStartDate, "\",",
			"\"", v.MeterID, "\",",
			"\"", v.MeterName, "\",",
			"\"", v.ProductID, "\",",
			"\"", v.SkuID, "\",",
			"\"", v.AvailabilityID, "\",",
			"\"", v.ProductName, "\",",
			"\"", v.SkuName, "\",",
			"\"", v.ServiceName, "\",",
			"\"", v.ServiceID, "\",",
			"\"", v.ServiceFamily, "\",",
			"\"", v.UnitOfMeasure, "\",",
			"\"", v.Type, "\",",
			"\"", v.IsPrimaryMeterRegion, "\",",
			"\"", v.ArmSkuName, "\",",
			"\"", v.ReservationTerm, "\"")
		ret = append(ret, str)
	}
	return ret
}
