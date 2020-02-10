package users

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/metrics"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
	"gokit-poc/models"
	"time"
)

type InstrumentingMiddleware struct {
	RequestCount   metrics.Counter
	RequestLatency metrics.Histogram
	CountResult    metrics.Histogram
	Next           UserService
}

func (mw InstrumentingMiddleware) CreateUser(ctx context.Context, req CreateUserRequest) (user *models.User, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "CreateUser", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	return mw.Next.CreateUser(ctx, req)
}

func (mw InstrumentingMiddleware) GetUsers(ctx context.Context, req GetUsersRequest) (users []*models.User, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "GetUsers", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	return mw.Next.GetUsers(ctx, req)
}

func InstrumentingMiddlewareDecorator(svc UserService) UserService {
	fieldKeys := []string{"method", "error"}
	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "API",
		Subsystem: "UserService",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)
	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "API",
		Subsystem: "UserService",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys)
	countResult := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "API",
		Subsystem: "UserService",
		Name:      "count_result",
		Help:      "The result of each count method.",
	}, []string{})

	return InstrumentingMiddleware{
		requestCount,
		requestLatency,
		countResult,
		svc,
	}
}
