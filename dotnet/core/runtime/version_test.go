package runtime_test

import (
	"reflect"
	"testing"

	"github.com/MaineK00n/go-microsoft-version/dotnet/core/runtime"
)

func TestNewVersion(t *testing.T) {
	type args struct {
		ver string
	}
	tests := []struct {
		name    string
		args    args
		want    runtime.Version
		wantErr bool
	}{
		{
			name: "8.0.2",
			args: args{ver: "8.0.2"},
			want: runtime.Version{Major: 8, Minor: 0, Patch: 2},
		},
		{
			name: "9.0.13",
			args: args{ver: "9.0.13"},
			want: runtime.Version{Major: 9, Minor: 0, Patch: 13},
		},
		{
			name: "trailing space 6.0.36 ",
			args: args{ver: "6.0.36 "},
			want: runtime.Version{Major: 6, Minor: 0, Patch: 36},
		},
		{
			name: "11.0.0-rc.2",
			args: args{ver: "11.0.0-rc.2"},
			want: runtime.Version{Major: 11, Minor: 0, Patch: 0, Prerelease: "rc.2"},
		},
		{
			name: "11.0.0-rc.2.26504.105",
			args: args{ver: "11.0.0-rc.2.26504.105"},
			want: runtime.Version{Major: 11, Minor: 0, Patch: 0, Prerelease: "rc.2.26504.105"},
		},
		{
			name: "SDK 11.0.100-rc.3",
			args: args{ver: "11.0.100-rc.3"},
			want: runtime.Version{Major: 11, Minor: 0, Patch: 100, Prerelease: "rc.3"},
		},
		{
			name: "SDK 11.0.100-rc.3.26522.148",
			args: args{ver: "11.0.100-rc.3.26522.148"},
			want: runtime.Version{Major: 11, Minor: 0, Patch: 100, Prerelease: "rc.3.26522.148"},
		},
		{
			name: "5.0.6-servicing.21220.11",
			args: args{ver: "5.0.6-servicing.21220.11"},
			want: runtime.Version{Major: 5, Minor: 0, Patch: 6, Prerelease: "servicing.21220.11"},
		},
		{
			name: "SDK 9.0.100-preview.7.24407.12",
			args: args{ver: "9.0.100-preview.7.24407.12"},
			want: runtime.Version{Major: 9, Minor: 0, Patch: 100, Prerelease: "preview.7.24407.12"},
		},
		{
			name: "hyphen in prerelease identifier 8.0.0-alpha-1.0",
			args: args{ver: "8.0.0-alpha-1.0"},
			want: runtime.Version{Major: 8, Minor: 0, Patch: 0, Prerelease: "alpha-1.0"},
		},
		{
			name: "build metadata 8.0.2+abc.001",
			args: args{ver: "8.0.2+abc.001"},
			want: runtime.Version{Major: 8, Minor: 0, Patch: 2, Build: "abc.001"},
		},
		{
			name: "prerelease and build metadata 11.0.0-rc.2+sha.5114f85",
			args: args{ver: "11.0.0-rc.2+sha.5114f85"},
			want: runtime.Version{Major: 11, Minor: 0, Patch: 0, Prerelease: "rc.2", Build: "sha.5114f85"},
		},
		{
			name:    "two parts",
			args:    args{ver: "8.0"},
			wantErr: true,
		},
		{
			name:    "four parts",
			args:    args{ver: "8.0.2.1"},
			wantErr: true,
		},
		{
			name:    "four parts with prerelease",
			args:    args{ver: "8.0.2.1-rc.1"},
			wantErr: true,
		},
		{
			name:    "non-numeric",
			args:    args{ver: "8.0.abc"},
			wantErr: true,
		},
		{
			name:    "non-numeric core with prerelease",
			args:    args{ver: "8.x.0-rc.1"},
			wantErr: true,
		},
		{
			name:    "empty string",
			args:    args{ver: ""},
			wantErr: true,
		},
		{
			name:    "empty prerelease",
			args:    args{ver: "11.0.0-"},
			wantErr: true,
		},
		{
			name:    "empty prerelease identifier",
			args:    args{ver: "11.0.0-rc..1"},
			wantErr: true,
		},
		{
			name:    "trailing dot in prerelease",
			args:    args{ver: "11.0.0-rc.1."},
			wantErr: true,
		},
		{
			name:    "invalid character in prerelease",
			args:    args{ver: "11.0.0-rc_1"},
			wantErr: true,
		},
		{
			name:    "non-ASCII character in prerelease",
			args:    args{ver: "11.0.0-rc.é"},
			wantErr: true,
		},
		{
			name:    "leading zero in numeric prerelease identifier",
			args:    args{ver: "11.0.0-rc.02"},
			wantErr: true,
		},
		{
			name:    "empty build metadata",
			args:    args{ver: "8.0.2+"},
			wantErr: true,
		},
		{
			name:    "empty build metadata identifier",
			args:    args{ver: "8.0.2+abc..def"},
			wantErr: true,
		},
		{
			name:    "invalid character in build metadata",
			args:    args{ver: "8.0.2+abc+def"},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := runtime.NewVersion(tt.args.ver)
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
			name: "8.0.2 = 8.0.2",
			args: args{v1: "8.0.2", v2: "8.0.2"},
			want: 0,
		},
		{
			name: "6.0.36 < 8.0.2",
			args: args{v1: "6.0.36", v2: "8.0.2"},
			want: -1,
		},
		{
			name: "same major, minor diff 8.0.2 < 8.1.0",
			args: args{v1: "8.0.2", v2: "8.1.0"},
			want: -1,
		},
		{
			name: "same major/minor, patch diff 8.0.2 < 8.0.13",
			args: args{v1: "8.0.2", v2: "8.0.13"},
			want: -1,
		},
		{
			name: "release > prerelease 11.0.0 > 11.0.0-rc.2",
			args: args{v1: "11.0.0", v2: "11.0.0-rc.2"},
			want: +1,
		},
		{
			name: "prerelease < release 11.0.0-rc.2.26504.105 < 11.0.0",
			args: args{v1: "11.0.0-rc.2.26504.105", v2: "11.0.0"},
			want: -1,
		},
		{
			name: "core takes precedence over prerelease 10.0.5 < 11.0.0-rc.2",
			args: args{v1: "10.0.5", v2: "11.0.0-rc.2"},
			want: -1,
		},
		{
			name: "numeric identifiers compare numerically rc.2 < rc.10",
			args: args{v1: "11.0.0-rc.2", v2: "11.0.0-rc.10"},
			want: -1,
		},
		{
			name: "longer identifier set is greater rc.2 < rc.2.26504.105",
			args: args{v1: "11.0.0-rc.2", v2: "11.0.0-rc.2.26504.105"},
			want: -1,
		},
		{
			name: "rc.2.26504.105 < rc.3",
			args: args{v1: "11.0.0-rc.2.26504.105", v2: "11.0.0-rc.3"},
			want: -1,
		},
		{
			name: "SDK rc.3.26522.148 > rc.3",
			args: args{v1: "11.0.100-rc.3.26522.148", v2: "11.0.100-rc.3"},
			want: +1,
		},
		{
			name: "alphanumeric identifiers compare in ASCII order preview.7 < rc.1",
			args: args{v1: "9.0.100-preview.7.24407.12", v2: "9.0.100-rc.1"},
			want: -1,
		},
		{
			name: "numeric identifier < alphanumeric identifier rc.1 < rc.beta",
			args: args{v1: "11.0.0-rc.1", v2: "11.0.0-rc.beta"},
			want: -1,
		},
		{
			name: "numeric identifiers larger than int64",
			args: args{v1: "11.0.0-rc.99999999999999999999", v2: "11.0.0-rc.100000000000000000000"},
			want: -1,
		},
		{
			name: "same prerelease 5.0.6-servicing.21220.11 = 5.0.6-servicing.21220.11",
			args: args{v1: "5.0.6-servicing.21220.11", v2: "5.0.6-servicing.21220.11"},
			want: 0,
		},
		{
			name: "build metadata is ignored 8.0.2+abc = 8.0.2",
			args: args{v1: "8.0.2+abc", v2: "8.0.2"},
			want: 0,
		},
		{
			name: "build metadata is ignored 11.0.0-rc.2+abc = 11.0.0-rc.2+def",
			args: args{v1: "11.0.0-rc.2+abc", v2: "11.0.0-rc.2+def"},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v1, err := runtime.NewVersion(tt.args.v1)
			if err != nil {
				t.Fatalf("NewVersion(%q) error = %v", tt.args.v1, err)
			}
			v2, err := runtime.NewVersion(tt.args.v2)
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
	type fields struct {
		Major      int
		Minor      int
		Patch      int
		Prerelease string
		Build      string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "8.0.2",
			fields: fields{Major: 8, Minor: 0, Patch: 2},
			want:   "8.0.2",
		},
		{
			name:   "9.0.13",
			fields: fields{Major: 9, Minor: 0, Patch: 13},
			want:   "9.0.13",
		},
		{
			name:   "11.0.0-rc.2.26504.105",
			fields: fields{Major: 11, Minor: 0, Patch: 0, Prerelease: "rc.2.26504.105"},
			want:   "11.0.0-rc.2.26504.105",
		},
		{
			name:   "8.0.2+abc",
			fields: fields{Major: 8, Minor: 0, Patch: 2, Build: "abc"},
			want:   "8.0.2+abc",
		},
		{
			name:   "11.0.0-rc.2+abc",
			fields: fields{Major: 11, Minor: 0, Patch: 0, Prerelease: "rc.2", Build: "abc"},
			want:   "11.0.0-rc.2+abc",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := runtime.Version{
				Major:      tt.fields.Major,
				Minor:      tt.fields.Minor,
				Patch:      tt.fields.Patch,
				Prerelease: tt.fields.Prerelease,
				Build:      tt.fields.Build,
			}
			if got := v.String(); got != tt.want {
				t.Errorf("Version.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVersion_String_RoundTrip(t *testing.T) {
	tests := []string{
		"8.0.2",
		"11.0.0-rc.2",
		"11.0.0-rc.2.26504.105",
		"11.0.100-rc.3",
		"11.0.100-rc.3.26522.148",
		"5.0.6-servicing.21220.11",
		"9.0.100-preview.7.24407.12",
		"11.0.0-rc.2+sha.5114f85",
	}
	for _, ver := range tests {
		t.Run(ver, func(t *testing.T) {
			v, err := runtime.NewVersion(ver)
			if err != nil {
				t.Fatalf("NewVersion() error = %v", err)
			}
			if got := v.String(); got != ver {
				t.Errorf("Version.String() = %v, want %v", got, ver)
			}
		})
	}
}
