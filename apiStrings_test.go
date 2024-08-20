package azrp

import (
	"testing"
)

func Test_getPssdString(t *testing.T) {
	type args struct {
		pdisk string
		loc   string
		cur   string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Good", args{"P10", "uksouth", "GBP"},
			"https://prices.azure.com/api/retail/prices?api-version=2023-01-01-preview&CurrencyCode=GBP&$filter=armRegionName%20eq%20%27uksouth%27%20and%20serviceFamily%20eq%20%27Storage%27%20and%20skuName%20eq%20%27P10%20LRS%27%20and%20productName%20eq%20%27Premium%20SSD%20Managed%20Disks%27%20and%20meterName%20eq%20%27P10%20LRS%20Disk%27%20and%20priceType%20eq%20%27Consumption%27"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getPssdString(tt.args.pdisk, tt.args.loc, tt.args.cur); got != tt.want {
				t.Errorf("getPssdString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getVmString(t *testing.T) {
	type args struct {
		vm  string
		loc string
		cur string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Good", args{"Standard_D2as_v5", "uksouth", "GBP"},
			"https://prices.azure.com/api/retail/prices?api-version=2023-01-01-preview&CurrencyCode=GBP&$filter=armRegionName%20eq%20%27uksouth%27%20and%20serviceFamily%20eq%20%27Compute%27%20and%20serviceName%20eq%20%27Virtual%20Machines%27%20and%20armSkuName%20eq%20%27Standard_D2as_v5%27"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getVmString(tt.args.vm, tt.args.loc, tt.args.cur); got != tt.want {
				t.Errorf("getVmString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getSssdStrings(t *testing.T) {
	expDisk := "https://prices.azure.com/api/retail/prices?api-version=2023-01-01-preview&CurrencyCode=GBP&$filter=armRegionName%20eq%20%27uksouth%27%20and%20serviceFamily%20eq%20%27Storage%27%20and%20productName%20eq%20%27Standard%20SSD%20Managed%20Disks%27%20and%20skuName%20eq%20%27E10%20LRS%27%20and%20meterName%20eq%20%27E10%20LRS%20Disk%27%20and%20priceType%20eq%20%27Consumption%27"
	expOps := "https://prices.azure.com/api/retail/prices?api-version=2023-01-01-preview&CurrencyCode=GBP&$filter=armRegionName%20eq%20%27uksouth%27%20and%20serviceFamily%20eq%20%27Storage%27%20and%20productName%20eq%20%27Standard%20SSD%20Managed%20Disks%27%20and%20skuName%20eq%20%27E10%20LRS%27%20and%20meterName%20eq%20%27E10%20LRS%20Disk%20Operations%27%20and%20priceType%20eq%20%27Consumption%27"
	type args struct {
		sssdName string
		region   string
		currency string
	}
	tests := []struct {
		name     string
		args     args
		wantDisk string
		wantOps  string
	}{
		{"Good", args{"E10", "uksouth", "GBP"}, expDisk, expOps},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotDisk, gotOps := getSssdStrings(tt.args.sssdName, tt.args.region, tt.args.currency)
			if gotDisk != tt.wantDisk {
				t.Errorf("getSssdStrings() gotDisk = %v, want %v", gotDisk, tt.wantDisk)
			}
			if gotOps != tt.wantOps {
				t.Errorf("getSssdStrings() gotOps = %v, want %v", gotOps, tt.wantOps)
			}
		})
	}
}
