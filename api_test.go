package azrp

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

var expectedGoodApiString string = `{"BillingCurrency":"GBP","CustomerEntityId":"Default","CustomerEntityType":"Retail",` +
	`"Items":[{"currencyCode":"GBP","tierMinimumUnits":0,"retailPrice":9.613436,` +
	`"unitPrice":9.613436,"armRegionName":"uksouth","location":"UK South",` +
	`"effectiveStartDate":"2018-05-01T00:00:00Z","meterId":"53c34f93-1bfc-4528-9edc-5579adab17b9",` +
	`"meterName":"P6 LRS Disk","productId":"DZH318Z0BP04","skuId":"DZH318Z0BP04/0014",` +
	`"productName":"Premium SSD Managed Disks","skuName":"P6 LRS","serviceName":"Storage",` +
	`"serviceId":"DZH317F1HKN0","serviceFamily":"Storage","unitOfMeasure":"1/Month",` +
	`"type":"Consumption","isPrimaryMeterRegion":true,"armSkuName":"Premium_SSD_Managed_Disk_P6"}],` +
	`"NextPageLink":null,"Count":1}`

var expectedGoodApiResponse ApiResponse = ApiResponse{
	BillingCurrency:    "GBP",
	CustomerEntityID:   "Default",
	CustomerEntityType: "Retail",
	NextPageLink:       "",
	Count:              1,
	Items: []Item{
		{
			CurrencyCode:         "GBP",
			TierMinimumUnits:     0.0,
			RetailPrice:          9.613436,
			UnitPrice:            9.613436,
			ArmRegionName:        "uksouth",
			Location:             "UK South",
			EffectiveStartDate:   "2018-05-01T00:00:00Z",
			MeterID:              "53c34f93-1bfc-4528-9edc-5579adab17b9",
			MeterName:            "P6 LRS Disk",
			ProductID:            "DZH318Z0BP04",
			SkuID:                "DZH318Z0BP04/0014",
			AvailabilityID:       "",
			ProductName:          "Premium SSD Managed Disks",
			SkuName:              "P6 LRS",
			ServiceName:          "Storage",
			ServiceID:            "DZH317F1HKN0",
			ServiceFamily:        "Storage",
			UnitOfMeasure:        "1/Month",
			Type:                 "Consumption",
			IsPrimaryMeterRegion: true,
			ArmSkuName:           "Premium_SSD_Managed_Disk_P6",
			ReservationTerm:      "",
			SavingsPlan:          nil,
		},
	},
}

func Test_apiGet(t *testing.T) {
	goodSrv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, expectedGoodApiString)
	}))
	defer goodSrv.Close()

	malformedSrv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "This isn't JSON")
	}))
	defer malformedSrv.Close()

	type args struct {
		url string
	}
	tests := []struct {
		name    string
		args    args
		want    ApiResponse
		wantErr bool
	}{
		{"Good", args{goodSrv.URL}, expectedGoodApiResponse, false},
		{"MalformedJson", args{malformedSrv.URL}, ApiResponse{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := apiGet(tt.args.url)
			if (err != nil) != tt.wantErr {
				t.Errorf("apiGet() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("apiGet() = %v, want %v", got, tt.want)
			}
		})
	}
}
