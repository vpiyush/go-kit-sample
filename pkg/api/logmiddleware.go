package api

import (
	"time"

	"github.com/go-kit/kit/log"
	"github.com/vpiyush/go-kit-sample/pkg/services"
	"github.com/vpiyush/go-kit-sample/pkg/storage/inmem"
)

type loggingMiddleware struct {
	logger log.Logger
	next   services.PairService
}

func NewLoggingMiddleware(l log.Logger, s services.PairService) *loggingMiddleware {
	return &loggingMiddleware{l, s}
}
func (mw loggingMiddleware) Insert(key, value string) (pair *inmem.Pair, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "Insert",
			"input", key, value,
			"output", pair,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	pair, err = mw.next.Insert(key, value)
	return
}

func (mw loggingMiddleware) Get(key string) (pair *inmem.Pair, err error) {
	defer func(begin time.Time) {
		mw.logger.Log(
			"method", "Get",
			"input", key,
			"output", pair,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())
	pair, err = mw.next.Get(key)
	return
}
