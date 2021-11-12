package log

import (
	"context"
	"github.com/go-chi/chi/middleware"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// ContextLogger is logger that will use context for additional fields
type ContextLogger struct {
	Log     *zap.Logger
	Context context.Context
}

// Debug is zap debug with context
func (l *ContextLogger) Debug(msg string, userFields ...zap.Field) {
	getDefaultLoggerOrPanic().Debug(msg, accumulateLogFields(l.Context, userFields)...)
}

// Info is zap Info with context
func (l *ContextLogger) Info(msg string, userFields ...zap.Field) {
	getDefaultLoggerOrPanic().Info(msg, accumulateLogFields(l.Context, userFields)...)
}

// Error is zap Error with context
func (l *ContextLogger) Error(msg string, userFields ...zap.Field) {
	getDefaultLoggerOrPanic().Error(msg, accumulateLogFields(l.Context, userFields)...)
}

// Panic is zap Panic with context
func (l *ContextLogger) Panic(ctx context.Context, msg string, userFields ...zap.Field) {
	getDefaultLoggerOrPanic().Panic(msg, accumulateLogFields(l.Context, userFields)...)
}

// Warn is zap Warn with context
func (l *ContextLogger) Warn(ctx context.Context, msg string, userFields ...zap.Field) {
	getDefaultLoggerOrPanic().Warn(msg, accumulateLogFields(l.Context, userFields)...)
}

// Check is zap Check with context
func (l *ContextLogger) Check(lvl zapcore.Level, msg string) *zapcore.CheckedEntry {
	return getDefaultLoggerOrPanic().Check(lvl, msg)
}

func accumulateLogFields(ctx context.Context, newFields []zap.Field) []zap.Field {
	previousFields := GetLogFieldsFromCtx(ctx)
	return append(previousFields, newFields...)
}

// GetLogFieldsFromCtx extracts fields from context
func GetLogFieldsFromCtx(ctx context.Context) []zap.Field {
	var fields []zap.Field
	if ctx != nil {
		v := ctx.Value(middleware.RequestIDKey)
		if v != nil {
			fields = append(fields, zap.String("reqID", v.(string)))
		}
	}
	return fields
}
