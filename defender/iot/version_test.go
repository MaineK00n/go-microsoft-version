package iot_test

import (
	"reflect"
	"testing"

	"github.com/MaineK00n/go-microsoft-version/defender/iot"
)

func TestNewVersion(t *testing.T) {
	type args struct {
		ver string
	}
	tests := []struct {
		name    string
		args    args
		want    iot.Version
		wantErr bool
	}{
		{
			name: "10.5.2",
			args: args{ver: "10.5.2"},
			want: iot.Version{Major: 10, Minor: 5, Patch: 2},
		},
		{
			name: "22.2.6",
			args: args{ver: "22.2.6"},
			want: iot.Version{Major: 22, Minor: 2, Patch: 6},
		},
		{
			name:    "four parts",
			args:    args{ver: "10.5.2.0"},
			wantErr: true,
		},
		{
			name:    "non-numeric",
			args:    args{ver: "10.abc.2"},
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
			got, err := iot.NewVersion(tt.args.ver)
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
		v2 iot.Version
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name:   "10.5.2 = 10.5.2",
			fields: fields{Major: 10, Minor: 5, Patch: 2},
			args:   args{v2: iot.Version{Major: 10, Minor: 5, Patch: 2}},
			want:   0,
		},
		{
			name:   "10.5.2 < 22.2.6",
			fields: fields{Major: 10, Minor: 5, Patch: 2},
			args:   args{v2: iot.Version{Major: 22, Minor: 2, Patch: 6}},
			want:   -1,
		},
		{
			name:   "same major, minor diff 10.5.2 < 10.6.2",
			fields: fields{Major: 10, Minor: 5, Patch: 2},
			args:   args{v2: iot.Version{Major: 10, Minor: 6, Patch: 2}},
			want:   -1,
		},
		{
			name:   "same major/minor, patch diff 10.5.2 < 10.5.3",
			fields: fields{Major: 10, Minor: 5, Patch: 2},
			args:   args{v2: iot.Version{Major: 10, Minor: 5, Patch: 3}},
			want:   -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v1 := iot.Version{
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
			name:   "10.5.2",
			fields: fields{Major: 10, Minor: 5, Patch: 2},
			want:   "10.5.2",
		},
		{
			name:   "22.2.6",
			fields: fields{Major: 22, Minor: 2, Patch: 6},
			want:   "22.2.6",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := iot.Version{
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
