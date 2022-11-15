package logr_test

import (
	"context"

	"github.com/go-courier/logr"

	"github.com/go-courier/logr/slog"
)

func ExampleLogger() {
	ctx := logr.WithLogger(context.Background(), slog.Logger(slog.Default()))

	log := logr.FromContext(ctx).WithValues("k", "k")

	log.Debug("test %d", 1)
	log.Info("test %d", 1)
	// Output:
}

func ExampleLogger_Start() {
	ctx := logr.WithLogger(context.Background(), slog.Logger(slog.Default()))

	_, log := logr.Start(ctx, "span", "k", "k")
	defer log.End()

	log.Debug("test %d", 1)
	log.Info("test %d", 1)
	// Output:
}
