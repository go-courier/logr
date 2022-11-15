package slog

import (
	"golang.org/x/exp/slog"
)

func Default() *slog.Logger {
	return slog.New(&handler{h: slog.Default().Handler()})
}

type handler struct {
	h slog.Handler
}

func (h *handler) Handle(r slog.Record) error {
	return h.h.Handle(r)
}

func (h *handler) WithAttrs(attrs []slog.Attr) slog.Handler {
	return &handler{h: h.h.WithAttrs(attrs)}
}

func (h *handler) WithGroup(name string) slog.Handler {
	return &handler{h: h.h.WithGroup(name)}
}

func (h *handler) Enabled(l slog.Level) bool {
	return l >= slog.DebugLevel
}
