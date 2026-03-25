package client_test

import (
	"reflect"
	"testing"

	"github.com/MaineK00n/go-microsoft-version/teams/client"
)

func TestNewVersion(t *testing.T) {
	type args struct {
		ver string
	}
	tests := []struct {
		name    string
		args    args
		want    client.Version
		wantErr bool
	}{
		{
			name: "2.10.1",
			args: args{ver: "2.10.1"},
			want: client.Version{Major: 2, Minor: 10, Patch: 1},
		},
		{
			name: "1.3.0",
			args: args{ver: "1.3.0"},
			want: client.Version{Major: 1, Minor: 3, Patch: 0},
		},
		{
			name:    "four parts",
			args:    args{ver: "2.10.1.0"},
			wantErr: true,
		},
		{
			name:    "non-numeric",
			args:    args{ver: "2.abc.1"},
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
			got, err := client.NewVersion(tt.args.ver)
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
		v2 client.Version
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name:   "2.10.1 = 2.10.1",
			fields: fields{Major: 2, Minor: 10, Patch: 1},
			args:   args{v2: client.Version{Major: 2, Minor: 10, Patch: 1}},
			want:   0,
		},
		{
			name:   "1.3.0 < 2.10.1",
			fields: fields{Major: 1, Minor: 3, Patch: 0},
			args:   args{v2: client.Version{Major: 2, Minor: 10, Patch: 1}},
			want:   -1,
		},
		{
			name:   "same major, minor diff 2.3.1 < 2.10.1",
			fields: fields{Major: 2, Minor: 3, Patch: 1},
			args:   args{v2: client.Version{Major: 2, Minor: 10, Patch: 1}},
			want:   -1,
		},
		{
			name:   "same major/minor, patch diff 2.10.0 < 2.10.1",
			fields: fields{Major: 2, Minor: 10, Patch: 0},
			args:   args{v2: client.Version{Major: 2, Minor: 10, Patch: 1}},
			want:   -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v1 := client.Version{
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
			name:   "2.10.1",
			fields: fields{Major: 2, Minor: 10, Patch: 1},
			want:   "2.10.1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := client.Version{
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
