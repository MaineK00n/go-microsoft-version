package edge_test

import (
	"reflect"
	"testing"

	"github.com/MaineK00n/go-microsoft-version/edge"
)

func TestNewVersion(t *testing.T) {
	type args struct {
		ver string
	}
	tests := []struct {
		name    string
		args    args
		want    edge.Version
		wantErr bool
	}{
		{
			name: "EdgeHTML: 20.10240",
			args: args{ver: "20.10240"},
			want: edge.Version{Type: edge.EdgeHTML, Major: 20, Build: 10240},
		},
		{
			name: "EdgeHTML: trailing space 44.17763 ",
			args: args{ver: "44.17763 "},
			want: edge.Version{Type: edge.EdgeHTML, Major: 44, Build: 17763},
		},
		{
			name: "Chromium: 88.0.705.18",
			args: args{ver: "88.0.705.18"},
			want: edge.Version{Type: edge.Chromium, Major: 88, Minor: 0, Build: 705, Revision: 18},
		},
		{
			name:    "single number",
			args:    args{ver: "88"},
			wantErr: true,
		},
		{
			name:    "three parts",
			args:    args{ver: "1.2.3"},
			wantErr: true,
		},
		{
			name:    "non-numeric",
			args:    args{ver: "abc.def"},
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
			got, err := edge.NewVersion(tt.args.ver)
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
		Type     edge.VersionType
		Major    int
		Minor    int
		Build    int
		Revision int
	}
	type args struct {
		v2 edge.Version
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   int
	}{
		{
			name:   "EdgeHTML: 20.10240 = 20.10240",
			fields: fields{Type: edge.EdgeHTML, Major: 20, Build: 10240},
			args:   args{v2: edge.Version{Type: edge.EdgeHTML, Major: 20, Build: 10240}},
			want:   0,
		},
		{
			name:   "EdgeHTML: 20.10240 < 44.17763",
			fields: fields{Type: edge.EdgeHTML, Major: 20, Build: 10240},
			args:   args{v2: edge.Version{Type: edge.EdgeHTML, Major: 44, Build: 17763}},
			want:   -1,
		},
		{
			name:   "EdgeHTML: same major, 20.10240 < 20.10586",
			fields: fields{Type: edge.EdgeHTML, Major: 20, Build: 10240},
			args:   args{v2: edge.Version{Type: edge.EdgeHTML, Major: 20, Build: 10586}},
			want:   -1,
		},
		{
			name:   "Chromium: 88.0.705.18 = 88.0.705.18",
			fields: fields{Type: edge.Chromium, Major: 88, Minor: 0, Build: 705, Revision: 18},
			args:   args{v2: edge.Version{Type: edge.Chromium, Major: 88, Minor: 0, Build: 705, Revision: 18}},
			want:   0,
		},
		{
			name:   "Chromium: 88.0.705.18 < 146.0.3856.13",
			fields: fields{Type: edge.Chromium, Major: 88, Minor: 0, Build: 705, Revision: 18},
			args:   args{v2: edge.Version{Type: edge.Chromium, Major: 146, Minor: 0, Build: 3856, Revision: 13}},
			want:   -1,
		},
		{
			name:   "Chromium: same major, minor diff 88.0.705.18 < 88.1.705.18",
			fields: fields{Type: edge.Chromium, Major: 88, Minor: 0, Build: 705, Revision: 18},
			args:   args{v2: edge.Version{Type: edge.Chromium, Major: 88, Minor: 1, Build: 705, Revision: 18}},
			want:   -1,
		},
		{
			name:   "Chromium: same major/minor, build diff 88.0.705.18 < 88.0.706.18",
			fields: fields{Type: edge.Chromium, Major: 88, Minor: 0, Build: 705, Revision: 18},
			args:   args{v2: edge.Version{Type: edge.Chromium, Major: 88, Minor: 0, Build: 706, Revision: 18}},
			want:   -1,
		},
		{
			name:   "Chromium: same major/minor/build, revision diff 88.0.705.18 < 88.0.705.19",
			fields: fields{Type: edge.Chromium, Major: 88, Minor: 0, Build: 705, Revision: 18},
			args:   args{v2: edge.Version{Type: edge.Chromium, Major: 88, Minor: 0, Build: 705, Revision: 19}},
			want:   -1,
		},
		{
			name:   "cross-type: EdgeHTML < Chromium",
			fields: fields{Type: edge.EdgeHTML, Major: 44, Build: 17763},
			args:   args{v2: edge.Version{Type: edge.Chromium, Major: 88, Minor: 0, Build: 705, Revision: 18}},
			want:   -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v1 := edge.Version{
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
		Type     edge.VersionType
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
			name:   "EdgeHTML: 20.10240",
			fields: fields{Type: edge.EdgeHTML, Major: 20, Build: 10240},
			want:   "20.10240",
		},
		{
			name:   "Chromium: 88.0.705.18",
			fields: fields{Type: edge.Chromium, Major: 88, Minor: 0, Build: 705, Revision: 18},
			want:   "88.0.705.18",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := edge.Version{
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
