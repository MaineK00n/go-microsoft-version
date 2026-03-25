package android_test

import (
	"reflect"
	"testing"

	"github.com/MaineK00n/go-microsoft-version/teams/android"
)

func TestNewVersion(t *testing.T) {
	type args struct {
		ver string
	}
	tests := []struct {
		name    string
		args    args
		want    android.Version
		wantErr bool
	}{
		{
			name: "1.0.0.2024022302",
			args: args{ver: "1.0.0.2024022302"},
			want: android.Version{Major: 1, Minor: 0, Build: 0, Revision: 2024022302},
		},
		{
			name: "1.0.0.2025042801",
			args: args{ver: "1.0.0.2025042801"},
			want: android.Version{Major: 1, Minor: 0, Build: 0, Revision: 2025042801},
		},
		{
			name:    "three parts",
			args:    args{ver: "1.0.0"},
			wantErr: true,
		},
		{
			name:    "non-numeric",
			args:    args{ver: "1.0.0.abc"},
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
			got, err := android.NewVersion(tt.args.ver)
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
		v2 android.Version
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name:   "1.0.0.2024022302 = 1.0.0.2024022302",
			fields: fields{Major: 1, Minor: 0, Build: 0, Revision: 2024022302},
			args:   args{v2: android.Version{Major: 1, Minor: 0, Build: 0, Revision: 2024022302}},
			want:   0,
		},
		{
			name:   "1.0.0.2024022302 < 1.0.0.2025042801",
			fields: fields{Major: 1, Minor: 0, Build: 0, Revision: 2024022302},
			args:   args{v2: android.Version{Major: 1, Minor: 0, Build: 0, Revision: 2025042801}},
			want:   -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v1 := android.Version{
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
			name:   "1.0.0.2024022302",
			fields: fields{Major: 1, Minor: 0, Build: 0, Revision: 2024022302},
			want:   "1.0.0.2024022302",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := android.Version{
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
