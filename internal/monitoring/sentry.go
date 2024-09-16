package monitoring

import (
	"fmt"

	"github.com/getsentry/sentry-go"
	sentryhttp "github.com/getsentry/sentry-go/http"
	"github.com/sylphritz/go-microservice-template/internal/config"
)

func InitSentry() *sentryhttp.Handler {
	if err := sentry.Init(sentry.ClientOptions{
		Dsn: config.C.SentryUrl,
		// TODO: Adjust sample rate based on environment
		TracesSampleRate: 1.0,
	}); err != nil {
		fmt.Printf("Sentry initialization failed: %v\n", err)
	}

	return sentryhttp.New(sentryhttp.Options{
		// Pass down the panic in case other services/libraries need the errors as well
		Repanic: true,
	})
}
