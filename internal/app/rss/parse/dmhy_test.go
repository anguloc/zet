package parse

import (
	"context"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/anguloc/zet/internal/app/rss/data"
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
	fileData, err := os.ReadFile(file)
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
				data: fileData,
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

// go test ./internal/app/rss/parse -v -run TestDmhy_parsePubDate$ --count=1
func TestDmhy_parsePubDate(t *testing.T) {
	type args struct {
		date string
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			name: "东八区",
			args: args{
				date: "Sun, 11 Jun 2023 20:52:13 +0800",
			},
			want: time.Unix(1686487933, 0),
		},
		{
			name: "非东八区1",
			args: args{
				date: "Sun, 11 Jun 2023 20:52:13 +0700",
			},
			want: time.Unix(1686491533, 0),
		},
		{
			name: "非东八区2",
			args: args{
				date: "Sun, 11 Jun 2023 20:52:13 +0900",
			},
			want: time.Unix(1686484333, 0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dmhy{}
			assert.Equalf(t, tt.want, d.parsePubDate(tt.args.date), "parsePubDate(%v)", tt.args.date)
		})
	}
}

// go test ./internal/app/rss/parse -v -run TestDmhy_parseUk$ --count=1
func TestDmhy_parseUk(t *testing.T) {
	tests := []struct {
		name string
		args string
		want string
	}{
		{
			name: "全数字",
			args: "123345",
			want: "123345",
		},
		{
			name: "数字开头",
			args: "123345ad",
			want: "123345",
		},
		{
			name: "数字结尾",
			args: "fsd123345",
			want: "123345",
		},
		{
			name: "数字在中间",
			args: "dsag123345gdrfgd",
			want: "123345",
		},
		{
			name: "多段数字",
			args: "ad123345ads25dsf623sdf4623fs",
			want: "123345",
		},
		{
			name: "符号分隔数字",
			args: "&%$^&123345_324523",
			want: "123345",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dmhy{}
			assert.Equalf(t, tt.want, d.parseUk(tt.args), "parseUk(%v)", tt.args)
		})
	}
}
