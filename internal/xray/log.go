package xray

// LogConfig represents the logging configuration for the Xray application.
type LogConfig struct {
	LogLevel string `json:"loglevel"` // e.g., "debug", "info", "warning", "error"
}

// NewLogConfig creates a new LogConfig with the specified log level.
func NewLogConfig(level string) *LogConfig {
	return &LogConfig{LogLevel: level}
}
