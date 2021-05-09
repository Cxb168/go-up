package main

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"golang.org/x/sync/errgroup"
	"log"
	"os"
	"os/signal"
	"syscall"
)

type Server interface {
	Start() error
	Stop() error
	Name() string
}

type App struct {
	opts   options
	ctx    context.Context
	cancel func()
}

func NewApp(opts ...Option) *App {
	o := options{
		ctx:     context.Background(),
		logger:  log.Default(),
		signals: []os.Signal{syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT},
	}
	if id, err := uuid.NewUUID(); err == nil {
		o.id = id.String()
	}
	for _, opt := range opts {
		opt(&o)
	}
	ctx, cancel := context.WithCancel(o.ctx)
	return &App{
		opts:   o,
		ctx:    ctx,
		cancel: cancel,
	}
}

func (app *App) Run() error {
	log.Printf("app %s run...\n", app.opts.name)
	g, ctx := errgroup.WithContext(app.ctx)
	for _, srv := range app.opts.servers {
		srv := srv
		g.Go(func() error {
			<-ctx.Done()
			return srv.Stop()
		})
		g.Go(func() error {
			return srv.Start()
		})
	}
	c := make(chan os.Signal, 1)
	signal.Notify(c, app.opts.signals...)
	g.Go(func() error {
		for {
			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-c:
				return app.Stop()
			}
		}
	})

	if err := g.Wait(); err != nil && !errors.Is(err, context.Canceled) {
		return err
	}
	return nil
}

func (app *App) Stop() error {
	if app.cancel != nil {
		app.cancel()
	}
	return nil
}
