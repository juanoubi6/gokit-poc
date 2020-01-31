package authentications

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
	Next   AuthenticationService
}

func (mw LoggingMiddleware) SignUp(ctx context.Context, req SignUpRequest) (account *models.Account, err error) {
	defer func(begin time.Time) {
		input, _ := json.Marshal(req)
		output, _ := json.Marshal(account)
		_ = mw.Logger.Log(
			"Endpoint", "SignUp",
			"Input", input,
			"Output", output,
			"Err", err,
			"Took", time.Since(begin),
		)
	}(time.Now())

	return mw.Next.SignUp(ctx, req)
}

func (mw LoggingMiddleware) Login(ctx context.Context, req LoginRequest) (token string, err error) {
	defer func(begin time.Time) {
		input, _ := json.Marshal(req)
		_ = mw.Logger.Log(
			"Endpoint", "SignUp",
			"Input", input,
			"Output", token,
			"Err", err,
			"Took", time.Since(begin),
		)
	}(time.Now())

	return mw.Next.Login(ctx, req)
}

func LoggingMiddlewareDecorator(svc AuthenticationService) AuthenticationService {
	logger := log.NewLogfmtLogger(os.Stderr)

	return LoggingMiddleware{logger, svc}
}
