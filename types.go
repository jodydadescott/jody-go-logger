package logging

import (
	"fmt"
	"strings"

	"github.com/jinzhu/copier"
)

// LogFail should we log or should we crash
type LogLevel string

const (
	Empty      LogLevel = ""
	TraceLevel          = "trace"
	DebugLevel          = "debug"
	InfoLevel           = "info"
	WarnLevel           = "warn"
	ErrorLevel          = "error"
)

type Config struct {
	LogLevel           LogLevel `json:"logLevel,omitempty" yaml:"logLevel,omitempty"`
	SamplingInitial    int      `json:"samplingInitial,omitempty" yaml:"samplingInitial,omitempty"`
	SamplingThereafter int      `json:"samplingThereafter,omitempty" yaml:"samplingThereafter,omitempty"`
	Development        bool     `json:"development,omitempty" yaml:"development,omitempty"`
	Encoding           string   `json:"encoding,omitempty" yaml:"encoding,omitempty"`
}

// Clone return copy
func (t *Config) Clone() *Config {
	c := &Config{}
	copier.Copy(&c, &t)
	return c
}

func (t *Config) ParseLogLevel(s string) {

	switch strings.ToLower(s) {

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

	default:
		panic(fmt.Sprintf("LogLevel %s is invalid", s))

	}

}
