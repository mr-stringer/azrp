package azrp

import (
	"fmt"
	"reflect"
	"testing"
)

func pssdv2GoodGetter(url string) (ApiResponse, error) {
	ar := ApiResponse{}
	ar.BillingCurrency = "GBP"
	ar.CustomerEntityID = "Default"
	ar.CustomerEntityType = "Retail"
	ar.Count = 5
	ar.NextPageLink = ""
	it := []Item{
		{
			"GBP",
			0,
			0.000099,
			0.000099,
			"uksouth",
			"UK South",
			"2023-04-01T00:00:00Z",
			"972b1104-d07e-5d08-9a98-24eb1cabc63b",
			"Premium LRS Provisioned Capacity",
			"DZH318Z09V07",
			"DZH318Z09V07/0003",
			"",
			"Azure Premium SSD v2",
			"Premium LRS",
			"Storage",
			"DZH317F1HKN0",
			"Storage",
			"1 GiB/Hour",
			"Consumption",
			true,
			"",
			"",
			nil,
		},
		{
			"GBP",
			0,
			0,
			0,
			"uksouth",
			"UK South",
			"2023-08-01T00:00:00Z",
			"c261860b-7b49-5309-b8d5-e07b0258135c",
			"Premium LRS Provisioned Throughput (MBps)",
			"DZH318Z09V07",
			"DZH318Z09V07/0003",
			"",
			"Azure Premium SSD v2",
			"Premium LRS",
			"Storage",
			"DZH317F1HKN0",
			"Storage",
			"1/Hour",
			"Consumption",
			true,
			"",
			"",
			nil,
		},
		{
			"GBP",
			125,
			0.000049,
			0.000049,
			"uksouth",
			"UK South",
			"2023-08-01T00:00:00Z",
			"c261860b-7b49-5309-b8d5-e07b0258135c",
			"Premium LRS Provisioned Throughput (MBps)",
			"DZH318Z09V07",
			"DZH318Z09V07/0003",
			"",
			"Azure Premium SSD v2",
			"Premium LRS",
			"Storage",
			"DZH317F1HKN0",
			"Storage",
			"1/Hour",
			"Consumption",
			true,
			"",
			"",
			nil,
		},
		{
			"GBP",
			0,
			0,
			0,
			"uksouth",
			"UK South",
			"2023-04-01T00:00:00Z",
			"f59b359a-2c37-50cf-893b-9c5acbe20e86",
			"Premium LRS Provisioned IOPS",
			"DZH318Z09V07",
			"DZH318Z09V07/0003",
			"",
			"Azure Premium SSD v2",
			"Premium LRS",
			"Storage",
			"DZH317F1HKN0",
			"Storage",
			"1/Hour",
			"Consumption",
			true,
			"",
			"",
			nil,
		},
		{
			"GBP",
			3000,
			0.000006,
			0.000006,
			"uksouth",
			"UK South",
			"2023-04-01T00:00:00Z",
			"f59b359a-2c37-50cf-893b-9c5acbe20e86",
			"Premium LRS Provisioned IOPS",
			"DZH318Z09V07",
			"DZH318Z09V07/0003",
			"",
			"Azure Premium SSD v2",
			"Premium LRS",
			"Storage",
			"DZH317F1HKN0",
			"Storage",
			"1/Hour",
			"Consumption",
			true,
			"",
			"",
			nil,
		},
	}
	ar.Items = it

	return ar, nil
}

func pssdv2ErrorGetter(url string) (ApiResponse, error) {
	/*just error*/
	return ApiResponse{}, fmt.Errorf("it didn't work out")
}

func pssdv2NoResult(url string) (ApiResponse, error) {
	/*just error*/
	ar := ApiResponse{}
	ar.BillingCurrency = "GBP"
	ar.CustomerEntityID = "Default"
	ar.CustomerEntityType = "Retail"
	ar.Count = 5
	ar.NextPageLink = ""
	it := []Item{}
	ar.Items = it
	return ar, nil
}

func pssdv2NoIOPS(url string) (ApiResponse, error) {
	ar := ApiResponse{}
	ar.BillingCurrency = "GBP"
	ar.CustomerEntityID = "Default"
	ar.CustomerEntityType = "Retail"
	ar.Count = 5
	ar.NextPageLink = ""
	it := []Item{
		{
			"GBP",
			0,
			0.000099,
			0.000099,
			"uksouth",
			"UK South",
			"2023-04-01T00:00:00Z",
			"972b1104-d07e-5d08-9a98-24eb1cabc63b",
			"Premium LRS Provisioned Capacity",
			"DZH318Z09V07",
			"DZH318Z09V07/0003",
			"",
			"Azure Premium SSD v2",
			"Premium LRS",
			"Storage",
			"DZH317F1HKN0",
			"Storage",
			"1 GiB/Hour",
			"Consumption",
			true,
			"",
			"",
			nil,
		},
		{
			"GBP",
			125,
			0.000049,
			0.000049,
			"uksouth",
			"UK South",
			"2023-08-01T00:00:00Z",
			"c261860b-7b49-5309-b8d5-e07b0258135c",
			"Premium LRS Provisioned Throughput (MBps)",
			"DZH318Z09V07",
			"DZH318Z09V07/0003",
			"",
			"Azure Premium SSD v2",
			"Premium LRS",
			"Storage",
			"DZH317F1HKN0",
			"Storage",
			"1/Hour",
			"Consumption",
			true,
			"",
			"",
			nil,
		},
	}
	ar.Items = it

	return ar, nil
}

func pssdv2NoMBs(url string) (ApiResponse, error) {
	ar := ApiResponse{}
	ar.BillingCurrency = "GBP"
	ar.CustomerEntityID = "Default"
	ar.CustomerEntityType = "Retail"
	ar.Count = 5
	ar.NextPageLink = ""
	it := []Item{
		{
			"GBP",
			0,
			0.000099,
			0.000099,
			"uksouth",
			"UK South",
			"2023-04-01T00:00:00Z",
			"972b1104-d07e-5d08-9a98-24eb1cabc63b",
			"Premium LRS Provisioned Capacity",
			"DZH318Z09V07",
			"DZH318Z09V07/0003",
			"",
			"Azure Premium SSD v2",
			"Premium LRS",
			"Storage",
			"DZH317F1HKN0",
			"Storage",
			"1 GiB/Hour",
			"Consumption",
			true,
			"",
			"",
			nil,
		},
		{
			"GBP",
			3000,
			0.000006,
			0.000006,
			"uksouth",
			"UK South",
			"2023-04-01T00:00:00Z",
			"f59b359a-2c37-50cf-893b-9c5acbe20e86",
			"Premium LRS Provisioned IOPS",
			"DZH318Z09V07",
			"DZH318Z09V07/0003",
			"",
			"Azure Premium SSD v2",
			"Premium LRS",
			"Storage",
			"DZH317F1HKN0",
			"Storage",
			"1/Hour",
			"Consumption",
			true,
			"",
			"",
			nil,
		},
	}
	ar.Items = it

	return ar, nil
}

func pssdv2NoCapacity(url string) (ApiResponse, error) {
	ar := ApiResponse{}
	ar.BillingCurrency = "GBP"
	ar.CustomerEntityID = "Default"
	ar.CustomerEntityType = "Retail"
	ar.Count = 5
	ar.NextPageLink = ""
	it := []Item{
		{
			"GBP",
			3000,
			0.000006,
			0.000006,
			"uksouth",
			"UK South",
			"2023-04-01T00:00:00Z",
			"f59b359a-2c37-50cf-893b-9c5acbe20e86",
			"Premium LRS Provisioned IOPS",
			"DZH318Z09V07",
			"DZH318Z09V07/0003",
			"",
			"Azure Premium SSD v2",
			"Premium LRS",
			"Storage",
			"DZH317F1HKN0",
			"Storage",
			"1/Hour",
			"Consumption",
			true,
			"",
			"",
			nil,
		},
		{
			"GBP",
			125,
			0.000049,
			0.000049,
			"uksouth",
			"UK South",
			"2023-08-01T00:00:00Z",
			"c261860b-7b49-5309-b8d5-e07b0258135c",
			"Premium LRS Provisioned Throughput (MBps)",
			"DZH318Z09V07",
			"DZH318Z09V07/0003",
			"",
			"Azure Premium SSD v2",
			"Premium LRS",
			"Storage",
			"DZH317F1HKN0",
			"Storage",
			"1/Hour",
			"Consumption",
			true,
			"",
			"",
			nil,
		},
	}
	ar.Items = it

	return ar, nil
}

func TestPricer_GetPssdv2Price(t *testing.T) {
	type fields struct {
		apg apiGetter
	}
	type args struct {
		region   string
		currency string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    Pssdv2Price
		wantErr bool
	}{
		{"Good", fields{pssdv2GoodGetter}, args{"uksouth", "GBP"}, Pssdv2Price{"uksouth", "GBP", 0.000099, 0.000006, 0.000049}, false},
		{"ApiError", fields{pssdv2ErrorGetter}, args{"uksouth", "GBP"}, Pssdv2Price{}, true},
		{"NoResults", fields{pssdv2NoResult}, args{"uksouth", "GBP"}, Pssdv2Price{}, true},
		{"NoIOPS", fields{pssdv2NoIOPS}, args{"uksouth", "GBP"}, Pssdv2Price{}, true},
		{"NoMBs", fields{pssdv2NoMBs}, args{"uksouth", "GBP"}, Pssdv2Price{}, true},
		{"NoCapacity", fields{pssdv2NoCapacity}, args{"uksouth", "GBP"}, Pssdv2Price{}, true},
		{"BadRegion", fields{apiGet}, args{"uknorth", "GBP"}, Pssdv2Price{}, true},
		{"BadCurrency", fields{apiGet}, args{"uksouth", "DLZ"}, Pssdv2Price{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Pricer{
				apg: tt.fields.apg,
			}
			got, err := p.GetPssdv2Price(tt.args.region, tt.args.currency)
			if (err != nil) != tt.wantErr {
				t.Errorf("Pricer.GetPssdv2Price() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pricer.GetPssdv2Price() = %v, want %v", got, tt.want)
			}
		})
	}
}
