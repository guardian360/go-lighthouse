package client

// Logger is a leveled, structured logger interface. Implementations should
// treat keysAndValues as alternating key-value pairs (e.g. "method", "GET",
// "status", 502).
type Logger interface {
	Debug(msg string, keysAndValues ...any)
	Info(msg string, keysAndValues ...any)
	Warn(msg string, keysAndValues ...any)
	Error(msg string, keysAndValues ...any)
}
