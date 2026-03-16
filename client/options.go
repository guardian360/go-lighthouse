package client

// Option configures the client or HTTP client.
type Option func(*options)

type options struct {
	insecure bool
	logger   Logger
}

// WithInsecure disables TLS certificate verification.
func WithInsecure(insecure bool) Option {
	return func(o *options) { o.insecure = insecure }
}

// WithLogger sets the logger for retry logging and other diagnostics.
func WithLogger(logger Logger) Option {
	return func(o *options) { o.logger = logger }
}

func applyOptions(opts []Option) options {
	var o options
	for _, opt := range opts {
		opt(&o)
	}
	return o
}
