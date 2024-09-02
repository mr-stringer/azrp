package azrp

import "testing"

func TestPricer_GetString(t *testing.T) {
	type fields struct {
		apg           apiGetter
		Currency      string
		ApiVersion    string
		ArmRegionName string
		Location      string
		MeterId       string
		MeterName     string
		ProductId     string
		SkuId         string
		ProductName   string
		SkuName       string
		ServiceName   string
		ServiceId     string
		ServiceFamily string
		PriceType     string
		ArmSkuName    string
	}
	tests := []struct {
		name    string
		fields  fields
		want    string
		wantErr bool
	}{
		{"GoodNoFiltersNoApiVersion", fields{apg: apiGet, Currency: "GBP"}, "https://prices.azure.com/api/retail/prices?CurrencyCode=GBP", false},
		{"GoodNoFilters", fields{apg: apiGet, Currency: "USD", ApiVersion: "2023-01-01-preview"}, "https://prices.azure.com/api/retail/prices?api-version=2023-01-01-preview&CurrencyCode=USD", false},
		{"GoodArmRegionName", fields{apg: apiGet, Currency: "GBP", ArmRegionName: "uksouth"}, "https://prices.azure.com/api/retail/prices?CurrencyCode=GBP&$filter=armRegionName%20eq%20%27uksouth%27", false},
		{"GoodLocation", fields{apg: apiGet, Currency: "GBP", Location: "UK South"}, "https://prices.azure.com/api/retail/prices?CurrencyCode=GBP&$filter=location%20eq%20%27UK%20South%27", false},
		{"GoodMeterId", fields{apg: apiGet, Currency: "GBP", MeterId: "001f39e8-7ec4-54cf-88b3-fd78f260394f"}, "https://prices.azure.com/api/retail/prices?CurrencyCode=GBP&$filter=meterId%20eq%20%27001f39e8-7ec4-54cf-88b3-fd78f260394f%27", false},
		{"GoodMeterName", fields{apg: apiGet, Currency: "GBP", MeterName: "D8ls v5"}, "https://prices.azure.com/api/retail/prices?CurrencyCode=GBP&$filter=meterName%20eq%20%27D8ls%20v5%27", false},
		{"GoodProductId", fields{apg: apiGet, Currency: "GBP", ProductId: "DZH318Z0CM95"}, "https://prices.azure.com/api/retail/prices?CurrencyCode=GBP&$filter=productId%20eq%20%27DZH318Z0CM95%27", false},
		{"GoodSkuId", fields{apg: apiGet, Currency: "GBP", SkuId: "DZH318Z0CM2D/01QC"}, "https://prices.azure.com/api/retail/prices?CurrencyCode=GBP&$filter=skuId%20eq%20%27DZH318Z0CM2D/01QC%27", false},
		{"GoodProductName", fields{apg: apiGet, Currency: "GBP", ProductName: "Azure Monitor"}, "https://prices.azure.com/api/retail/prices?CurrencyCode=GBP&$filter=productName%20eq%20%27Azure%20Monitor%27", false},
		{"GoodSkuName", fields{apg: apiGet, Currency: "GBP", SkuName: "vCore"}, "https://prices.azure.com/api/retail/prices?CurrencyCode=GBP&$filter=skuName%20eq%20%27vCore%27", false},
		{"GoodServiceName", fields{apg: apiGet, Currency: "GBP", ServiceName: "SQL Database"}, "https://prices.azure.com/api/retail/prices?CurrencyCode=GBP&$filter=serviceName%20eq%20%27SQL%20Database%27", false},
		{"GoodServiceId", fields{apg: apiGet, Currency: "GBP", ServiceId: "DZH313Z7MMC8"}, "https://prices.azure.com/api/retail/prices?CurrencyCode=GBP&$filter=serviceId%20eq%20%27DZH313Z7MMC8%27", false},
		{"GoodServiceFamily", fields{apg: apiGet, Currency: "GBP", ServiceFamily: "Storage"}, "https://prices.azure.com/api/retail/prices?CurrencyCode=GBP&$filter=serviceFamily%20eq%20%27Storage%27", false},
		{"GoodPriceType", fields{apg: apiGet, Currency: "GBP", PriceType: "Reservation"}, "https://prices.azure.com/api/retail/prices?CurrencyCode=GBP&$filter=priceType%20eq%20%27Reservation%27", false},
		{"GoodArmSkuName", fields{apg: apiGet, Currency: "GBP", ArmSkuName: "Standard_E48ads_v5"}, "https://prices.azure.com/api/retail/prices?CurrencyCode=GBP&$filter=armSkuName%20eq%20%27Standard_E48ads_v5%27", false},
		{"GoodCombo", fields{apg: apiGet, Currency: "GBP", ArmRegionName: "uksouth", ArmSkuName: "Standard_E48ads_v5"}, "https://prices.azure.com/api/retail/prices?CurrencyCode=GBP&$filter=armRegionName%20eq%20%27uksouth%27%20and%20armSkuName%20eq%20%27Standard_E48ads_v5%27", false},
		{"NoCurrency", fields{apg: apiGet, Currency: ""}, "", true},
		{"UnsupportedCurrency", fields{apg: apiGet, Currency: "ZOP"}, "", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Pricer{
				apg:           tt.fields.apg,
				Currency:      tt.fields.Currency,
				ApiVersion:    tt.fields.ApiVersion,
				ArmRegionName: tt.fields.ArmRegionName,
				Location:      tt.fields.Location,
				MeterId:       tt.fields.MeterId,
				MeterName:     tt.fields.MeterName,
				ProductId:     tt.fields.ProductId,
				SkuId:         tt.fields.SkuId,
				ProductName:   tt.fields.ProductName,
				SkuName:       tt.fields.SkuName,
				ServiceName:   tt.fields.ServiceName,
				ServiceId:     tt.fields.ServiceId,
				ServiceFamily: tt.fields.ServiceFamily,
				PriceType:     tt.fields.PriceType,
				ArmSkuName:    tt.fields.ArmSkuName,
			}
			got, err := p.GetString()
			if (err != nil) != tt.wantErr {
				t.Errorf("Pricer.GetString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Pricer.GetString() = %v, want %v", got, tt.want)
			}
		})
	}
}
