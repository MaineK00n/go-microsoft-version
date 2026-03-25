package ie_test

import (
	"reflect"
	"testing"

	"github.com/MaineK00n/go-microsoft-version/ie"
)

func TestNewVersion(t *testing.T) {
	type args struct {
		ver string
	}
	tests := []struct {
		name    string
		args    args
		want    ie.Version
		wantErr bool
	}{
		{
			name: "11.0.9600.21117",
			args: args{ver: "11.0.9600.21117"},
			want: ie.Version{Major: 11, Minor: 0, Build: 9600, Revision: 21117},
		},
		{
			name: "9.0.8112.21631",
			args: args{ver: "9.0.8112.21631"},
			want: ie.Version{Major: 9, Minor: 0, Build: 8112, Revision: 21631},
		},
		{
			name:    "three parts",
			args:    args{ver: "11.0.9600"},
			wantErr: true,
		},
		{
			name:    "non-numeric",
			args:    args{ver: "11.0.abc.21117"},
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
			got, err := ie.NewVersion(tt.args.ver)
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
		v2 ie.Version
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name:   "11.0.9600.21117 = 11.0.9600.21117",
			fields: fields{Major: 11, Minor: 0, Build: 9600, Revision: 21117},
			args:   args{v2: ie.Version{Major: 11, Minor: 0, Build: 9600, Revision: 21117}},
			want:   0,
		},
		{
			name:   "9.0.8112.21631 < 11.0.9600.21117",
			fields: fields{Major: 9, Minor: 0, Build: 8112, Revision: 21631},
			args:   args{v2: ie.Version{Major: 11, Minor: 0, Build: 9600, Revision: 21117}},
			want:   -1,
		},
		{
			name:   "same major/minor, build diff 11.0.9600.21117 < 11.0.9601.0",
			fields: fields{Major: 11, Minor: 0, Build: 9600, Revision: 21117},
			args:   args{v2: ie.Version{Major: 11, Minor: 0, Build: 9601, Revision: 0}},
			want:   -1,
		},
		{
			name:   "same major/minor/build, revision diff 11.0.9600.21117 < 11.0.9600.22000",
			fields: fields{Major: 11, Minor: 0, Build: 9600, Revision: 21117},
			args:   args{v2: ie.Version{Major: 11, Minor: 0, Build: 9600, Revision: 22000}},
			want:   -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v1 := ie.Version{
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
			name:   "11.0.9600.21117",
			fields: fields{Major: 11, Minor: 0, Build: 9600, Revision: 21117},
			want:   "11.0.9600.21117",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := ie.Version{
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
