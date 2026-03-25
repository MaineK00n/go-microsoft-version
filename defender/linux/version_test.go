package linux_test

import (
	"reflect"
	"testing"

	"github.com/MaineK00n/go-microsoft-version/defender/linux"
)

func TestNewVersion(t *testing.T) {
	type args struct {
		ver string
	}
	tests := []struct {
		name    string
		args    args
		want    linux.Version
		wantErr bool
	}{
		{
			name: "101.24052.2",
			args: args{ver: "101.24052.2"},
			want: linux.Version{Major: 101, Minor: 24052, Patch: 2},
		},
		{
			name: "101.98.84",
			args: args{ver: "101.98.84"},
			want: linux.Version{Major: 101, Minor: 98, Patch: 84},
		},
		{
			name: "leading zeros 101.24052.0002",
			args: args{ver: "101.24052.0002"},
			want: linux.Version{Major: 101, Minor: 24052, Patch: 2},
		},
		{
			name:    "four parts",
			args:    args{ver: "101.24052.0002.0"},
			wantErr: true,
		},
		{
			name:    "non-numeric",
			args:    args{ver: "101.abc.0002"},
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
			got, err := linux.NewVersion(tt.args.ver)
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
		v2 linux.Version
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name:   "101.98.84 = 101.98.84",
			fields: fields{Major: 101, Minor: 98, Patch: 84},
			args:   args{v2: linux.Version{Major: 101, Minor: 98, Patch: 84}},
			want:   0,
		},
		{
			name:   "101.98.84 < 101.24052.2",
			fields: fields{Major: 101, Minor: 98, Patch: 84},
			args:   args{v2: linux.Version{Major: 101, Minor: 24052, Patch: 2}},
			want:   -1,
		},
		{
			name:   "same major/minor, patch diff 101.98.84 < 101.98.85",
			fields: fields{Major: 101, Minor: 98, Patch: 84},
			args:   args{v2: linux.Version{Major: 101, Minor: 98, Patch: 85}},
			want:   -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v1 := linux.Version{
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
			name:   "101.98.84",
			fields: fields{Major: 101, Minor: 98, Patch: 84},
			want:   "101.98.84",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := linux.Version{
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
