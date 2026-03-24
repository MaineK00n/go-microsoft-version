package windows_test

import (
	"reflect"
	"testing"

	"github.com/MaineK00n/go-microsoft-version/windows"
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
			name: "3-part: 6.1.7601",
			args: args{ver: "6.1.7601"},
			want: windows.Version{Major: 6, Minor: 1, Build: 7601},
		},
		{
			name: "4-part: 10.0.19045.7058",
			args: args{ver: "10.0.19045.7058"},
			want: windows.Version{Major: 10, Minor: 0, Build: 19045, Revision: new(7058)},
		},
		{
			name: "trailing space: 10.0.26100.8037 ",
			args: args{ver: "10.0.26100.8037 "},
			want: windows.Version{Major: 10, Minor: 0, Build: 26100, Revision: new(8037)},
		},
		{
			name:    "2-part",
			args:    args{ver: "10.0"},
			wantErr: true,
		},
		{
			name:    "non-numeric",
			args:    args{ver: "a.b.c"},
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
		Revision *int
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
			name:   "6.1.7601 = 6.1.7601",
			fields: fields{Major: 6, Minor: 1, Build: 7601},
			args:   args{v2: windows.Version{Major: 6, Minor: 1, Build: 7601}},
			want:   0,
		},
		{
			name:   "6.1.7601 < 6.2.9200",
			fields: fields{Major: 6, Minor: 1, Build: 7601},
			args:   args{v2: windows.Version{Major: 6, Minor: 2, Build: 9200}},
			want:   -1,
		},
		{
			name:   "10.0.19045 < 10.0.22000",
			fields: fields{Major: 10, Minor: 0, Build: 19045},
			args:   args{v2: windows.Version{Major: 10, Minor: 0, Build: 22000}},
			want:   -1,
		},
		{
			name:   "10.0.19045.7058 = 10.0.19045.7058",
			fields: fields{Major: 10, Minor: 0, Build: 19045, Revision: new(7058)},
			args:   args{v2: windows.Version{Major: 10, Minor: 0, Build: 19045, Revision: new(7058)}},
			want:   0,
		},
		{
			name:   "revision diff: 10.0.19045.4894 < 10.0.19045.7058",
			fields: fields{Major: 10, Minor: 0, Build: 19045, Revision: new(4894)},
			args:   args{v2: windows.Version{Major: 10, Minor: 0, Build: 19045, Revision: new(7058)}},
			want:   -1,
		},
		{
			name:   "major diff: 6.3.9600 < 10.0.19045",
			fields: fields{Major: 6, Minor: 3, Build: 9600},
			args:   args{v2: windows.Version{Major: 10, Minor: 0, Build: 19045}},
			want:   -1,
		},
		{
			name:   "3-part vs 4-part: 10.0.19045 = 10.0.19045.7058",
			fields: fields{Major: 10, Minor: 0, Build: 19045},
			args:   args{v2: windows.Version{Major: 10, Minor: 0, Build: 19045, Revision: new(7058)}},
			want:   0,
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
		Revision *int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "3-part: 6.1.7601",
			fields: fields{Major: 6, Minor: 1, Build: 7601},
			want:   "6.1.7601",
		},
		{
			name:   "4-part: 10.0.19045.7058",
			fields: fields{Major: 10, Minor: 0, Build: 19045, Revision: new(7058)},
			want:   "10.0.19045.7058",
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
