package core_test

import (
	"reflect"
	"testing"

	"github.com/MaineK00n/go-microsoft-version/dotnet/core"
)

func TestNewVersion(t *testing.T) {
	type args struct {
		ver string
	}
	tests := []struct {
		name    string
		args    args
		want    core.Version
		wantErr bool
	}{
		{
			name: "8.0.2",
			args: args{ver: "8.0.2"},
			want: core.Version{Major: 8, Minor: 0, Patch: 2},
		},
		{
			name: "9.0.13",
			args: args{ver: "9.0.13"},
			want: core.Version{Major: 9, Minor: 0, Patch: 13},
		},
		{
			name: "trailing space 6.0.36 ",
			args: args{ver: "6.0.36 "},
			want: core.Version{Major: 6, Minor: 0, Patch: 36},
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
			name:    "non-numeric",
			args:    args{ver: "8.0.abc"},
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
			got, err := core.NewVersion(tt.args.ver)
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
		v2 core.Version
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name:   "8.0.2 = 8.0.2",
			fields: fields{Major: 8, Minor: 0, Patch: 2},
			args:   args{v2: core.Version{Major: 8, Minor: 0, Patch: 2}},
			want:   0,
		},
		{
			name:   "6.0.36 < 8.0.2",
			fields: fields{Major: 6, Minor: 0, Patch: 36},
			args:   args{v2: core.Version{Major: 8, Minor: 0, Patch: 2}},
			want:   -1,
		},
		{
			name:   "same major, minor diff 8.0.2 < 8.1.0",
			fields: fields{Major: 8, Minor: 0, Patch: 2},
			args:   args{v2: core.Version{Major: 8, Minor: 1, Patch: 0}},
			want:   -1,
		},
		{
			name:   "same major/minor, patch diff 8.0.2 < 8.0.13",
			fields: fields{Major: 8, Minor: 0, Patch: 2},
			args:   args{v2: core.Version{Major: 8, Minor: 0, Patch: 13}},
			want:   -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v1 := core.Version{
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
			name:   "8.0.2",
			fields: fields{Major: 8, Minor: 0, Patch: 2},
			want:   "8.0.2",
		},
		{
			name:   "9.0.13",
			fields: fields{Major: 9, Minor: 0, Patch: 13},
			want:   "9.0.13",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := core.Version{
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
