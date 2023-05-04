package safe

import "context"

func Go(ctx context.Context, handle func(ctx context.Context), rh func(r interface{})) {
	p := func() {
		if r := recover(); r != nil {
			// TODO log
			if rh == nil {
				return
			}
			Go(ctx, func(_ context.Context) {
				rh(r)
			}, nil)
		}
	}

	go func() {
		defer p()
		handle(ctx)
	}()
}
