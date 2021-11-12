package log

import (
	"context"
	"fmt"
	"github.com/go-chi/chi/middleware"
	"go.uber.org/zap"
)

// ContextLogger is logger that will use context for additional fields
type ContextLogger struct {
	Log     *zap.SugaredLogger
	Context context.Context
}

// Debug calls log.Debug
func (l *ContextLogger) Debug(args ...interface{}) {
	getDefaultLoggerOrPanic().Debugf(accumulateTemplate(l.Context, "%s"), args...)
}

// Info calls log.Info
func (l *ContextLogger) Info(args ...interface{}) {
	getDefaultLoggerOrPanic().Infof(accumulateTemplate(l.Context, "%s"), args...)
}

// Warn calls log.Warn  with context fields
func (l *ContextLogger) Warn(args ...interface{}) {
	args = appendStackTraceMaybeArgs(args)
	getDefaultLoggerOrPanic().Warnf(accumulateTemplate(l.Context, "%s"), args...)
}

// Error calls log.Error  with context fields
func (l *ContextLogger) Error(args ...interface{}) {
	args = appendStackTraceMaybeArgs(args)
	getDefaultLoggerOrPanic().Errorf(accumulateTemplate(l.Context, "%s"), args...)
}

// Fatal calls log.Fatal  with context fields
func (l *ContextLogger) Fatal(args ...interface{}) {
	args = appendStackTraceMaybeArgs(args)
	getDefaultLoggerOrPanic().Fatalf(accumulateTemplate(l.Context, "%s"), args...)
}

// Debugf calls log.Debugf  with context fields
func (l *ContextLogger) Debugf(template string, args ...interface{}) {
	getDefaultLoggerOrPanic().Debugf(accumulateTemplate(l.Context, template), args...)
}

// Infof calls log.Infof  with context fields
func (l *ContextLogger) Infof(template string, args ...interface{}) {
	getDefaultLoggerOrPanic().Infof(accumulateTemplate(l.Context, template), args)
}

// Warnf calls log.Warnf with context fields
func (l *ContextLogger) Warnf(template string, args ...interface{}) {
	getDefaultLoggerOrPanic().Warnf(accumulateTemplate(l.Context, template), args)
}

// Fatalf calls log.Fatalf   with context fields
func (l *ContextLogger) Fatalf(template string, args ...interface{}) {
	getDefaultLoggerOrPanic().Fatalf(accumulateTemplate(l.Context, template), args)
}

// Errorf calls log.Errorf and stores the error message into the ErrorFile
func (l *ContextLogger) Errorf(template string, args ...interface{}) {
	getDefaultLoggerOrPanic().Errorf(accumulateTemplate(l.Context, template), args)
}

// Debugw is zap debug with context
func (l *ContextLogger) Debugw(msg string, kv ...interface{}) {
	getDefaultLoggerOrPanic().Debugw(accumulateTemplate(l.Context, msg), kv...)
}

// Infow is zap Infow with context
func (l *ContextLogger) Infow(msg string, kv ...interface{}) {
	getDefaultLoggerOrPanic().Infow(accumulateTemplate(l.Context, msg), kv...)
}

// Errorw is zap Errorw with context
func (l *ContextLogger) Errorw(msg string, kv ...interface{}) {
	getDefaultLoggerOrPanic().Errorw(accumulateTemplate(l.Context, msg), kv...)
}

// Panicw is zap Panicw with context
func (l *ContextLogger) Panicw(msg string, kv ...interface{}) {
	getDefaultLoggerOrPanic().Panicw(accumulateTemplate(l.Context, msg), kv...)
}

// Warnw is zap Warnw with context
func (l *ContextLogger) Warnw(msg string, kv ...interface{}) {
	getDefaultLoggerOrPanic().Warnw(accumulateTemplate(l.Context, msg), kv...)
}

func accumulateTemplate(ctx context.Context, template string) string {
	requestID := GetRequestIDFromContext(ctx)
	template = fmt.Sprintf("%v\t%s", requestID, template)
	return template
}

// GetRequestIDFromContext extracts requestID from context
func GetRequestIDFromContext(ctx context.Context) string {
	return middleware.GetReqID(ctx)
}
