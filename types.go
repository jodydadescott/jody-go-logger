package logging

import (
	"fmt"
	"strings"

	"github.com/jinzhu/copier"
)

type Config struct {
	LogLevel           LogLevel `json:"logLevel,omitempty" yaml:"logLevel,omitempty"`
	SamplingInitial    int      `json:"samplingInitial,omitempty" yaml:"samplingInitial,omitempty"`
	SamplingThereafter int      `json:"samplingThereafter,omitempty" yaml:"samplingThereafter,omitempty"`
	Development        bool     `json:"development,omitempty" yaml:"development,omitempty"`
	Encoding           Encoding `json:"encoding,omitempty" yaml:"encoding,omitempty"`
}

// Clone return copy
func (t *Config) Clone() *Config {
	c := &Config{}
	copier.Copy(&c, &t)
	return c
}

func (t *Config) ParseLogLevel(s string) error {

	switch strings.ToLower(s) {

	case string(WireLevel):
		t.LogLevel = WireLevel

	case string(TraceLevel):
		t.LogLevel = TraceLevel

	case string(DebugLevel):
		t.LogLevel = DebugLevel

	case string(InfoLevel):
		t.LogLevel = InfoLevel

	case string(WarnLevel):
		t.LogLevel = WarnLevel

	case string(ErrorLevel):
		t.LogLevel = ErrorLevel

	default:
		return fmt.Errorf("LogLevel %s is invalid", s)

	}

	return nil
}

func (t *Config) ParseEncoding(s string) error {

	switch strings.ToLower(s) {

	case string(EncodingJSON):
		t.Encoding = EncodingJSON

	case string(EncodingConsole):
		t.Encoding = EncodingConsole

	default:
		return fmt.Errorf("Encoding %s is invalid", s)

	}

	return nil
}
