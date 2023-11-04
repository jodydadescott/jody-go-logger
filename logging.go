package logging

import "go.uber.org/zap"

// New returns logger
func SetConfig(config *Config) {

	if config == nil {
		panic("config is nil")
	}

	if config.LogLevel == Empty {
		config.LogLevel = DefaultLogLevel
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
		Sampling: &zap.SamplingConfig{
			Initial:    config.SamplingInitial,
			Thereafter: config.SamplingThereafter,
		},
		Encoding:      config.Encoding,
		EncoderConfig: zap.NewProductionEncoderConfig(),
	}

	switch config.LogLevel {

	case Empty:
		zapConfig.Level = zap.NewAtomicLevelAt(zap.InfoLevel)

	case TraceLevel:
		zapConfig.Level = zap.NewAtomicLevelAt(zap.DebugLevel)

	case DebugLevel:
		zapConfig.Level = zap.NewAtomicLevelAt(zap.DebugLevel)

	case InfoLevel:
		zapConfig.Level = zap.NewAtomicLevelAt(zap.InfoLevel)

	case WarnLevel:
		zapConfig.Level = zap.NewAtomicLevelAt(zap.WarnLevel)

	case ErrorLevel:
		zapConfig.Level = zap.NewAtomicLevelAt(zap.ErrorLevel)

	}

	zapConfig.OutputPaths = append(zapConfig.OutputPaths, "stderr")
	zapConfig.ErrorOutputPaths = append(zapConfig.ErrorOutputPaths, "stderr")

	t, err := zapConfig.Build()
	if err != nil {
		panic(err)
	}

	zap.ReplaceGlobals(t)
}
