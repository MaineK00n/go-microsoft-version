package windows_test

import (
	"reflect"
	"testing"

	"github.com/MaineK00n/go-microsoft-version/defender/windows"
)

func TestNewVersion(t *testing.T) {
	type args struct {
		ver string
	}
	tests := []struct {
		name    string
		args    args
		want    windows.Version
		wantErr bool
	}{
		{
			name: "4.18.23100.2009",
			args: args{ver: "4.18.23100.2009"},
			want: windows.Version{Major: 4, Minor: 18, Build: 23100, Revision: 2009},
		},
		{
			name: "4.18.24070.5",
			args: args{ver: "4.18.24070.5"},
			want: windows.Version{Major: 4, Minor: 18, Build: 24070, Revision: 5},
		},
		{
			name:    "three parts",
			args:    args{ver: "4.18.23100"},
			wantErr: true,
		},
		{
			name:    "non-numeric",
			args:    args{ver: "4.18.abc.2009"},
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
			got, err := windows.NewVersion(tt.args.ver)
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
		v2 windows.Version
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name:   "4.18.23100.2009 = 4.18.23100.2009",
			fields: fields{Major: 4, Minor: 18, Build: 23100, Revision: 2009},
			args:   args{v2: windows.Version{Major: 4, Minor: 18, Build: 23100, Revision: 2009}},
			want:   0,
		},
		{
			name:   "4.18.23100.2009 < 4.18.24070.5",
			fields: fields{Major: 4, Minor: 18, Build: 23100, Revision: 2009},
			args:   args{v2: windows.Version{Major: 4, Minor: 18, Build: 24070, Revision: 5}},
			want:   -1,
		},
		{
			name:   "same major/minor/build, revision diff 4.18.23100.2009 < 4.18.23100.3000",
			fields: fields{Major: 4, Minor: 18, Build: 23100, Revision: 2009},
			args:   args{v2: windows.Version{Major: 4, Minor: 18, Build: 23100, Revision: 3000}},
			want:   -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v1 := windows.Version{
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
			name:   "4.18.23100.2009",
			fields: fields{Major: 4, Minor: 18, Build: 23100, Revision: 2009},
			want:   "4.18.23100.2009",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := windows.Version{
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
