package mac_test

import (
	"reflect"
	"testing"

	"github.com/MaineK00n/go-microsoft-version/teams/mac"
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
			name: "1.6.0.17554",
			args: args{ver: "1.6.0.17554"},
			want: mac.Version{Major: 1, Minor: 6, Build: 0, Revision: 17554},
		},
		{
			name: "24295.606.3252.8961",
			args: args{ver: "24295.606.3252.8961"},
			want: mac.Version{Major: 24295, Minor: 606, Build: 3252, Revision: 8961},
		},
		{
			name: "leading zeros 1.6.00.17554",
			args: args{ver: "1.6.00.17554"},
			want: mac.Version{Major: 1, Minor: 6, Build: 0, Revision: 17554},
		},
		{
			name:    "three parts",
			args:    args{ver: "1.6.00"},
			wantErr: true,
		},
		{
			name:    "non-numeric",
			args:    args{ver: "1.6.abc.17554"},
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
		Major    int
		Minor    int
		Build    int
		Revision int
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
			name:   "1.6.0.17554 = 1.6.0.17554",
			fields: fields{Major: 1, Minor: 6, Build: 0, Revision: 17554},
			args:   args{v2: mac.Version{Major: 1, Minor: 6, Build: 0, Revision: 17554}},
			want:   0,
		},
		{
			name:   "1.6.0.17554 < 24295.606.3252.8961",
			fields: fields{Major: 1, Minor: 6, Build: 0, Revision: 17554},
			args:   args{v2: mac.Version{Major: 24295, Minor: 606, Build: 3252, Revision: 8961}},
			want:   -1,
		},
		{
			name:   "same major/minor/build, revision diff 1.6.0.17554 < 1.6.0.26474",
			fields: fields{Major: 1, Minor: 6, Build: 0, Revision: 17554},
			args:   args{v2: mac.Version{Major: 1, Minor: 6, Build: 0, Revision: 26474}},
			want:   -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v1 := mac.Version{
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
			name:   "24295.606.3252.8961",
			fields: fields{Major: 24295, Minor: 606, Build: 3252, Revision: 8961},
			want:   "24295.606.3252.8961",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := mac.Version{
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
