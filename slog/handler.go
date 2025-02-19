package slog

import (
	"context"
	"log/slog"
)

func EnableLevel(lvl slog.Level) func(h *handler) {
	return func(h *handler) {
		h.lvl = lvl
	}
}

func Default(optFns ...func(h *handler)) *slog.Logger {
	h := &handler{h: slog.Default().Handler()}
	for _, fn := range optFns {
		fn(h)
	}
	return slog.New(h)
}

type handler struct {
	h   slog.Handler
	lvl slog.Level
}

func (h *handler) Handle(ctx context.Context, r slog.Record) error {
	return h.h.Handle(ctx, r)
}

func (h *handler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &handler{h: h.h.WithAttrs(attrs)}
}

func (h *handler) WithGroup(name string) slog.Handler {
	return &handler{h: h.h.WithGroup(name)}
}

func (h *handler) Enabled(ctx context.Context, l slog.Level) bool {
	return l >= h.lvl
}
