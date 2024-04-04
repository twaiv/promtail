package promtail

type Level uint8

const (
	Debug Level = 0
	Info  Level = 1
	Warn  Level = 2
	Error Level = 3
	Fatal Level = 4
	Panic Level = 5
)

func (l Level) String() string {
	switch l {
	case Debug:
		return "DEBUG"
	case Info:
		return "INFO"
	case Warn:
		return "WARN"
	case Error:
		return "ERROR"
	case Fatal:
		return "FATAL"
	case Panic:
		return "PANIC"

	}
	return "unknown"
}

type Client interface {
	PushLogEntry(entry *LogEntry)
	Ping() (*PongResponse, error)
	Flush()
	Close()
}

type PongResponse struct {
	IsReady bool
}
