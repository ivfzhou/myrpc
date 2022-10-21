package transport

import "context"

type ServerTransport interface {
	Listen(context.Context, ...ListenOption)
}

type ClientTransport interface {
	RoundTrip(context.Context, ...ClientOption)
}
