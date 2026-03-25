package mac_test

import (
	"reflect"
	"testing"

	"github.com/MaineK00n/go-microsoft-version/defender/mac"
)

func TestNewVersion(t *testing.T) {
	type args struct {
		ver string
	}
	tests := []struct {
		name    string
		args    args
		want    mac.Version
		wantErr bool
	}{
		{
			name: "101.60.91",
			args: args{ver: "101.60.91"},
			want: mac.Version{Major: 101, Minor: 60, Patch: 91},
		},
		{
			name: "101.24082.6",
			args: args{ver: "101.24082.6"},
			want: mac.Version{Major: 101, Minor: 24082, Patch: 6},
		},
		{
			name:    "four parts",
			args:    args{ver: "101.60.91.0"},
			wantErr: true,
		},
		{
			name:    "non-numeric",
			args:    args{ver: "101.abc.91"},
			wantErr: true,
		},
		{
			name:    "empty string",
			args:    args{ver: ""},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := mac.NewVersion(tt.args.ver)
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
	type fields struct {
		Major int
		Minor int
		Patch int
	}
	type args struct {
		v2 mac.Version
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name:   "101.60.91 = 101.60.91",
			fields: fields{Major: 101, Minor: 60, Patch: 91},
			args:   args{v2: mac.Version{Major: 101, Minor: 60, Patch: 91}},
			want:   0,
		},
		{
			name:   "101.60.91 < 101.24082.6",
			fields: fields{Major: 101, Minor: 60, Patch: 91},
			args:   args{v2: mac.Version{Major: 101, Minor: 24082, Patch: 6}},
			want:   -1,
		},
		{
			name:   "same major/minor, patch diff 101.60.91 < 101.60.92",
			fields: fields{Major: 101, Minor: 60, Patch: 91},
			args:   args{v2: mac.Version{Major: 101, Minor: 60, Patch: 92}},
			want:   -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v1 := mac.Version{
				Major: tt.fields.Major,
				Minor: tt.fields.Minor,
				Patch: tt.fields.Patch,
			}
			if got := v1.Compare(tt.args.v2); got != tt.want {
				t.Errorf("Version.Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVersion_String(t *testing.T) {
	type fields struct {
		Major int
		Minor int
		Patch int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "101.60.91",
			fields: fields{Major: 101, Minor: 60, Patch: 91},
			want:   "101.60.91",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := mac.Version{
				Major: tt.fields.Major,
				Minor: tt.fields.Minor,
				Patch: tt.fields.Patch,
			}
			if got := v.String(); got != tt.want {
				t.Errorf("Version.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
