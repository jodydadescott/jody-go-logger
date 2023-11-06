package logging

const (
	DefaultSamplingInitial    = 100
	DefaultSamplingThereafter = 100
	DefaultLogLevel           = InfoLevel
	DefaultEncoding           = EncodingJSON
)

type LogLevel string

const (
	EmptyLevel LogLevel = ""
	WireLevel           = "wire"
	TraceLevel          = "trace"
	DebugLevel          = "debug"
	InfoLevel           = "info"
	WarnLevel           = "warn"
	ErrorLevel          = "error"
)

type Encoding string

const (
	EmptyEncoding   Encoding = ""
	EncodingJSON             = "json"
	EncodingConsole          = "console"
)
