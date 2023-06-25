package rss

import (
	"testing"

	"github.com/anguloc/zet/internal/app/rss/data"
	"github.com/stretchr/testify/assert"
)

func TestDmhy_filterTitle(t *testing.T) {
	type args struct {
		item *data.Item
	}

	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "health",
			args:    args{},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &Dmhy{}
			err := d.filterTitle(tt.args.item)
			assert.Equal(t, tt.wantErr, err != nil)
		})
	}
}
