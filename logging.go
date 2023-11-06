package logging

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Wire = false
var Trace = false

// New returns logger
func SetConfig(config *Config) {

	if config == nil {
		panic("config is nil")
	}

	if config.SamplingInitial <= 0 {
		config.SamplingInitial = DefaultSamplingInitial
	}

	if config.SamplingThereafter <= 0 {
		config.SamplingThereafter = DefaultSamplingThereafter
	}

	if config.Encoding == EmptyEncoding {
		config.Encoding = DefaultEncoding
	}

	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeCaller = zapcore.FullCallerEncoder

	zapConfig := &zap.Config{
		Development: config.Development,
		Sampling: &zap.SamplingConfig{
			Initial:    config.SamplingInitial,
			Thereafter: config.SamplingThereafter,
		},
		Encoding:      string(config.Encoding),
		EncoderConfig: encoderConfig,
	}

	switch config.LogLevel {

	case EmptyLevel:
		zapConfig.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
		Wire = false
		Trace = false

	case WireLevel:
		zapConfig.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
		Wire = true
		Trace = true

	case TraceLevel:
		zapConfig.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
		Wire = false
		Trace = true

	case DebugLevel:
		zapConfig.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
		Wire = false
		Trace = false

	case InfoLevel:
		zapConfig.Level = zap.NewAtomicLevelAt(zap.InfoLevel)
		Wire = false
		Trace = false

	case WarnLevel:
		zapConfig.Level = zap.NewAtomicLevelAt(zap.WarnLevel)
		Wire = false
		Trace = false

	case ErrorLevel:
		zapConfig.Level = zap.NewAtomicLevelAt(zap.ErrorLevel)
		Wire = false
		Trace = false

	}

	zapConfig.OutputPaths = append(zapConfig.OutputPaths, "stderr")
	zapConfig.ErrorOutputPaths = append(zapConfig.ErrorOutputPaths, "stderr")

	t, err := zapConfig.Build()
	if err != nil {
		panic(err)
	}

	zap.ReplaceGlobals(t)
}
