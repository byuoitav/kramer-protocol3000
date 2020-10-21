package protocol3000

import (
	"time"

	"go.uber.org/zap"
)

var (
	_defaultTTL   = 30 * time.Second
	_defaultDelay = 500 * time.Millisecond
)

type options struct {
	ttl    time.Duration
	delay  time.Duration
	logger *zap.Logger
}

type Option interface {
	apply(*options)
}

type optionFunc func(*options)

func (f optionFunc) apply(o *options) {
	f(o)
}

func WithLogger(l *zap.Logger) Option {
	return optionFunc(func(o *options) {
		o.logger = l
	})
}
