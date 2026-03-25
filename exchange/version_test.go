package exchange_test

import (
	"reflect"
	"testing"

	"github.com/MaineK00n/go-microsoft-version/exchange"
)

func TestNewVersion(t *testing.T) {
	type args struct {
		ver string
	}
	tests := []struct {
		name    string
		args    args
		want    exchange.Version
		wantErr bool
	}{
		{
			name: "15.0.1497.48",
			args: args{ver: "15.0.1497.48"},
			want: exchange.Version{Major: 15, Minor: 0, Build: 1497, Revision: 48},
		},
		{
			name: "leading zeros 15.01.2375.012",
			args: args{ver: "15.01.2375.012"},
			want: exchange.Version{Major: 15, Minor: 1, Build: 2375, Revision: 12},
		},
		{
			name: "15.02.1544.9",
			args: args{ver: "15.02.1544.9"},
			want: exchange.Version{Major: 15, Minor: 2, Build: 1544, Revision: 9},
		},
		{
			name:    "three parts",
			args:    args{ver: "15.0.1497"},
			wantErr: true,
		},
		{
			name:    "non-numeric",
			args:    args{ver: "15.0.abc.48"},
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
			got, err := exchange.NewVersion(tt.args.ver)
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
		v2 exchange.Version
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name:   "15.1.2375.12 = 15.1.2375.12",
			fields: fields{Major: 15, Minor: 1, Build: 2375, Revision: 12},
			args:   args{v2: exchange.Version{Major: 15, Minor: 1, Build: 2375, Revision: 12}},
			want:   0,
		},
		{
			name:   "15.0.1497.48 < 15.1.2375.12",
			fields: fields{Major: 15, Minor: 0, Build: 1497, Revision: 48},
			args:   args{v2: exchange.Version{Major: 15, Minor: 1, Build: 2375, Revision: 12}},
			want:   -1,
		},
		{
			name:   "same major/minor, build diff 15.1.2375.12 < 15.1.2507.6",
			fields: fields{Major: 15, Minor: 1, Build: 2375, Revision: 12},
			args:   args{v2: exchange.Version{Major: 15, Minor: 1, Build: 2507, Revision: 6}},
			want:   -1,
		},
		{
			name:   "same major/minor/build, revision diff 15.1.2375.12 < 15.1.2375.13",
			fields: fields{Major: 15, Minor: 1, Build: 2375, Revision: 12},
			args:   args{v2: exchange.Version{Major: 15, Minor: 1, Build: 2375, Revision: 13}},
			want:   -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v1 := exchange.Version{
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
			name:   "15.1.2375.12",
			fields: fields{Major: 15, Minor: 1, Build: 2375, Revision: 12},
			want:   "15.1.2375.12",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := exchange.Version{
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
