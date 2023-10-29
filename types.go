package logging

import (
	"strings"
)

// LogFail should we log or should we crash
type LogLevel string

const (
	DefaultSamplingInitial    = 100
	DefaultSamplingThereafter = 100
	DefaultEncoding           = "console"
	DefaultLogLevel           = InfoLevel
)

const (
	Empty      LogLevel = ""
	Disabled            = "disabled"
	InfoLevel           = "info"
	TraceLevel          = "trace"
	DebugLevel          = "debug"
	WarnLevel           = "warn"
	ErrorLevel          = "error"
)

type Config struct {
	LogLevel           LogLevel
	SamplingInitial    int
	SamplingThereafter int
	Development        bool
	Encoding           string
}

func (t *Config) ParseLogLevel(s string) error {

	switch strings.ToLower(s) {

	case string(Disabled):
		t.LogLevel = Disabled

	case string(InfoLevel):
		t.LogLevel = InfoLevel

	case string(TraceLevel):
		t.LogLevel = TraceLevel

	case string(DebugLevel):
		t.LogLevel = DebugLevel

	case string(WarnLevel):
		t.LogLevel = WarnLevel

	case string(ErrorLevel):
		t.LogLevel = ErrorLevel

	}

	return nil
}
