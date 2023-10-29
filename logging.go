package logging

import "go.uber.org/zap"

var logLevel LogLevel
var zapLogger *zap.Logger

// New returns logger
func SetConfig(config *Config) {

	if config == nil {
		panic("config is nil")
	}

	logLevel = config.LogLevel

	if logLevel == Empty {
		logLevel = DefaultLogLevel
	}

	if config.SamplingInitial <= 0 {
		config.SamplingInitial = DefaultSamplingInitial
	}

	if config.SamplingThereafter <= 0 {
		config.SamplingThereafter = DefaultSamplingThereafter
	}

	if config.Encoding == "" {
		config.Encoding = DefaultEncoding
	}

	zapConfig := &zap.Config{
		Development: config.Development,
		Level:       zap.NewAtomicLevelAt(zap.DebugLevel),
		Sampling: &zap.SamplingConfig{
			Initial:    config.SamplingInitial,
			Thereafter: config.SamplingThereafter,
		},
		Encoding:      config.Encoding,
		EncoderConfig: zap.NewProductionEncoderConfig(),
	}

	zapConfig.OutputPaths = append(zapConfig.OutputPaths, "stderr")
	zapConfig.ErrorOutputPaths = append(zapConfig.ErrorOutputPaths, "stderr")

	t, err := zapConfig.Build()
	if err != nil {
		panic(err)
	}

	zap.ReplaceGlobals(t)
	zapLogger = t
}

func baseset() {

	if zapLogger != nil {
		return
	}

	logLevel = DefaultLogLevel

	zapConfig := &zap.Config{
		Development: false,
		Level:       zap.NewAtomicLevelAt(zap.DebugLevel),
		Sampling: &zap.SamplingConfig{
			Initial:    DefaultSamplingInitial,
			Thereafter: DefaultSamplingThereafter,
		},
		Encoding:      DefaultEncoding,
		EncoderConfig: zap.NewProductionEncoderConfig(),
	}

	zapConfig.OutputPaths = append(zapConfig.OutputPaths, "stderr")
	zapConfig.ErrorOutputPaths = append(zapConfig.ErrorOutputPaths, "stderr")

	t, err := zapConfig.Build()
	if err != nil {
		panic(err)
	}

	zap.ReplaceGlobals(zapLogger)
	zapLogger = t
}

func Trace(msg string, fields ...zap.Field) {
	baseset()
	switch logLevel {
	case TraceLevel:
		zap.L().Debug(msg, fields...)
	}
}

func Debug(msg string, fields ...zap.Field) {
	baseset()
	switch logLevel {
	case TraceLevel, DebugLevel:
		zap.L().Debug(msg, fields...)
	}
}

func Info(msg string, fields ...zap.Field) {
	baseset()
	switch logLevel {
	case TraceLevel, DebugLevel, InfoLevel:
		zap.L().Info(msg, fields...)
	}
}

func Warn(msg string, fields ...zap.Field) {
	baseset()
	switch logLevel {
	case TraceLevel, DebugLevel, InfoLevel, WarnLevel:
		zap.L().Warn(msg, fields...)
	}
}

func Error(msg string, fields ...zap.Field) {
	baseset()
	switch logLevel {
	case TraceLevel, DebugLevel, InfoLevel, WarnLevel, ErrorLevel:
		zap.L().Error(msg, fields...)
	}
}
