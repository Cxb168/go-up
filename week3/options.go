package main

import (
	"context"
	"log"
	"os"
)

type Option func(o *options)

type options struct {
	id      string
	name    string
	ctx     context.Context
	logger  *log.Logger
	servers []Server
	signals []os.Signal
}

func WithId(id string) Option {
	return func(o *options) {
		o.id = id
	}
}

func WithName(name string) Option {
	return func(o *options) {
		o.name = name
	}
}

func WithSignals(sig ...os.Signal) Option {
	return func(o *options) {
		o.signals = sig
	}
}

func WithServer(servers ...Server) Option {
	return func(o *options) {
		o.servers = servers
	}
}

func WithContext(ctx context.Context) Option {
	return func(o *options) {
		o.ctx = ctx
	}
}
