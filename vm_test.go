package azrp

import (
	"fmt"
	"reflect"
	"testing"
)

func goodGetter(url string) (ApiResponse, error) {
	ar := ApiResponse{}
	ar.BillingCurrency = "GBP"
	ar.CustomerEntityID = "Default"
	ar.CustomerEntityType = "Retail"
	ar.Count = 13
	ar.Items = []Item{
		{
			"GBP",
			0.0,
			0.059783,
			0.059783,
			"uksouth",
			"UK South",
			"2023-10-01T00:00:00Z",
			"24e8084c-e15a-57dc-888a-dd883c1d78a1",
			"D2as v5 Low Priority",
			"DZH318Z08W0R",
			"DZH318Z08W0R/002W",
			"",
			"Dasv5 Series Cloud Services",
			"Standard_D2as_v5 Low Priority",
			"Virtual Machines",
			"DZH313Z7MMC8",
			"Compute",
			"1 Hour",
			"Consumption",
			true,
			"Standard_D2as_v5",
			"",
			nil},
		{
			"GBP",
			0.0,
			0.010119,
			0.010119,
			"uksouth",
			"UK South",
			"2024-08-01T00:00:00Z",
			"51541b9e-5e7c-5866-956c-b92eb075c00d",
			"D2as v5 Spot",
			"DZH318Z08W0T",
			"DZH318Z08W0T/0005",
			"",
			"Virtual Machines Dasv5 Series",
			"Standard_D2as_v5 Spot",
			"Virtual Machines",
			"DZH313Z7MMC8",
			"Compute",
			"1 Hour",
			"Consumption",
			true,
			"Standard_D2as_v5",
			"",
			nil,
		},
		{
			"GBP",
			0.0,
			0.149457,
			0.149457,
			"uksouth",
			"UK South",
			"2023-10-01T00:00:00Z",
			"7e35f19a-40a7-5ae7-9892-13e7b607d63a",
			"D2as v5",
			"DZH318Z08W0R",
			"DZH318Z08W0R/002P",
			"",
			"Dasv5 Series Cloud Services",
			"Standard_D2as_v5",
			"Virtual Machines",
			"DZH313Z7MMC8",
			"Compute",
			"1 Hour",
			"Consumption",
			true,
			"Standard_D2as_v5",
			"",
			nil,
		},
		{
			"GBP",
			0.0,
			0.077842,
			0.077842,
			"uksouth",
			"UK South",
			"2021-11-01T00:00:00Z",
			"9278400e-3afa-5985-a200-c8edc2448631",
			"D2as v5",
			"DZH318Z08W0T",
			"DZH318Z08W0T/005X",
			"",
			"Virtual Machines Dasv5 Series",
			"Standard_D2as_v5",
			"Virtual Machines",
			"DZH313Z7MMC8",
			"Compute",
			"1 Hour",
			"Consumption",
			true,
			"Standard_D2as_v5",
			"",
			[]SavingsPlan{
				{
					0.041085,
					0.041085,
					"3 years",
				},
				{
					0.058475,
					0.058475,
					"1 Year",
				},
			},
		},
		{
			"GBP",
			0.0,
			402.444246,
			402.444246,
			"uksouth",
			"UK South",
			"2021-11-01T00:00:00Z",
			"9278400e-3afa-5985-a200-c8edc2448631",
			"D2as v5",
			"DZH318Z08W0T",
			"DZH318Z08W0T/00Z0",
			"",
			"Virtual Machines Dasv5 Series",
			"Standard_D2as_v5",
			"Virtual Machines",
			"DZH313Z7MMC8",
			"Compute",
			"1 Hour",
			"Reservation",
			true,
			"Standard_D2as_v5",
			"1 Year",
			nil,
		},
		{
			"GBP",
			0.0,
			777.643716,
			777.643716,
			"uksouth",
			"UK South",
			"2021-11-01T00:00:00Z",
			"9278400e-3afa-5985-a200-c8edc2448631",
			"D2as v5",
			"DZH318Z08W0T",
			"DZH318Z08W0T/00XR",
			"",
			"Virtual Machines Dasv5 Series",
			"Standard_D2as_v5",
			"Virtual Machines",
			"DZH313Z7MMC8",
			"Compute",
			"1 Hour",
			"Reservation",
			true,
			"Standard_D2as_v5",
			"3 Years",
			nil,
		},
		{
			"GBP",
			0.0,
			0.059783,
			0.059783,
			"uksouth",
			"UK South",
			"2021-11-01T00:00:00Z",
			"9506f50a-49bf-502a-8072-2a893a6a7287",
			"D2as v5 Low Priority",
			"DZH318Z08W0Q",
			"DZH318Z08W0Q/001Q",
			"",
			"Virtual Machines Dasv5 Series Windows",
			"Standard_D2as_v5 Low Priority",
			"Virtual Machines",
			"DZH313Z7MMC8",
			"Compute",
			"1 Hour",
			"Consumption",
			true,
			"Standard_D2as_v5",
			"",
			nil,
		},
		{
			"GBP",
			0.0,
			0.015568,
			0.015568,
			"uksouth",
			"UK South",
			"2021-11-01T00:00:00Z",
			"9506f50a-49bf-502a-8072-2a893a6a7287",
			"D2as v5 Low Priority",
			"DZH318Z08W0Q",
			"DZH318Z08W0Q/001Q",
			"",
			"Virtual Machines Dasv5 Series Windows",
			"Standard_D2as_v5 Low Priority",
			"Virtual Machines",
			"DZH313Z7MMC8",
			"Compute",
			"1 Hour",
			"DevTestConsumption",
			true,
			"",
			"Standard_D2as_v5",
			nil,
		},
		{
			"GBP",
			0,
			0.015568,
			0.015568,
			"uksouth",
			"UK South",
			"2021-11-01T00:00:00Z",
			"c105bfeb-4904-5ac5-8166-bf14962d7284",
			"D2as v5 Low Priority",
			"DZH318Z08W0T",
			"DZH318Z08W0T/0053",
			"",
			"Virtual Machines Dasv5 Series",
			"Standard_D2as_v5 Low Priority",
			"Virtual Machines",
			"DZH313Z7MMC8",
			"Compute",
			"1 Hour",
			"Consumption",
			true,
			"",
			"Standard_D2as_v5",
			nil,
		},
		{
			"GBP",
			0.0,
			0.149457,
			0.149457,
			"uksouth",
			"UK South",
			"2021-11-01T00:00:00Z",
			"cdb8880b-0267-5c3a-a35b-ccd8c9f53db9",
			"D2as v5",
			"DZH318Z08W0Q",
			"DZH318Z08W0Q/001H",
			"",
			"Virtual Machines Dasv5 Series Windows",
			"Standard_D2as_v5",
			"Virtual Machines",
			"DZH313Z7MMC8",
			"Compute",
			"1 Hour",
			"Consumption",
			true,
			"",
			"Standard_D2as_v5",
			nil,
		},
		{
			"GBP",
			0,
			0.077842,
			0.077842,
			"uksouth",
			"UK South",
			"2021-11-01T00:00:00Z",
			"cdb8880b-0267-5c3a-a35b-ccd8c9f53db9",
			"D2as v5",
			"DZH318Z08W0Q",
			"DZH318Z08W0Q/001H",
			"",
			"Virtual Machines Dasv5 Series Windows",
			"Standard_D2as_v5",
			"Virtual Machines",
			"DZH313Z7MMC8",
			"Compute",
			"1 Hour",
			"DevTestConsumption",
			true,
			"Standard_D2as_v5",
			"",
			nil,
		},
		{
			"GBP",
			0,
			0.019429,
			0.019429,
			"uksouth",
			"UK South",
			"2024-08-01T00:00:00Z",
			"ffc3fffc-b410-52ce-9ea4-bcaf8d158a98",
			"D2as v5 Spot",
			"DZH318Z08W0Q",
			"DZH318Z08W0Q/0027",
			"",
			"Virtual Machines Dasv5 Series Windows",
			"Standard_D2as_v5 Spot",
			"Virtual Machines",
			"DZH313Z7MMC8",
			"Compute",
			"1 Hour",
			"Consumption",
			true,
			"Standard_D2as_v5",
			"",
			nil,
		},
		{
			"GBP",
			0,
			0.010119,
			0.010119,
			"uksouth",
			"UK South",
			"2024-08-01T00:00:00Z",
			"ffc3fffc-b410-52ce-9ea4-bcaf8d158a98",
			"D2as v5 Spot",
			"DZH318Z08W0Q",
			"DZH318Z08W0Q/0027",
			"",
			"Virtual Machines Dasv5 Series Windows",
			"Standard_D2as_v5 Spot",
			"Virtual Machines",
			"DZH313Z7MMC8",
			"Compute",
			"1 Hour",
			"DevTestConsumption",
			true,
			"Standard_D2as_v5",
			"",
			nil,
		},
	}
	return ar, nil
}

func no1yeRiGetter(url string) (ApiResponse, error) {
	ar := ApiResponse{}
	ar.BillingCurrency = "GBP"
	ar.CustomerEntityID = "Default"
	ar.CustomerEntityType = "Retail"
	ar.Count = 13
	ar.Items = []Item{
		{
			"GBP",
			0.0,
			0.077842,
			0.077842,
			"uksouth",
			"UK South",
			"2021-11-01T00:00:00Z",
			"9278400e-3afa-5985-a200-c8edc2448631",
			"D2as v5",
			"DZH318Z08W0T",
			"DZH318Z08W0T/005X",
			"",
			"Virtual Machines Dasv5 Series",
			"Standard_D2as_v5",
			"Virtual Machines",
			"DZH313Z7MMC8",
			"Compute",
			"1 Hour",
			"Consumption",
			true,
			"Standard_D2as_v5",
			"",
			[]SavingsPlan{
				{
					0.041085,
					0.041085,
					"3 years",
				},
				{
					0.058475,
					0.058475,
					"1 Year",
				},
			},
		},
		{
			"GBP",
			0.0,
			777.643716,
			777.643716,
			"uksouth",
			"UK South",
			"2021-11-01T00:00:00Z",
			"9278400e-3afa-5985-a200-c8edc2448631",
			"D2as v5",
			"DZH318Z08W0T",
			"DZH318Z08W0T/00XR",
			"",
			"Virtual Machines Dasv5 Series",
			"Standard_D2as_v5",
			"Virtual Machines",
			"DZH313Z7MMC8",
			"Compute",
			"1 Hour",
			"Reservation",
			true,
			"Standard_D2as_v5",
			"3 Years",
			nil,
		},
	}
	return ar, nil
}

func no3yeRiGetter(url string) (ApiResponse, error) {
	ar := ApiResponse{}
	ar.BillingCurrency = "GBP"
	ar.CustomerEntityID = "Default"
	ar.CustomerEntityType = "Retail"
	ar.Count = 13
	ar.Items = []Item{
		{
			"GBP",
			0.0,
			0.077842,
			0.077842,
			"uksouth",
			"UK South",
			"2021-11-01T00:00:00Z",
			"9278400e-3afa-5985-a200-c8edc2448631",
			"D2as v5",
			"DZH318Z08W0T",
			"DZH318Z08W0T/005X",
			"",
			"Virtual Machines Dasv5 Series",
			"Standard_D2as_v5",
			"Virtual Machines",
			"DZH313Z7MMC8",
			"Compute",
			"1 Hour",
			"Consumption",
			true,
			"Standard_D2as_v5",
			"",
			[]SavingsPlan{
				{
					0.041085,
					0.041085,
					"3 years",
				},
				{
					0.058475,
					0.058475,
					"1 Year",
				},
			},
		},
		{
			"GBP",
			0.0,
			402.444246,
			402.444246,
			"uksouth",
			"UK South",
			"2021-11-01T00:00:00Z",
			"9278400e-3afa-5985-a200-c8edc2448631",
			"D2as v5",
			"DZH318Z08W0T",
			"DZH318Z08W0T/00Z0",
			"",
			"Virtual Machines Dasv5 Series",
			"Standard_D2as_v5",
			"Virtual Machines",
			"DZH313Z7MMC8",
			"Compute",
			"1 Hour",
			"Reservation",
			true,
			"Standard_D2as_v5",
			"1 Year",
			nil,
		},
	}
	return ar, nil
}

func noPaygRiGetter(url string) (ApiResponse, error) {
	ar := ApiResponse{}
	ar.BillingCurrency = "GBP"
	ar.CustomerEntityID = "Default"
	ar.CustomerEntityType = "Retail"
	ar.Count = 13
	ar.Items = []Item{

		{
			"GBP",
			0.0,
			402.444246,
			402.444246,
			"uksouth",
			"UK South",
			"2021-11-01T00:00:00Z",
			"9278400e-3afa-5985-a200-c8edc2448631",
			"D2as v5",
			"DZH318Z08W0T",
			"DZH318Z08W0T/00Z0",
			"",
			"Virtual Machines Dasv5 Series",
			"Standard_D2as_v5",
			"Virtual Machines",
			"DZH313Z7MMC8",
			"Compute",
			"1 Hour",
			"Reservation",
			true,
			"Standard_D2as_v5",
			"1 Year",
			nil,
		},
		{
			"GBP",
			0.0,
			777.643716,
			777.643716,
			"uksouth",
			"UK South",
			"2021-11-01T00:00:00Z",
			"9278400e-3afa-5985-a200-c8edc2448631",
			"D2as v5",
			"DZH318Z08W0T",
			"DZH318Z08W0T/00XR",
			"",
			"Virtual Machines Dasv5 Series",
			"Standard_D2as_v5",
			"Virtual Machines",
			"DZH313Z7MMC8",
			"Compute",
			"1 Hour",
			"Reservation",
			true,
			"Standard_D2as_v5",
			"3 Years",
			nil,
		},
	}
	return ar, nil
}

func badGetter(url string) (ApiResponse, error) {
	return ApiResponse{}, fmt.Errorf("it didn't work out")
}

func emptyGetter(url string) (ApiResponse, error) {
	/*empty object with no error*/
	return ApiResponse{}, nil
}

func TestPricer_GetVmPrice(t *testing.T) {
	type fields struct {
		apg apiGetter
	}
	type args struct {
		vmSku    string
		region   string
		currency string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    VmPrice
		wantErr bool
	}{
		{"Good", fields{goodGetter}, args{"Standard_D2as_v5", "uksouth", "GBP"},
			VmPrice{"Standard_D2as_v5", "uksouth", "GBP", 0.077842, 33.53702, 21.601215, 0, 0}, false},
		{"BadCurrency", fields{apiGet}, args{"Standard_D2ds_v5", "uksouth", "ZOP"}, VmPrice{}, true},
		{"BadLocation", fields{apiGet}, args{"Standard_D2ds_v5", "francewest", "EUR"}, VmPrice{}, true},
		{"FailedToPrice", fields{badGetter}, args{"Standard_D2ds_v5", "uksouth", "GBP"}, VmPrice{}, true},
		{"EmptyResults", fields{emptyGetter}, args{"Standard_D2ds_v5", "uksouth", "GBP"}, VmPrice{}, true},
		{"No1YrRi", fields{no1yeRiGetter}, args{"Standard_D2ds_v5", "uksouth", "GBP"}, VmPrice{}, true},
		{"No3YrRi", fields{no3yeRiGetter}, args{"Standard_D2ds_v5", "uksouth", "GBP"}, VmPrice{}, true},
		{"NoPayg", fields{noPaygRiGetter}, args{"Standard_D2ds_v5", "uksouth", "GBP"}, VmPrice{}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Pricer{
				apg: tt.fields.apg,
			}
			got, err := p.GetVmPrice(tt.args.vmSku, tt.args.region, tt.args.currency)
			if (err != nil) != tt.wantErr {
				t.Errorf("Pricer.GetVmPrice() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Pricer.GetVmPrice() = %v, want %v", got, tt.want)
			}
		})
	}
}
