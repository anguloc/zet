package application

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"golang.org/x/sync/errgroup"
)

type IApplication interface {
	Init(ctx context.Context, conf string) error
	Run(ctx context.Context) error
	RegisterWorker(name string, worker IWorker)
}

type Application struct {
	initOnce  sync.Once
	stopOnce  sync.Once
	startOnce sync.Once

	mu *sync.Mutex
	wg *sync.WaitGroup

	workers      map[string]IWorker
	workerCancel map[string]context.CancelFunc
}

func New() IApplication {
	return &Application{
		mu:           &sync.Mutex{},
		wg:           &sync.WaitGroup{},
		workers:      make(map[string]IWorker),
		workerCancel: make(map[string]context.CancelFunc),
	}
}

func (app *Application) Init(ctx context.Context, conf string) (err error) {
	app.initOnce.Do(func() {
		if err = app.initConf(ctx, conf); err != nil {
			return
		}

		if err = app.initWorker(ctx); err != nil {
			return
		}
	})
	return
}

func (app *Application) initConf(_ context.Context, conf string) error {
	return nil
}

func (app *Application) initWorker(ctx context.Context) error {
	if len(app.workers) == 0 {
		return fmt.Errorf("failed to missing worker of run")
	}
	eg, _ := errgroup.WithContext(ctx)
	for _, worker := range app.workers {
		w := worker
		eg.Go(func() error {
			return w.Init(ctx)
		})
	}
	return eg.Wait()
}

func (app *Application) Run(ctx context.Context) (err error) {
	app.startOnce.Do(func() {
		if err = app.startWorker(ctx); err != nil {
			return
		}

		if err = app.wait(ctx); err != nil {
			return
		}
	})
	return
}

func (app *Application) startWorker(ctx context.Context) error {
	app.mu.Lock()
	defer app.mu.Unlock()
	app.wg.Add(len(app.workers))
	for name, worker := range app.workers {
		w := worker
		cCtx, cancel := context.WithCancel(ctx)
		app.workerCancel[name] = cancel
		Go(cCtx, func(ctx context.Context) {
			if err := w.Start(ctx); err != nil {
				// TODO log
			}
		}, nil)
	}
	return nil
}

func (app *Application) stopWorker(ctx context.Context, graceful bool) error {
	app.stopOnce.Do(func() {
		for name, worker := range app.workers {
			w := worker
			cancelFunc := app.workerCancel[name]
			Go(ctx, func(ctx context.Context) {
				defer app.wg.Done()
				cancelFunc()
				if graceful {
					_ = w.GracefulStop(ctx)
				} else {
					_ = w.Stop(ctx)
				}
			}, nil)
		}
	})
	return nil
}

func (app *Application) wait(ctx context.Context) error {
	sig := make(chan os.Signal, 2)
	signal.Notify(sig, []os.Signal{syscall.SIGQUIT, syscall.SIGINT, syscall.SIGTERM}...)

	go func() {
		s := <-sig
		go func() {
			_ = app.stopWorker(ctx, s != syscall.SIGQUIT)
		}()
		<-sig
		os.Exit(128 + int(s.(syscall.Signal)))
	}()

	app.wg.Wait()
	fmt.Println("success quit")
	return nil
}

func (app *Application) RegisterWorker(name string, worker IWorker) {
	app.mu.Lock()
	defer app.mu.Unlock()
	app.workers[name] = worker
}

func Go(ctx context.Context, handle func(ctx context.Context), rh func(r interface{})) {
	p := func() {
		if r := recover(); r != nil {
			// TODO log
			if rh == nil {
				return
			}
			Go(ctx, func(ctx context.Context) {
				rh(r)
			}, nil)
		}
	}

	go func() {
		defer p()
		handle(ctx)
	}()
}
