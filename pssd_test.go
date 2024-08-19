package azrp

import (
	"fmt"
	"reflect"
	"testing"
)

func pssdGoodGetter(url string) (ApiResponse, error) {
	ar := ApiResponse{}
	ar.BillingCurrency = "GBP"
	ar.CustomerEntityID = "Default"
	ar.CustomerEntityType = "Retail"
	ar.Count = 1
	ar.NextPageLink = ""
	itm := Item{}
	itm.CurrencyCode = "GBP"
	itm.TierMinimumUnits = 0
	itm.RetailPrice = 18.563811
	itm.UnitPrice = 18.563811
	itm.ArmRegionName = "uksouth"
	itm.Location = "UK South"
	itm.EffectiveStartDate = "2018-05-01T00:00:00Z"
	itm.MeterID = "1e0cf0cf-919e-4cfe-a2a0-63e312f80718"
	itm.MeterName = "P10 LRS Disk"
	itm.ProductID = "DZH318Z0BP04"
	itm.SkuID = "DZH318Z0BP04/007N"
	itm.ProductName = "Premium SSD Managed Disks"
	itm.SkuName = "P10 LRS"
	itm.ServiceName = "Storage"
	itm.ServiceID = "DZH317F1HKN0"
	itm.ServiceFamily = "Storage"
	itm.UnitOfMeasure = "1/Month"
	itm.Type = "Consumption"
	itm.IsPrimaryMeterRegion = true
	itm.ArmSkuName = "Premium_SSD_Managed_Disk_P10"
	ar.Items = append(ar.Items, itm)
	return ar, nil
}

func pssdApiError(url string) (ApiResponse, error) {
	return ApiResponse{}, fmt.Errorf("sorry, it didn't work out")
}

func pssdNoItems(url string) (ApiResponse, error) {
	ar := ApiResponse{}
	ar.BillingCurrency = "GBP"
	ar.CustomerEntityID = "Default"
	ar.CustomerEntityType = "Retail"
	ar.Count = 0
	ar.NextPageLink = ""
	return ar, nil
}

func pssdTwoItems(url string) (ApiResponse, error) {
	ar := ApiResponse{}
	ar.BillingCurrency = "GBP"
	ar.CustomerEntityID = "Default"
	ar.CustomerEntityType = "Retail"
	ar.Count = 1
	ar.NextPageLink = ""
	itm := Item{}
	ar.Items = append(ar.Items, itm)
	ar.Items = append(ar.Items, itm)
	return ar, nil
}

//{"BillingCurrency":"GBP","CustomerEntityId":"Default","CustomerEntityType":"Retail","Items"
//,"skuName":"P10 LRS","serviceName":"Storage","serviceId":"DZH317F1HKN0","serviceFamily":"Storage","unitOfMeasure":"1/Month","type":"Consumption","isPrimaryMeterRegion":true,"armSkuName":"Premium_SSD_Managed_Disk_P10"}],"NextPageLink":null,"Count":1}

func TestPricer_GetPssdPrice(t *testing.T) {
	type fields struct {
		apg apiGetter
	}
	type args struct {
		name     string
		region   string
		currency string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    PssdPrice
		wantErr bool
	}{
		{"Good", fields{pssdGoodGetter}, args{"P10", "uksouth", "GBP"}, PssdPrice{"P10", 128, "uksouth", "GBP", 18.563811}, false},
		{"WrongPdiskName", fields{apiGet}, args{"P11", "germanywestcentral", "EUR"}, PssdPrice{}, true},
		{"WrongRegionName", fields{apiGet}, args{"P10", "lapland", "EUR"}, PssdPrice{}, true},
		{"ApiError", fields{pssdApiError}, args{"P10", "uksouth", "GBP"}, PssdPrice{}, true},
		{"NoItemsReturned", fields{pssdNoItems}, args{"P10", "uksouth", "GBP"}, PssdPrice{}, true},
		{"TooManyItemsReturned", fields{pssdTwoItems}, args{"P10", "uksouth", "GBP"}, PssdPrice{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Pricer{
				apg: tt.fields.apg,
			}
			got, err := p.GetPssdPrice(tt.args.name, tt.args.region, tt.args.currency)
			if (err != nil) != tt.wantErr {
				t.Errorf("Pricer.GetPssdPrice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pricer.GetPssdPrice() = %v, want %v", got, tt.want)
			}
		})
	}
}
