package visualstudio_test

import (
	"reflect"
	"testing"

	"github.com/MaineK00n/go-microsoft-version/visualstudio"
)

func TestNewVersion(t *testing.T) {
	type args struct {
		ver string
	}
	tests := []struct {
		name    string
		args    args
		want    visualstudio.Version
		wantErr bool
	}{
		{
			name: "VS 2022 3-segment: 17.8.6",
			args: args{ver: "17.8.6"},
			want: visualstudio.Version{Type: visualstudio.Modern, Major: 17, Minor: 8, Build: 6},
		},
		{
			name: "VS 2019 3-segment: 16.11.41",
			args: args{ver: "16.11.41"},
			want: visualstudio.Version{Type: visualstudio.Modern, Major: 16, Minor: 11, Build: 41},
		},
		{
			name: "VS 2017 3-segment: 15.9.38",
			args: args{ver: "15.9.38"},
			want: visualstudio.Version{Type: visualstudio.Modern, Major: 15, Minor: 9, Build: 38},
		},
		{
			name: "VS 2015 4-segment: 14.0.27552.0",
			args: args{ver: "14.0.27552.0"},
			want: visualstudio.Version{Type: visualstudio.Legacy, Major: 14, Minor: 0, Build: 27552, Revision: 0},
		},
		{
			name: "VS 2013 4-segment: 12.0.40700.0",
			args: args{ver: "12.0.40700.0"},
			want: visualstudio.Version{Type: visualstudio.Legacy, Major: 12, Minor: 0, Build: 40700, Revision: 0},
		},
		{
			name: "VS 2012 4-segment: 11.0.61234.0",
			args: args{ver: "11.0.61234.0"},
			want: visualstudio.Version{Type: visualstudio.Legacy, Major: 11, Minor: 0, Build: 61234, Revision: 0},
		},
		{
			name:    "invalid 2-segment",
			args:    args{ver: "17.8"},
			wantErr: true,
		},
		{
			name:    "invalid 5-segment",
			args:    args{ver: "17.8.6.1.0"},
			wantErr: true,
		},
		{
			name:    "non-numeric",
			args:    args{ver: "17.abc.6"},
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
			got, err := visualstudio.NewVersion(tt.args.ver)
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
		Type     visualstudio.VersionType
		Major    int
		Minor    int
		Build    int
		Revision int
	}
	type args struct {
		v2 visualstudio.Version
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name:   "modern: 17.8.6 = 17.8.6",
			fields: fields{Type: visualstudio.Modern, Major: 17, Minor: 8, Build: 6},
			args:   args{v2: visualstudio.Version{Type: visualstudio.Modern, Major: 17, Minor: 8, Build: 6}},
			want:   0,
		},
		{
			name:   "modern: 15.9.38 < 17.8.6",
			fields: fields{Type: visualstudio.Modern, Major: 15, Minor: 9, Build: 38},
			args:   args{v2: visualstudio.Version{Type: visualstudio.Modern, Major: 17, Minor: 8, Build: 6}},
			want:   -1,
		},
		{
			name:   "modern: patch diff 17.8.7 > 17.8.6",
			fields: fields{Type: visualstudio.Modern, Major: 17, Minor: 8, Build: 7},
			args:   args{v2: visualstudio.Version{Type: visualstudio.Modern, Major: 17, Minor: 8, Build: 6}},
			want:   1,
		},
		{
			name:   "legacy: 14.0.27552.0 = 14.0.27552.0",
			fields: fields{Type: visualstudio.Legacy, Major: 14, Minor: 0, Build: 27552, Revision: 0},
			args:   args{v2: visualstudio.Version{Type: visualstudio.Legacy, Major: 14, Minor: 0, Build: 27552, Revision: 0}},
			want:   0,
		},
		{
			name:   "legacy: major diff 12.0.40700.0 < 14.0.27552.0",
			fields: fields{Type: visualstudio.Legacy, Major: 12, Minor: 0, Build: 40700, Revision: 0},
			args:   args{v2: visualstudio.Version{Type: visualstudio.Legacy, Major: 14, Minor: 0, Build: 27552, Revision: 0}},
			want:   -1,
		},
		{
			name:   "legacy: revision diff 14.0.27552.1 > 14.0.27552.0",
			fields: fields{Type: visualstudio.Legacy, Major: 14, Minor: 0, Build: 27552, Revision: 1},
			args:   args{v2: visualstudio.Version{Type: visualstudio.Legacy, Major: 14, Minor: 0, Build: 27552, Revision: 0}},
			want:   1,
		},
		{
			name:   "cross-type: Legacy < Modern",
			fields: fields{Type: visualstudio.Legacy, Major: 14, Minor: 0, Build: 27552, Revision: 0},
			args:   args{v2: visualstudio.Version{Type: visualstudio.Modern, Major: 15, Minor: 9, Build: 38}},
			want:   -1,
		},
		{
			name:   "cross-type: Modern > Legacy",
			fields: fields{Type: visualstudio.Modern, Major: 17, Minor: 8, Build: 6},
			args:   args{v2: visualstudio.Version{Type: visualstudio.Legacy, Major: 14, Minor: 0, Build: 27552, Revision: 0}},
			want:   1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v1 := visualstudio.Version{
				Type:     tt.fields.Type,
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
		Type     visualstudio.VersionType
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
			name:   "modern: 17.8.6",
			fields: fields{Type: visualstudio.Modern, Major: 17, Minor: 8, Build: 6},
			want:   "17.8.6",
		},
		{
			name:   "legacy: 14.0.27552.0",
			fields: fields{Type: visualstudio.Legacy, Major: 14, Minor: 0, Build: 27552, Revision: 0},
			want:   "14.0.27552.0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := visualstudio.Version{
				Type:     tt.fields.Type,
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
