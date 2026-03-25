package framework_test

import (
	"reflect"
	"testing"

	"github.com/MaineK00n/go-microsoft-version/dotnet/framework"
)

func TestNewVersion(t *testing.T) {
	type args struct {
		ver string
	}
	tests := []struct {
		name    string
		args    args
		want    framework.Version
		wantErr bool
	}{
		{
			name: "4.8.4584.8",
			args: args{ver: "4.8.4584.8"},
			want: framework.Version{Major: 4, Minor: 8, Build: 4584, Revision: 8},
		},
		{
			name: "4.8.4465.0",
			args: args{ver: "4.8.4465.0"},
			want: framework.Version{Major: 4, Minor: 8, Build: 4465, Revision: 0},
		},
		{
			name: "leading zeros 4.8.04584.08",
			args: args{ver: "4.8.04584.08"},
			want: framework.Version{Major: 4, Minor: 8, Build: 4584, Revision: 8},
		},
		{
			name:    "three parts",
			args:    args{ver: "4.8.04584"},
			wantErr: true,
		},
		{
			name:    "non-numeric",
			args:    args{ver: "4.8.abc.08"},
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
			got, err := framework.NewVersion(tt.args.ver)
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
		Major    int
		Minor    int
		Build    int
		Revision int
	}
	type args struct {
		v2 framework.Version
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name:   "4.8.4584.8 = 4.8.4584.8",
			fields: fields{Major: 4, Minor: 8, Build: 4584, Revision: 8},
			args:   args{v2: framework.Version{Major: 4, Minor: 8, Build: 4584, Revision: 8}},
			want:   0,
		},
		{
			name:   "3.0.6920.8954 < 4.8.4584.8",
			fields: fields{Major: 3, Minor: 0, Build: 6920, Revision: 8954},
			args:   args{v2: framework.Version{Major: 4, Minor: 8, Build: 4584, Revision: 8}},
			want:   -1,
		},
		{
			name:   "same major/minor, build diff 4.8.4465.0 < 4.8.4584.8",
			fields: fields{Major: 4, Minor: 8, Build: 4465, Revision: 0},
			args:   args{v2: framework.Version{Major: 4, Minor: 8, Build: 4584, Revision: 8}},
			want:   -1,
		},
		{
			name:   "same major/minor/build, revision diff 4.8.4584.0 < 4.8.4584.8",
			fields: fields{Major: 4, Minor: 8, Build: 4584, Revision: 0},
			args:   args{v2: framework.Version{Major: 4, Minor: 8, Build: 4584, Revision: 8}},
			want:   -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v1 := framework.Version{
				Major:    tt.fields.Major,
				Minor:    tt.fields.Minor,
				Build:    tt.fields.Build,
				Revision: tt.fields.Revision,
			}
			if got := v1.Compare(tt.args.v2); got != tt.want {
				t.Errorf("Version.Compare() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVersion_String(t *testing.T) {
	type fields struct {
		Major    int
		Minor    int
		Build    int
		Revision int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "4.8.4584.8",
			fields: fields{Major: 4, Minor: 8, Build: 4584, Revision: 8},
			want:   "4.8.4584.8",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := framework.Version{
				Major:    tt.fields.Major,
				Minor:    tt.fields.Minor,
				Build:    tt.fields.Build,
				Revision: tt.fields.Revision,
			}
			if got := v.String(); got != tt.want {
				t.Errorf("Version.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
