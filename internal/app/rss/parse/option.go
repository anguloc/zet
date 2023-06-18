package parse

import "github.com/anguloc/zet/internal/dto"

type Option struct {
	module *dto.Module
}

type OptionFunc func(opt *Option)

func WithModule(module *dto.Module) OptionFunc {
	return func(opt *Option) {
		opt.module = module
	}
}
