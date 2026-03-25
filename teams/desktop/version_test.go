package desktop_test

import (
	"reflect"
	"testing"

	"github.com/MaineK00n/go-microsoft-version/teams/desktop"
)

func TestNewVersion(t *testing.T) {
	type args struct {
		ver string
	}
	tests := []struct {
		name    string
		args    args
		want    desktop.Version
		wantErr bool
	}{
		{
			name: "25122.1415.3698.6812",
			args: args{ver: "25122.1415.3698.6812"},
			want: desktop.Version{Major: 25122, Minor: 1415, Build: 3698, Revision: 6812},
		},
		{
			name: "1.6.0.18681",
			args: args{ver: "1.6.0.18681"},
			want: desktop.Version{Major: 1, Minor: 6, Build: 0, Revision: 18681},
		},
		{
			name: "leading zeros 1.6.00.18681",
			args: args{ver: "1.6.00.18681"},
			want: desktop.Version{Major: 1, Minor: 6, Build: 0, Revision: 18681},
		},
		{
			name:    "three parts",
			args:    args{ver: "25122.1415.3698"},
			wantErr: true,
		},
		{
			name:    "non-numeric",
			args:    args{ver: "25122.abc.3698.6812"},
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
			got, err := desktop.NewVersion(tt.args.ver)
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
		v2 desktop.Version
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name:   "25122.1415.3698.6812 = 25122.1415.3698.6812",
			fields: fields{Major: 25122, Minor: 1415, Build: 3698, Revision: 6812},
			args:   args{v2: desktop.Version{Major: 25122, Minor: 1415, Build: 3698, Revision: 6812}},
			want:   0,
		},
		{
			name:   "1.6.0.18681 < 25122.1415.3698.6812",
			fields: fields{Major: 1, Minor: 6, Build: 0, Revision: 18681},
			args:   args{v2: desktop.Version{Major: 25122, Minor: 1415, Build: 3698, Revision: 6812}},
			want:   -1,
		},
		{
			name:   "same major/minor/build, revision diff 25122.1415.3698.6812 < 25122.1415.3698.7000",
			fields: fields{Major: 25122, Minor: 1415, Build: 3698, Revision: 6812},
			args:   args{v2: desktop.Version{Major: 25122, Minor: 1415, Build: 3698, Revision: 7000}},
			want:   -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v1 := desktop.Version{
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
			name:   "25122.1415.3698.6812",
			fields: fields{Major: 25122, Minor: 1415, Build: 3698, Revision: 6812},
			want:   "25122.1415.3698.6812",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := desktop.Version{
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
