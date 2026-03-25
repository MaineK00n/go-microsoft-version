package vscode_test

import (
	"reflect"
	"testing"

	"github.com/MaineK00n/go-microsoft-version/vscode"
)

func TestNewVersion(t *testing.T) {
	type args struct {
		ver string
	}
	tests := []struct {
		name    string
		args    args
		want    vscode.Version
		wantErr bool
	}{
		{
			name: "1.100.1",
			args: args{ver: "1.100.1"},
			want: vscode.Version{Major: 1, Minor: 100, Patch: 1},
		},
		{
			name: "1.104.0",
			args: args{ver: "1.104.0"},
			want: vscode.Version{Major: 1, Minor: 104, Patch: 0},
		},
		{
			name:    "two parts",
			args:    args{ver: "1.100"},
			wantErr: true,
		},
		{
			name:    "non-numeric",
			args:    args{ver: "1.abc.0"},
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
			got, err := vscode.NewVersion(tt.args.ver)
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
		v2 vscode.Version
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name:   "1.100.1 = 1.100.1",
			fields: fields{Major: 1, Minor: 100, Patch: 1},
			args:   args{v2: vscode.Version{Major: 1, Minor: 100, Patch: 1}},
			want:   0,
		},
		{
			name:   "1.56.0 < 1.100.1",
			fields: fields{Major: 1, Minor: 56, Patch: 0},
			args:   args{v2: vscode.Version{Major: 1, Minor: 100, Patch: 1}},
			want:   -1,
		},
		{
			name:   "same major/minor, patch diff 1.100.1 < 1.100.2",
			fields: fields{Major: 1, Minor: 100, Patch: 1},
			args:   args{v2: vscode.Version{Major: 1, Minor: 100, Patch: 2}},
			want:   -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v1 := vscode.Version{
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
			name:   "1.100.1",
			fields: fields{Major: 1, Minor: 100, Patch: 1},
			want:   "1.100.1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := vscode.Version{
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
