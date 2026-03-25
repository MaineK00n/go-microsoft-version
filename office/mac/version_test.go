package mac_test

import (
	"reflect"
	"testing"

	"github.com/MaineK00n/go-microsoft-version/office/mac"
)

func TestNewVersion(t *testing.T) {
	type args struct {
		ver string
	}
	tests := []struct {
		name    string
		args    args
		want    mac.Version
		wantErr bool
	}{
		{
			name: "16.100.25081015",
			args: args{ver: "16.100.25081015"},
			want: mac.Version{Major: 16, Minor: 100, Build: 25081015},
		},
		{
			name: "16.54.21101001",
			args: args{ver: "16.54.21101001"},
			want: mac.Version{Major: 16, Minor: 54, Build: 21101001},
		},
		{
			name:    "four parts",
			args:    args{ver: "16.50.210.613"},
			wantErr: true,
		},
		{
			name:    "non-numeric",
			args:    args{ver: "16.50.abc"},
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
			got, err := mac.NewVersion(tt.args.ver)
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
		Build int
	}
	type args struct {
		v2 mac.Version
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name:   "16.100.25081015 = 16.100.25081015",
			fields: fields{Major: 16, Minor: 100, Build: 25081015},
			args:   args{v2: mac.Version{Major: 16, Minor: 100, Build: 25081015}},
			want:   0,
		},
		{
			name:   "16.54.21101001 < 16.100.25081015",
			fields: fields{Major: 16, Minor: 54, Build: 21101001},
			args:   args{v2: mac.Version{Major: 16, Minor: 100, Build: 25081015}},
			want:   -1,
		},
		{
			name:   "same major/minor, build diff 16.100.25031020 < 16.100.25081015",
			fields: fields{Major: 16, Minor: 100, Build: 25031020},
			args:   args{v2: mac.Version{Major: 16, Minor: 100, Build: 25081015}},
			want:   -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v1 := mac.Version{
				Major: tt.fields.Major,
				Minor: tt.fields.Minor,
				Build: tt.fields.Build,
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
		Build int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "16.100.25081015",
			fields: fields{Major: 16, Minor: 100, Build: 25081015},
			want:   "16.100.25081015",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := mac.Version{
				Major: tt.fields.Major,
				Minor: tt.fields.Minor,
				Build: tt.fields.Build,
			}
			if got := v.String(); got != tt.want {
				t.Errorf("Version.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
