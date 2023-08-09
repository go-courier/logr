package slog

import (
	"context"
	"fmt"
	"log/slog"
	"strings"

	"github.com/go-courier/logr"
)

func Logger(l *slog.Logger) logr.Logger {
	return &logger{slog: l, ctx: context.Background()}
}

type logger struct {
	ctx   context.Context
	slog  *slog.Logger
	spans []string
}

func (d *logger) WithValues(keyAndValues ...any) logr.Logger {
	return &logger{
		spans: d.spans,
		slog:  d.slog.With(keyAndValues...),
	}
}

func (d *logger) Start(ctx context.Context, name string, keyAndValues ...any) (context.Context, logr.Logger) {
	spans := append(d.spans, name)

	if len(keyAndValues) == 0 {
		return ctx, &logger{
			ctx:   ctx,
			spans: spans,
			slog:  d.slog.WithGroup(strings.Join(spans, "/")),
		}
	}

	return ctx, &logger{
		spans: spans,
		slog:  d.slog.WithGroup(strings.Join(spans, "/")).With(keyAndValues...),
	}
}

func (d *logger) End() {
	if len(d.spans) != 0 {
		d.spans = d.spans[0 : len(d.spans)-1]
	}
}

func (d *logger) Debug(format string, args ...any) {
	if !d.slog.Enabled(d.ctx, slog.LevelDebug) {
		return
	}
	d.slog.Log(d.ctx, slog.LevelDebug, fmt.Sprintf(format, args...))
}

func (d *logger) Info(format string, args ...any) {
	if !d.slog.Enabled(d.ctx, slog.LevelInfo) {
		return
	}
	d.slog.Log(d.ctx, slog.LevelInfo, fmt.Sprintf(format, args...))
}

func (d *logger) Warn(err error) {
	if !d.slog.Enabled(d.ctx, slog.LevelWarn) {
		return
	}
	d.slog.Log(d.ctx, slog.LevelWarn, err.Error(), slog.Any("err", err))
}

func (d *logger) Error(err error) {
	d.slog.Error(err.Error(), err)
}
