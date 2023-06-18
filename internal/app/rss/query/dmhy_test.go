package query

import (
	"context"
	"github.com/anguloc/zet/internal/app/rss/client"
	"github.com/anguloc/zet/internal/dto"
	"reflect"
	"testing"
)

func TestDmhy_Command(t *testing.T) {
	t.Skip("test")
	type fields struct {
		url    string
		client client.IClient
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dmhy{
				url:    tt.fields.url,
				client: tt.fields.client,
			}
			if err := d.Command(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Command() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

// go test ./internal/app/rss/query -v -run TestDmhy_Get$ --count=1
func TestDmhy_Get(t *testing.T) {
	t.Skip("test")
	ctx := context.Background()
	type fields struct {
		url    string
		client client.IClient
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Result
		wantErr bool
	}{
		{
			name: "正常请求",
			fields: fields{
				url:    "",
				client: client.New(client.WithModule(dto.DmhyModule)),
			},
			args: args{
				ctx: ctx,
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dmhy{
				url:    tt.fields.url,
				client: tt.fields.client,
			}
			got, err := d.Get(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Get() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDmhy_getUrl(t *testing.T) {
	t.Skip("test")
	type fields struct {
		url    string
		client client.IClient
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dmhy{
				url:    tt.fields.url,
				client: tt.fields.client,
			}
			if got := d.getUrl(); got != tt.want {
				t.Errorf("getUrl() = %v, want %v", got, tt.want)
			}
		})
	}
}
