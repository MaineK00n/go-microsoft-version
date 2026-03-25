package windows_test

import (
	"reflect"
	"testing"

	"github.com/MaineK00n/go-microsoft-version/office/windows"
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
			name: "16.0.5474.1001",
			args: args{ver: "16.0.5474.1001"},
			want: windows.Version{Major: 16, Minor: 0, Build: 5474, Revision: 1001},
		},
		{
			name: "15.0.5589.1002",
			args: args{ver: "15.0.5589.1002"},
			want: windows.Version{Major: 15, Minor: 0, Build: 5589, Revision: 1002},
		},
		{
			name:    "three parts",
			args:    args{ver: "16.0.5474"},
			wantErr: true,
		},
		{
			name:    "non-numeric",
			args:    args{ver: "16.0.xxxx.1000"},
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
			name:   "16.0.5474.1001 = 16.0.5474.1001",
			fields: fields{Major: 16, Minor: 0, Build: 5474, Revision: 1001},
			args:   args{v2: windows.Version{Major: 16, Minor: 0, Build: 5474, Revision: 1001}},
			want:   0,
		},
		{
			name:   "15.0.5589.1002 < 16.0.5474.1001",
			fields: fields{Major: 15, Minor: 0, Build: 5589, Revision: 1002},
			args:   args{v2: windows.Version{Major: 16, Minor: 0, Build: 5474, Revision: 1001}},
			want:   -1,
		},
		{
			name:   "same major/minor, build diff 16.0.5215.1000 < 16.0.5474.1001",
			fields: fields{Major: 16, Minor: 0, Build: 5215, Revision: 1000},
			args:   args{v2: windows.Version{Major: 16, Minor: 0, Build: 5474, Revision: 1001}},
			want:   -1,
		},
		{
			name:   "same major/minor/build, revision diff 16.0.5474.1001 < 16.0.5474.1002",
			fields: fields{Major: 16, Minor: 0, Build: 5474, Revision: 1001},
			args:   args{v2: windows.Version{Major: 16, Minor: 0, Build: 5474, Revision: 1002}},
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
			name:   "16.0.5474.1001",
			fields: fields{Major: 16, Minor: 0, Build: 5474, Revision: 1001},
			want:   "16.0.5474.1001",
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
