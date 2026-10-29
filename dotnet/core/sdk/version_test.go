package sdk_test

import (
	"reflect"
	"testing"

	"github.com/MaineK00n/go-microsoft-version/dotnet/core/sdk"
)

func TestNewVersion(t *testing.T) {
	type args struct {
		ver string
	}
	tests := []struct {
		name    string
		args    args
		want    sdk.Version
		wantErr bool
	}{
		{
			name: "8.0.130",
			args: args{ver: "8.0.130"},
			want: sdk.Version{Major: 8, Minor: 0, Patch: 130},
		},
		{
			name: "10.0.402",
			args: args{ver: "10.0.402"},
			want: sdk.Version{Major: 10, Minor: 0, Patch: 402},
		},
		{
			name: "trailing space 9.0.100 ",
			args: args{ver: "9.0.100 "},
			want: sdk.Version{Major: 9, Minor: 0, Patch: 100},
		},
		{
			name: "11.0.100-rc.3",
			args: args{ver: "11.0.100-rc.3"},
			want: sdk.Version{Major: 11, Minor: 0, Patch: 100, Prerelease: "rc.3"},
		},
		{
			name: "11.0.100-rc.3.26522.148",
			args: args{ver: "11.0.100-rc.3.26522.148"},
			want: sdk.Version{Major: 11, Minor: 0, Patch: 100, Prerelease: "rc.3.26522.148"},
		},
		{
			name: "9.0.100-preview.7.24407.12",
			args: args{ver: "9.0.100-preview.7.24407.12"},
			want: sdk.Version{Major: 9, Minor: 0, Patch: 100, Prerelease: "preview.7.24407.12"},
		},
		{
			name: "prerelease and build metadata 11.0.100-rc.3+sha.5114f85",
			args: args{ver: "11.0.100-rc.3+sha.5114f85"},
			want: sdk.Version{Major: 11, Minor: 0, Patch: 100, Prerelease: "rc.3", Build: "sha.5114f85"},
		},
		{
			name:    "patch < 100 8.0.32",
			args:    args{ver: "8.0.32"},
			wantErr: true,
		},
		{
			name:    "before feature bands 2.1.4",
			args:    args{ver: "2.1.4"},
			wantErr: true,
		},
		{
			name:    "patch 0 11.0.0-rc.2",
			args:    args{ver: "11.0.0-rc.2"},
			wantErr: true,
		},
		{
			name:    "two parts",
			args:    args{ver: "8.0"},
			wantErr: true,
		},
		{
			name:    "four parts",
			args:    args{ver: "8.0.100.1"},
			wantErr: true,
		},
		{
			name:    "non-numeric",
			args:    args{ver: "8.0.abc"},
			wantErr: true,
		},
		{
			name:    "empty string",
			args:    args{ver: ""},
			wantErr: true,
		},
		{
			name:    "empty prerelease",
			args:    args{ver: "11.0.100-"},
			wantErr: true,
		},
		{
			name:    "empty prerelease identifier",
			args:    args{ver: "11.0.100-rc..1"},
			wantErr: true,
		},
		{
			name:    "invalid character in prerelease",
			args:    args{ver: "11.0.100-rc_1"},
			wantErr: true,
		},
		{
			name:    "leading zero in numeric prerelease identifier",
			args:    args{ver: "11.0.100-rc.03"},
			wantErr: true,
		},
		{
			name:    "empty build metadata",
			args:    args{ver: "10.0.402+"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := sdk.NewVersion(tt.args.ver)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewVersion() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewVersion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVersion_FeatureBand(t *testing.T) {
	tests := []struct {
		name          string
		ver           string
		wantBand      int
		wantBandPatch int
	}{
		{
			name:          "8.0.130",
			ver:           "8.0.130",
			wantBand:      1,
			wantBandPatch: 30,
		},
		{
			name:          "10.0.402",
			ver:           "10.0.402",
			wantBand:      4,
			wantBandPatch: 2,
		},
		{
			name:          "11.0.100-rc.3",
			ver:           "11.0.100-rc.3",
			wantBand:      1,
			wantBandPatch: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v, err := sdk.NewVersion(tt.ver)
			if err != nil {
				t.Fatalf("NewVersion() error = %v", err)
			}
			if got := v.FeatureBand(); got != tt.wantBand {
				t.Errorf("Version.FeatureBand() = %v, want %v", got, tt.wantBand)
			}
			if got := v.BandPatch(); got != tt.wantBandPatch {
				t.Errorf("Version.BandPatch() = %v, want %v", got, tt.wantBandPatch)
			}
		})
	}
}

func TestVersion_Compare(t *testing.T) {
	type args struct {
		v1 string
		v2 string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "10.0.402 = 10.0.402",
			args: args{v1: "10.0.402", v2: "10.0.402"},
			want: 0,
		},
		{
			name: "8.0.130 < 9.0.100",
			args: args{v1: "8.0.130", v2: "9.0.100"},
			want: -1,
		},
		{
			name: "same band 10.0.401 < 10.0.402",
			args: args{v1: "10.0.401", v2: "10.0.402"},
			want: -1,
		},
		{
			name: "cross band 10.0.399 < 10.0.400",
			args: args{v1: "10.0.399", v2: "10.0.400"},
			want: -1,
		},
		{
			name: "cross band prerelease 10.0.399 < 10.0.400-rc.1",
			args: args{v1: "10.0.399", v2: "10.0.400-rc.1"},
			want: -1,
		},
		{
			name: "prerelease of a band's first release 10.0.400-rc.1 < 10.0.400",
			args: args{v1: "10.0.400-rc.1", v2: "10.0.400"},
			want: -1,
		},
		{
			name: "cross band 10.0.400 < 10.0.402",
			args: args{v1: "10.0.400", v2: "10.0.402"},
			want: -1,
		},
		{
			name: "cross band 10.0.399 < 10.0.402",
			args: args{v1: "10.0.399", v2: "10.0.402"},
			want: -1,
		},
		{
			name: "release > prerelease 11.0.100 > 11.0.100-rc.3",
			args: args{v1: "11.0.100", v2: "11.0.100-rc.3"},
			want: +1,
		},
		{
			name: "longer identifier set is greater rc.3 < rc.3.26522.148",
			args: args{v1: "11.0.100-rc.3", v2: "11.0.100-rc.3.26522.148"},
			want: -1,
		},
		{
			name: "numeric identifiers compare numerically rc.3 < rc.10",
			args: args{v1: "11.0.100-rc.3", v2: "11.0.100-rc.10"},
			want: -1,
		},
		{
			name: "alphanumeric identifiers compare in ASCII order preview.7 < rc.1",
			args: args{v1: "9.0.100-preview.7.24407.12", v2: "9.0.100-rc.1"},
			want: -1,
		},
		{
			name: "numeric identifier < alphanumeric identifier rc.1 < rc.beta",
			args: args{v1: "11.0.100-rc.1", v2: "11.0.100-rc.beta"},
			want: -1,
		},
		{
			name: "build metadata is ignored 10.0.402+abc = 10.0.402",
			args: args{v1: "10.0.402+abc", v2: "10.0.402"},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v1, err := sdk.NewVersion(tt.args.v1)
			if err != nil {
				t.Fatalf("NewVersion(%q) error = %v", tt.args.v1, err)
			}
			v2, err := sdk.NewVersion(tt.args.v2)
			if err != nil {
				t.Fatalf("NewVersion(%q) error = %v", tt.args.v2, err)
			}
			if got := v1.Compare(v2); got != tt.want {
				t.Errorf("Version.Compare() = %v, want %v", got, tt.want)
			}
			if got := v2.Compare(v1); got != -tt.want {
				t.Errorf("Version.Compare() reversed = %v, want %v", got, -tt.want)
			}
		})
	}
}

func TestVersion_String(t *testing.T) {
	tests := []string{
		"8.0.130",
		"10.0.402",
		"11.0.100-rc.3",
		"11.0.100-rc.3.26522.148",
		"9.0.100-preview.7.24407.12",
		"11.0.100-rc.3+sha.5114f85",
	}
	for _, ver := range tests {
		t.Run(ver, func(t *testing.T) {
			v, err := sdk.NewVersion(ver)
			if err != nil {
				t.Fatalf("NewVersion() error = %v", err)
			}
			if got := v.String(); got != ver {
				t.Errorf("Version.String() = %v, want %v", got, ver)
			}
		})
	}
}
