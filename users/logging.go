package users

import (
	"context"
	"encoding/json"
	"github.com/go-kit/kit/log"
	"gokit-poc/models"
	"os"
	"time"
)

type LoggingMiddleware struct {
	Logger log.Logger
	Next   UserService
}

func (mw LoggingMiddleware) CreateUser(ctx context.Context, req CreateUserRequest) (user *models.User, err error) {
	defer func(begin time.Time) {
		input, _ := json.Marshal(req)
		output, _ := json.Marshal(user)
		_ = mw.Logger.Log(
			"Endpoint", "CreateUser",
			"Input", input,
			"Output", output,
			"Err", err,
			"Took", time.Since(begin),
		)
	}(time.Now())

	return mw.Next.CreateUser(ctx, req)
}

func LoggingMiddlewareDecorator(svc UserService) UserService {
	logger := log.NewLogfmtLogger(os.Stderr)

	return LoggingMiddleware{logger, svc}
}
