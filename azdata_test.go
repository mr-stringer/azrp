package azrp

import (
	"testing"
)

func Test_validateCurrencyCode(t *testing.T) {
	type args struct {
		cur string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Good", args{"GBP"}, true},
		{"Bad", args{"ZOP"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateCurrencyCode(tt.args.cur); got != tt.want {
				t.Errorf("validateCurrencyCode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_validateLocation(t *testing.T) {
	type args struct {
		loc string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"Good", args{"uksouth"}, true},
		{"Bad", args{"moon"}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validateLocation(tt.args.loc); got != tt.want {
				t.Errorf("validateLocation() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_GetPssdFromSize(t *testing.T) {
	type args struct {
		sz uint
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"P1", args{3}, "P1"},
		{"P2", args{8}, "P2"},
		{"P3", args{9}, "P3"},
		{"P4", args{30}, "P4"},
		{"P6", args{64}, "P6"},
		{"P10", args{65}, "P10"},
		{"P15", args{256}, "P15"},
		{"P20", args{500}, "P20"},
		{"P30", args{1024}, "P30"},
		{"P40", args{1543}, "P40"},
		{"P50", args{4096}, "P50"},
		{"P60", args{8192}, "P60"},
		{"P70", args{16384}, "P70"},
		{"P80", args{32768}, "P80"},
		{"Error", args{32999}, "error"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetPssdFromSize(tt.args.sz); got != tt.want {
				t.Errorf("getPssdFromSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getSizeFromPssd(t *testing.T) {
	type args struct {
		pssd string
	}
	tests := []struct {
		name    string
		args    args
		want    uint
		wantErr bool
	}{
		{"GoodP1", args{"P1"}, 4, false},
		{"GoodP2", args{"P2"}, 8, false},
		{"GoodP3", args{"P3"}, 16, false},
		{"GoodP4", args{"P4"}, 32, false},
		{"GoodP6", args{"P6"}, 64, false},
		{"Good10", args{"P10"}, 128, false},
		{"GoodP15", args{"P15"}, 256, false},
		{"GoodP20", args{"P20"}, 512, false},
		{"GoodP30", args{"P30"}, 1024, false},
		{"GoodP40", args{"P40"}, 2048, false},
		{"GoodP50", args{"P50"}, 4096, false},
		{"GoodP60", args{"P60"}, 8192, false},
		{"GoodP70", args{"P70"}, 16384, false},
		{"GoodP80", args{"P80"}, 32768, false},
		{"BadDisk", args{"P128"}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetSizeFromPssd(tt.args.pssd)
			if (err != nil) != tt.wantErr {
				t.Errorf("getSizeFromPssd() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getSizeFromPssd() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetSssdFromSize(t *testing.T) {
	type args struct {
		sz uint
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"E1", args{3}, "E1"},
		{"E2", args{8}, "E2"},
		{"E3", args{9}, "E3"},
		{"E4", args{30}, "E4"},
		{"E6", args{64}, "E6"},
		{"E10", args{65}, "E10"},
		{"E15", args{256}, "E15"},
		{"E20", args{500}, "E20"},
		{"E30", args{1024}, "E30"},
		{"E40", args{1543}, "E40"},
		{"E50", args{4096}, "E50"},
		{"E60", args{8192}, "E60"},
		{"E70", args{16384}, "E70"},
		{"E80", args{32768}, "E80"},
		{"Error", args{32999}, "error"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetSssdFromSize(tt.args.sz); got != tt.want {
				t.Errorf("GetSssdFromSize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getSizeFromSssd(t *testing.T) {
	type args struct {
		pssd string
	}
	tests := []struct {
		name    string
		args    args
		want    uint
		wantErr bool
	}{
		{"GoodE1", args{"E1"}, 4, false},
		{"GoodE2", args{"E2"}, 8, false},
		{"GoodE3", args{"E3"}, 16, false},
		{"GoodE4", args{"E4"}, 32, false},
		{"GoodE6", args{"E6"}, 64, false},
		{"GoodE0", args{"E10"}, 128, false},
		{"GoodE15", args{"E15"}, 256, false},
		{"GoodE20", args{"E20"}, 512, false},
		{"GoodE30", args{"E30"}, 1024, false},
		{"GoodE40", args{"E40"}, 2048, false},
		{"GoodE50", args{"E50"}, 4096, false},
		{"GoodE60", args{"E60"}, 8192, false},
		{"GoodE70", args{"E70"}, 16384, false},
		{"GoodE80", args{"E80"}, 32768, false},
		{"BadDisk", args{"E128"}, 0, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetSizeFromSssd(tt.args.pssd)
			if (err != nil) != tt.wantErr {
				t.Errorf("getSizeFromSssd() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getSizeFromSssd() = %v, want %v", got, tt.want)
			}
		})
	}
}
