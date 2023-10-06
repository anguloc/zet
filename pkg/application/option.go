package application

type OptionFunc func(opt *option)

type option struct {
	conf string
}

func WithOptionConfig(conf string) OptionFunc {
	if conf == "" {
		conf = "conf/config.yml"
	}
	return func(opt *option) {
		opt.conf = conf
	}
}
