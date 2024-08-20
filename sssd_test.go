package azrp

import (
	"fmt"
	"reflect"
	"strings"
	"testing"
)

func goodSssdGetter(url string) (ApiResponse, error) {
	/*Check which request we are getting*/
	ar := ApiResponse{}
	if !strings.Contains(url, "Operations") {
		//Only populate required fields
		ar.BillingCurrency = "GBP"
		ar.CustomerEntityID = "Default"
		ar.CustomerEntityType = "Retail"
		ar.Count = 1
		ar.NextPageLink = ""
		it := []Item{
			{
				RetailPrice: 8.220138,
			},
		}
		ar.Items = it
	} else {
		ar.BillingCurrency = "GBP"
		ar.CustomerEntityID = "Default"
		ar.CustomerEntityType = "Retail"
		ar.Count = 1
		ar.NextPageLink = ""
		it := []Item{
			{
				RetailPrice: 0.001557,
			},
		}
		ar.Items = it
	}
	return ar, nil
}

func sssdDiskErrorGetter(url string) (ApiResponse, error) {
	/*just error*/
	return ApiResponse{}, fmt.Errorf("it didn't work out")
}

func sssdDiskNoItems(url string) (ApiResponse, error) {
	ar := ApiResponse{}
	//Only populate required fields
	ar.BillingCurrency = "GBP"
	ar.CustomerEntityID = "Default"
	ar.CustomerEntityType = "Retail"
	ar.Count = 1
	ar.NextPageLink = ""
	it := []Item{}
	ar.Items = it
	return ar, nil
}

func sssdDiskTwoItems(url string) (ApiResponse, error) {
	/*Check which request we are getting*/
	ar := ApiResponse{}

	//Only populate required fields
	ar.BillingCurrency = "GBP"
	ar.CustomerEntityID = "Default"
	ar.CustomerEntityType = "Retail"
	ar.Count = 1
	ar.NextPageLink = ""
	it := []Item{
		{
			RetailPrice: 8.220138,
		},
		{
			RetailPrice: 10.0000,
		},
	}
	ar.Items = it
	return ar, nil
}

func sssdOpsError(url string) (ApiResponse, error) {
	/*Check which request we are getting*/
	ar := ApiResponse{}
	if !strings.Contains(url, "Operations") {
		//Only populate required fields
		ar.BillingCurrency = "GBP"
		ar.CustomerEntityID = "Default"
		ar.CustomerEntityType = "Retail"
		ar.Count = 1
		ar.NextPageLink = ""
		it := []Item{
			{
				RetailPrice: 8.220138,
			},
		}
		ar.Items = it
	} else {
		return ApiResponse{}, fmt.Errorf("It didn't work out")
	}
	return ar, nil
}

func sssdOpsTwoItems(url string) (ApiResponse, error) {
	/*Check which request we are getting*/
	ar := ApiResponse{}
	if !strings.Contains(url, "Operations") {
		//Only populate required fields
		ar.BillingCurrency = "GBP"
		ar.CustomerEntityID = "Default"
		ar.CustomerEntityType = "Retail"
		ar.Count = 1
		ar.NextPageLink = ""
		it := []Item{
			{
				RetailPrice: 8.220138,
			},
		}
		ar.Items = it
	} else {
		//Only populate required fields
		ar.BillingCurrency = "GBP"
		ar.CustomerEntityID = "Default"
		ar.CustomerEntityType = "Retail"
		ar.Count = 1
		ar.NextPageLink = ""
		it := []Item{
			{
				RetailPrice: 10.1,
			},
			{
				RetailPrice: 0.2,
			},
		}
		ar.Items = it

	}
	return ar, nil
}

func sssdOpsNoItems(url string) (ApiResponse, error) {
	/*Check which request we are getting*/
	ar := ApiResponse{}
	if !strings.Contains(url, "Operations") {
		//Only populate required fields
		ar.BillingCurrency = "GBP"
		ar.CustomerEntityID = "Default"
		ar.CustomerEntityType = "Retail"
		ar.Count = 1
		ar.NextPageLink = ""
		it := []Item{
			{
				RetailPrice: 8.220138,
			},
		}
		ar.Items = it
	} else {
		//Only populate required fields
		ar.BillingCurrency = "GBP"
		ar.CustomerEntityID = "Default"
		ar.CustomerEntityType = "Retail"
		ar.Count = 1
		ar.NextPageLink = ""
		it := []Item{}
		ar.Items = it
		return ar, nil
	}
	return ar, nil
}

func TestPricer_GetSssdPrice(t *testing.T) {
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
		want    SssdPrice
		wantErr bool
	}{
		{"Good", fields{goodSssdGetter}, args{"E10", "uksouth", "GBP"}, SssdPrice{"E10", 128, "uksouth", "GBP", 8.220138, 0.001557}, false},
		{"BadDiskName", fields{apiGet}, args{"E99", "uksouth", "GBP"}, SssdPrice{}, true},
		{"BadRegionName", fields{apiGet}, args{"E4", "alaska", "GBP"}, SssdPrice{}, true},
		{"ApiDiskError", fields{sssdDiskErrorGetter}, args{"E20", "mexicocentral", "USD"}, SssdPrice{}, true},
		{"NoDiskNoItems", fields{sssdDiskNoItems}, args{"E10", "uksouth", "GBP"}, SssdPrice{}, true},
		{"TooDiskManyItems", fields{sssdDiskTwoItems}, args{"E10", "uksouth", "GBP"}, SssdPrice{}, true},
		{"ApiOpsError", fields{sssdOpsError}, args{"E10", "mexicocentral", "NOR"}, SssdPrice{}, true},
		{"ApiOpsNoItems", fields{sssdOpsNoItems}, args{"E10", "uksouth", "GBP"}, SssdPrice{}, true},
		{"TooManyOpsItems", fields{sssdOpsTwoItems}, args{"E10", "uksouth", "GBP"}, SssdPrice{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Pricer{
				apg: tt.fields.apg,
			}
			got, err := p.GetSssdPrice(tt.args.name, tt.args.region, tt.args.currency)
			if (err != nil) != tt.wantErr {
				t.Errorf("Pricer.GetSssdPrice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pricer.GetSssdPrice() = %v, want %v", got, tt.want)
			}
		})
	}
}
