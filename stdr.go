package logr

import (
	"context"
	"fmt"
	"log"
)

func StdLogger() Logger {
	return &stdLogger{}
}

type stdLogger struct {
	spans        []string
	keyAndValues []interface{}
}

func (d *stdLogger) WithValues(keyAndValues ...interface{}) Logger {
	return &stdLogger{spans: d.spans, keyAndValues: append(d.keyAndValues, keyAndValues...)}
}

func (d *stdLogger) Start(ctx context.Context, name string, keyAndValues ...interface{}) (context.Context, Logger) {
	return ctx, &stdLogger{spans: append(d.spans, name), keyAndValues: append(d.keyAndValues, keyAndValues...)}
}

func (d *stdLogger) End() {
	if len(d.spans) != 0 {
		d.spans = d.spans[0 : len(d.spans)-1]
	}
}

func (d *stdLogger) Trace(format string, args ...interface{}) {
	log.Println(append(keyValues(append(d.keyAndValues, "level", "trace")), fmt.Sprintf(format, args...))...)
}

func (d *stdLogger) Debug(format string, args ...interface{}) {
	log.Println(append(keyValues(append(d.keyAndValues, "level", "debug")), fmt.Sprintf(format, args...))...)
}

func (d *stdLogger) Info(format string, args ...interface{}) {
	log.Println(append(keyValues(append(d.keyAndValues, "level", "info")), fmt.Sprintf(format, args...))...)
}

func (d *stdLogger) Warn(err error) {
	log.Println(append(keyValues(append(d.keyAndValues, "level", "warn")), fmt.Sprintf("%v", err))...)
}

func (d *stdLogger) Error(err error) {
	log.Println(append(keyValues(append(d.keyAndValues, "level", "error")), fmt.Sprintf("%+v", err))...)
}

func (stdLogger) Fatal(err error) {
	log.Fatal(err)
}

func (stdLogger) Panic(err error) {
	log.Panic(err)
}

func keyValues(keyAndValues []interface{}) (values []interface{}) {
	if len(keyAndValues)%2 != 0 {
		return
	}

	for i := 0; i < len(keyAndValues); i += 2 {
		values = append(values, fmt.Sprintf("%v=%v", keyAndValues[i], keyAndValues[i+1]))
	}

	return
}
