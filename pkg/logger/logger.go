package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Field struct {
	Key   string
	Value any
}
type Fields []Field

var (
	logging *zap.SugaredLogger
)

func setEnvLevel(env string) (bool, zapcore.Level) {
	if env == "dev" || env == "stg" {
		return true, zap.DebugLevel
	}
	return false, zap.InfoLevel
}

func Setup(env string) {
	isDev, levelSeverity := setEnvLevel(env)

	config := zap.Config{
		Level:            zap.NewAtomicLevelAt(levelSeverity),
		Encoding:         "json",
		OutputPaths:      []string{"stderr"},
		ErrorOutputPaths: []string{"stderr"},
		Development:      isDev,
		DisableCaller:    false,

		EncoderConfig: zapcore.EncoderConfig{
			CallerKey:      "caller",
			LevelKey:       "level",
			EncodeLevel:    zapcore.LowercaseLevelEncoder,
			MessageKey:     "msg",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeTime:     zapcore.ISO8601TimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.FullCallerEncoder,
			FunctionKey:    "func",
			TimeKey:        "time",
		},
	}
	logg, err := config.Build()
	if err != nil {
		panic(err)
	}
	logging = logg.Sugar()
	defer logg.Sync()
}

func NewFields(p ...Field) Fields {
	x := Fields{}

	for i := 0; i < len(p); i++ {
		x.Append(p[i])
	}

	return x
}

func (f *Fields) Append(p Field) {
	*f = append(*f, p)
}

func Any(k string, v any) Field {
	return Field{
		Key:   k,
		Value: v,
	}
}

func EventName(v any) Field {
	return Field{
		Key:   "EventName",
		Value: v,
	}
}

func extract(args ...Field) []interface{} {
	var result []interface{}

	for _, field := range args {
		if field.Value == nil || field.Value == "" {
			continue
		}
		result = append(result, field.Key, field.Value)
	}

	return result
}

func Info(msg string, fl ...Field) {
	if len(fl) == 0 {
		logging.Desugar().WithOptions(zap.AddCallerSkip(1)).Sugar().Info(msg)
		return
	}
	logging.Desugar().WithOptions(zap.AddCallerSkip(1)).Sugar().Infow(msg, extract(fl...)...)
}

func Warn(msg string, fl ...Field) {
	if len(fl) == 0 {
		logging.Desugar().WithOptions(zap.AddCallerSkip(1)).Sugar().Warn(msg)
		return
	}
	logging.Desugar().WithOptions(zap.AddCallerSkip(1)).Sugar().Warnw(msg, extract(fl...)...)
}

func Debug(msg string, fl ...Field) {
	if len(fl) == 0 {
		logging.Desugar().WithOptions(zap.AddCallerSkip(1)).Sugar().Debug(msg)
		return
	}
	logging.Desugar().WithOptions(zap.AddCallerSkip(1)).Sugar().Debugw(msg, extract(fl...)...)
}

func Error(msg string, fl ...Field) {
	if len(fl) == 0 {
		logging.Desugar().WithOptions(zap.AddCallerSkip(1)).Sugar().Error(msg)
		return
	}
	logging.Desugar().WithOptions(zap.AddCallerSkip(1)).Sugar().Errorw(msg, extract(fl...)...)
}

func Fatal(msg string, fl ...Field) {
	if len(fl) == 0 {
		logging.Desugar().WithOptions(zap.AddCallerSkip(1)).Sugar().Error(msg)
		return
	}
	logging.Desugar().WithOptions(zap.AddCallerSkip(1)).Sugar().Fatalw(msg, extract(fl...)...)
}
