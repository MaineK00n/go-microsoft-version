package sqlserver_test

import (
	"reflect"
	"testing"

	"github.com/MaineK00n/go-microsoft-version/sqlserver"
)

func TestNewVersion(t *testing.T) {
	type args struct {
		ver string
	}
	tests := []struct {
		name    string
		args    args
		want    sqlserver.Version
		wantErr bool
	}{
		{
			name: "15.0.2095.3",
			args: args{ver: "15.0.2095.3"},
			want: sqlserver.Version{Major: 15, Minor: 0, Build: 2095, Revision: 3},
		},
		{
			name: "16.0.4185.3",
			args: args{ver: "16.0.4185.3"},
			want: sqlserver.Version{Major: 16, Minor: 0, Build: 4185, Revision: 3},
		},
		{
			name:    "three parts",
			args:    args{ver: "15.0.2095"},
			wantErr: true,
		},
		{
			name:    "non-numeric",
			args:    args{ver: "15.0.abc.3"},
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
			got, err := sqlserver.NewVersion(tt.args.ver)
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
		v2 sqlserver.Version
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name:   "15.0.2095.3 = 15.0.2095.3",
			fields: fields{Major: 15, Minor: 0, Build: 2095, Revision: 3},
			args:   args{v2: sqlserver.Version{Major: 15, Minor: 0, Build: 2095, Revision: 3}},
			want:   0,
		},
		{
			name:   "14.0.3465.1 < 15.0.2095.3",
			fields: fields{Major: 14, Minor: 0, Build: 3465, Revision: 1},
			args:   args{v2: sqlserver.Version{Major: 15, Minor: 0, Build: 2095, Revision: 3}},
			want:   -1,
		},
		{
			name:   "same major/minor, build diff 16.0.4115.5 < 16.0.4185.3",
			fields: fields{Major: 16, Minor: 0, Build: 4115, Revision: 5},
			args:   args{v2: sqlserver.Version{Major: 16, Minor: 0, Build: 4185, Revision: 3}},
			want:   -1,
		},
		{
			name:   "same major/minor/build, revision diff 15.0.2095.3 < 15.0.2095.5",
			fields: fields{Major: 15, Minor: 0, Build: 2095, Revision: 3},
			args:   args{v2: sqlserver.Version{Major: 15, Minor: 0, Build: 2095, Revision: 5}},
			want:   -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v1 := sqlserver.Version{
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
			name:   "15.0.2095.3",
			fields: fields{Major: 15, Minor: 0, Build: 2095, Revision: 3},
			want:   "15.0.2095.3",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := sqlserver.Version{
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
