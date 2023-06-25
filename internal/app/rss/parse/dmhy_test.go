package parse

import (
	"context"
	"os"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

// go test ./internal/app/rss/parse -v -run TestDmhy_Run$ --count=1
func TestDmhy_Run(t *testing.T) {
	type fields struct {
		data []byte
	}
	type args struct {
		ctx context.Context
	}

	ctx := context.Background()
	file := "testdata/dmhy.rss"
	data, err := os.ReadFile(file)
	assert.Nil(t, err)
	if err != nil {
		return
	}

	a := args{ctx: ctx}

	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *data.List
		wantErr bool
	}{
		{
			name: "正常解析",
			fields: fields{
				data: data,
			},
			args:    a,
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dmhy{
				data: tt.fields.data,
			}
			got, err := d.Run(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Run() got = %v, want %v", got, tt.want)
			}
		})
	}
}
