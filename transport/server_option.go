package transport

type ListenOptions struct{}

type ListenOption func(options *ListenOptions)
