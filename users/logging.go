package users

import (
	"context"
	"github.com/go-kit/kit/log"
	"gokit-poc/models"
	"time"
)

type LoggingMiddleware struct {
	Logger log.Logger
	Next   UserService
}

func (mw LoggingMiddleware) CreateUser(ctx context.Context, req CreateUserRequest) (user models.User, err error) {
	defer func(begin time.Time) {
		_ = mw.Logger.Log(
			"method", "CreateUser",
			"input", req,
			"output", user,
			"err", err,
			"took", time.Since(begin),
		)
	}(time.Now())

	return mw.Next.CreateUser(ctx, req)
}
