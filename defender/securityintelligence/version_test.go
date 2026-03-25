package securityintelligence_test

import (
	"reflect"
	"testing"

	"github.com/MaineK00n/go-microsoft-version/defender/securityintelligence"
)

func TestNewVersion(t *testing.T) {
	type args struct {
		ver string
	}
	tests := []struct {
		name    string
		args    args
		want    securityintelligence.Version
		wantErr bool
	}{
		{
			name: "1.379.200.0",
			args: args{ver: "1.379.200.0"},
			want: securityintelligence.Version{Major: 1, Minor: 379, Build: 200, Revision: 0},
		},
		{
			name: "1.411.234.0",
			args: args{ver: "1.411.234.0"},
			want: securityintelligence.Version{Major: 1, Minor: 411, Build: 234, Revision: 0},
		},
		{
			name:    "three parts",
			args:    args{ver: "1.379.200"},
			wantErr: true,
		},
		{
			name:    "non-numeric",
			args:    args{ver: "1.abc.200.0"},
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
			got, err := securityintelligence.NewVersion(tt.args.ver)
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
		v2 securityintelligence.Version
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name:   "1.379.200.0 = 1.379.200.0",
			fields: fields{Major: 1, Minor: 379, Build: 200, Revision: 0},
			args:   args{v2: securityintelligence.Version{Major: 1, Minor: 379, Build: 200, Revision: 0}},
			want:   0,
		},
		{
			name:   "1.379.200.0 < 1.411.234.0",
			fields: fields{Major: 1, Minor: 379, Build: 200, Revision: 0},
			args:   args{v2: securityintelligence.Version{Major: 1, Minor: 411, Build: 234, Revision: 0}},
			want:   -1,
		},
		{
			name:   "same major/minor, build diff 1.411.200.0 < 1.411.234.0",
			fields: fields{Major: 1, Minor: 411, Build: 200, Revision: 0},
			args:   args{v2: securityintelligence.Version{Major: 1, Minor: 411, Build: 234, Revision: 0}},
			want:   -1,
		},
		{
			name:   "same major/minor/build, revision diff 1.411.234.0 < 1.411.234.1",
			fields: fields{Major: 1, Minor: 411, Build: 234, Revision: 0},
			args:   args{v2: securityintelligence.Version{Major: 1, Minor: 411, Build: 234, Revision: 1}},
			want:   -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v1 := securityintelligence.Version{
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
			name:   "1.379.200.0",
			fields: fields{Major: 1, Minor: 379, Build: 200, Revision: 0},
			want:   "1.379.200.0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := securityintelligence.Version{
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
