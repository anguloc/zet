package worker

import (
	"context"
	"testing"

	"github.com/anguloc/zet/internal/dao"
	"github.com/anguloc/zet/internal/dao/zet_query"
	"github.com/golang/mock/gomock"
)

func TestMain(m *testing.M) {
	//if err := application.New().Init(context.TODO(), ""); err != nil {
	//	panic(err)
	//}
	m.Run()
}

type Foo interface {
	Bar
	A()
}

type Bar interface {
	a()
	B()
}

type Obj struct {
}

func (o *Obj) A() {

}

func (o *Obj) B() {

}

//func (o *Obj) a() {
//
//}

// go test ./internal/app/worker/ -v -run TestDmhyParse_Run$ --count=1
func TestDmhyParse_Run(t *testing.T) {
	ctx := context.TODO()

	ctrl := gomock.NewController(t)
	_ = ctrl

	//var aa zet_query.IRssDo = zet_query.NewMockIRssDo(ctrl)
	//_ = aa

	println(123)

	return

	type fields struct {
		requestTable zet_query.IRequestDo
		rssTable     zet_query.IRssDo
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
		{
			name: "正常解析",
			fields: fields{
				requestTable: dao.Zet().Request.WithContext(nil),
				rssTable:     nil,
			},
			args: args{
				ctx: ctx,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &DmhyParse{
				requestTable: tt.fields.requestTable,
				rssTable:     tt.fields.rssTable,
			}
			if err := w.Run(tt.args.ctx); (err != nil) != tt.wantErr {
				t.Errorf("Run() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
