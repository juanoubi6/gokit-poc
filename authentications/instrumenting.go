package authentications

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
	Next           AuthenticationService
}

func (mw InstrumentingMiddleware) SignUp(ctx context.Context, req SignUpRequest) (account *models.Account, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "SignUp", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	return mw.Next.SignUp(ctx, req)
}

func (mw InstrumentingMiddleware) Login(ctx context.Context, req LoginRequest) (token string, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "Login", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	return mw.Next.Login(ctx, req)
}

func InstrumentingMiddlewareDecorator(svc AuthenticationService) AuthenticationService {
	fieldKeys := []string{"method", "error"}
	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "API",
		Subsystem: "AuthenticationService",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)
	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "API",
		Subsystem: "AuthenticationService",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys)
	countResult := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "API",
		Subsystem: "AuthenticationService",
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
