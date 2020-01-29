package users

import (
	"context"
	"fmt"
	"github.com/go-kit/kit/metrics"
	"gokit-poc/models"
	"time"
)

type InstrumentingMiddleware struct {
	RequestCount   metrics.Counter
	RequestLatency metrics.Histogram
	CountResult    metrics.Histogram
	Next           UserService
}

func (mw InstrumentingMiddleware) CreateUser(ctx context.Context, req CreateUserRequest) (user models.User, err error) {
	defer func(begin time.Time) {
		lvs := []string{"method", "CreateUser", "error", fmt.Sprint(err != nil)}
		mw.RequestCount.With(lvs...).Add(1)
		mw.RequestLatency.With(lvs...).Observe(time.Since(begin).Seconds())
	}(time.Now())

	return mw.Next.CreateUser(ctx, req)
}
