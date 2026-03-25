package ios_test

import (
	"reflect"
	"testing"

	"github.com/MaineK00n/go-microsoft-version/defender/ios"
)

func TestNewVersion(t *testing.T) {
	type args struct {
		ver string
	}
	tests := []struct {
		name    string
		args    args
		want    ios.Version
		wantErr bool
	}{
		{
			name: "1.1.18090109",
			args: args{ver: "1.1.18090109"},
			want: ios.Version{Major: 1, Minor: 1, Patch: 18090109},
		},
		{
			name: "1.1.57210105",
			args: args{ver: "1.1.57210105"},
			want: ios.Version{Major: 1, Minor: 1, Patch: 57210105},
		},
		{
			name:    "four parts",
			args:    args{ver: "1.1.18090109.0"},
			wantErr: true,
		},
		{
			name:    "non-numeric",
			args:    args{ver: "1.1.abc"},
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
			got, err := ios.NewVersion(tt.args.ver)
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
		Patch int
	}
	type args struct {
		v2 ios.Version
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name:   "1.1.18090109 = 1.1.18090109",
			fields: fields{Major: 1, Minor: 1, Patch: 18090109},
			args:   args{v2: ios.Version{Major: 1, Minor: 1, Patch: 18090109}},
			want:   0,
		},
		{
			name:   "1.1.18090109 < 1.1.57210105",
			fields: fields{Major: 1, Minor: 1, Patch: 18090109},
			args:   args{v2: ios.Version{Major: 1, Minor: 1, Patch: 57210105}},
			want:   -1,
		},
		{
			name:   "same major, minor diff 1.0.18090109 < 1.1.18090109",
			fields: fields{Major: 1, Minor: 0, Patch: 18090109},
			args:   args{v2: ios.Version{Major: 1, Minor: 1, Patch: 18090109}},
			want:   -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v1 := ios.Version{
				Major: tt.fields.Major,
				Minor: tt.fields.Minor,
				Patch: tt.fields.Patch,
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
		Patch int
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name:   "1.1.18090109",
			fields: fields{Major: 1, Minor: 1, Patch: 18090109},
			want:   "1.1.18090109",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := ios.Version{
				Major: tt.fields.Major,
				Minor: tt.fields.Minor,
				Patch: tt.fields.Patch,
			}
			if got := v.String(); got != tt.want {
				t.Errorf("Version.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
