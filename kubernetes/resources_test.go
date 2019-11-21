package kubernetes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResourceParser_ParseYAML(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		wantObj interface{}
		wantErr bool
	}{
		{
			name: "parse regular dashboard",
			args: args{
				path: "test/dashboard.yaml",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			rp := NewResourceParser()
			gotObj, err := rp.ParseYAML(tt.args.path)
			assert.NoError(t, err)
			assert.Equal(t, tt.wantObj, gotObj)
		})
	}
}

func TestResourceParser_GetImages(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name       string
		args       args
		wantImages []string
		wantErr    bool
	}{
		{
			name: "parse dashboard.yaml",
			args: args{
				path: "test/dashboard.yaml",
			},
			wantImages: []string{"k8s.gcr.io/kubernetes-dashboard-amd64:v1.10.1"},
			wantErr:    false,
		},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			rp := NewResourceParser()
			gotImages, err := rp.GetImages(tt.args.path)
			assert.NoError(t, err)
			assert.Equal(t, tt.wantImages, gotImages)
		})
	}
}
