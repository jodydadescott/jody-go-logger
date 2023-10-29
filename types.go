package logging

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
