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
	WireLevel  LogLevel = "wire"
	TraceLevel LogLevel = "trace"
	DebugLevel LogLevel = "debug"
	InfoLevel  LogLevel = "info"
	WarnLevel  LogLevel = "warn"
	ErrorLevel LogLevel = "error"
)

type Encoding string

const (
	EmptyEncoding   Encoding = ""
	EncodingJSON    Encoding = "json"
	EncodingConsole Encoding = "console"
)
