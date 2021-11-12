package log

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"time"
)

// Duplicated constants from zap for more intuitive usage
const (
	// DebugLevel logs are typically voluminous, and are usually disabled in
	// production.
	DebugLevel = zap.DebugLevel
	// InfoLevel is the default logging priority.
	InfoLevel = zap.InfoLevel
	// WarnLevel logs are more important than Info, but don't need individual
	// human review.
	WarnLevel = zap.WarnLevel
	// ErrorLevel logs are high-priority. If an application is running smoothly,
	// it shouldn't generate any error-level logs.
	ErrorLevel = zap.ErrorLevel
	// DPanicLevel logs are particularly important errors. In development the
	// logger panics after writing the message.
	DPanicLevel = zap.DPanicLevel
	// PanicLevel logs a message, then panics.
	PanicLevel = zap.PanicLevel
	// FatalLevel logs a message, then calls os.Exit(1).
	FatalLevel = zap.FatalLevel
)

var log *zap.Logger

var logLevel *zap.AtomicLevel

// SetLevelStr sets level of default logger from level name
// Valid values: debug, info, warn, error, dpanic, panic, fatal
func SetLevelStr(levelStr string) {
	l := getDefaultLoggerOrPanic() // init logger if it hasn't yet been
	err := logLevel.UnmarshalText([]byte(levelStr))
	if err != nil {
		l.Error("can't change log level: invalid string value provided")
		return
	}
}

func getDefaultLoggerOrPanic() *zap.Logger {
	var err error
	if log != nil {
		return log
	}
	// default level: debug
	log, logLevel, err = NewLogger("debug", []string{"stdout"})
	if err != nil {
		panic(err)
	}
	return log
}

// NewLogger creates the logger with defined level. outputs defines the outputs where the
// logs will be sent. By default, outputs contains "stdout", which prints the
// logs at the output of the process. To add a log file as output, the path
// should be added at the outputs array. To avoid printing the logs but storing
// them on a file, can use []string{"pathtofile.log"}
func NewLogger(levelStr string, outputs []string) (*zap.Logger, *zap.AtomicLevel, error) {
	var level zap.AtomicLevel
	err := level.UnmarshalText([]byte(levelStr))
	if err != nil {
		return nil, nil, fmt.Errorf("error on setting log level: %s", err)
	}

	cfg := zap.Config{
		Level:            level,
		Encoding:         "console",
		OutputPaths:      outputs,
		ErrorOutputPaths: []string{"stderr"},
		EncoderConfig: zapcore.EncoderConfig{
			MessageKey:  "message",
			LevelKey:    "level",
			EncodeLevel: zapcore.CapitalColorLevelEncoder,
			TimeKey:     "timestamp",
			EncodeTime: func(ts time.Time, encoder zapcore.PrimitiveArrayEncoder) {
				encoder.AppendString(ts.Local().Format(time.RFC3339))
			},
			EncodeDuration: zapcore.SecondsDurationEncoder,
			CallerKey:      "caller",
			EncodeCaller:   zapcore.ShortCallerEncoder,

			// StacktraceKey: "stacktrace",
			StacktraceKey: "",
			LineEnding:    zapcore.DefaultLineEnding,
		},
	}

	logger, err := cfg.Build()
	if err != nil {
		return nil, nil, err
	}
	defer logger.Sync()
	l := logger.WithOptions(zap.WithCaller(false))
	return l, &level, nil
}

// WithContext creates logger with context
func WithContext(ctx context.Context) *ContextLogger {
	return &ContextLogger{
		Log:     getDefaultLoggerOrPanic(),
		Context: ctx,
	}
}

// Debug is zap debug with context
func Debug(msg string, userFields ...zap.Field) {
	getDefaultLoggerOrPanic().Debug(msg, userFields...)
}

// Info is zap Info with context
func Info(msg string, userFields ...zap.Field) {
	getDefaultLoggerOrPanic().Info(msg, userFields...)
}

// Error is zap Error with context
func Error(msg string, userFields ...zap.Field) {
	getDefaultLoggerOrPanic().Error(msg, userFields...)
}

// Panic is zap Panic with context
func Panic(msg string, userFields ...zap.Field) {
	getDefaultLoggerOrPanic().Panic(msg, userFields...)
}

// Warn is zap Warn with context
func Warn(msg string, userFields ...zap.Field) {
	getDefaultLoggerOrPanic().Warn(msg, userFields...)
}

// Check is zap Check with context
func Check(lvl zapcore.Level, msg string) *zapcore.CheckedEntry {
	return getDefaultLoggerOrPanic().Check(lvl, msg)
}
