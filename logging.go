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
}
