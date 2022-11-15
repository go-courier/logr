package slog

import (
	"context"
	"fmt"
	"strings"

	"github.com/go-courier/logr"
	"golang.org/x/exp/slog"
)

func Logger(l *slog.Logger) logr.Logger {
	return &logger{slog: l}
}

type logger struct {
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
	if !d.slog.Enabled(slog.DebugLevel) {
		return
	}
	d.slog.LogDepth(0, slog.DebugLevel, fmt.Sprintf(format, args...))
}

func (d *logger) Info(format string, args ...any) {
	if !d.slog.Enabled(slog.InfoLevel) {
		return
	}
	d.slog.LogDepth(0, slog.InfoLevel, fmt.Sprintf(format, args...))
}

func (d *logger) Warn(err error) {
	if !d.slog.Enabled(slog.WarnLevel) {
		return
	}
	d.slog.LogDepth(0, slog.WarnLevel, err.Error(), slog.Any("err", err))
}

func (d *logger) Error(err error) {
	d.slog.Error(err.Error(), err)
}
