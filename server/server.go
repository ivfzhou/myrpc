package server

import "context"

type Handler[Req any, Rsp any] interface {
	Name() string
	Handle(ctx context.Context, req Req) (rsp Rsp, err error)
}

type Server interface {
	Register(svr interface{}) error
	ListenAndServe() error
	Close(chan struct{}) error
}

type server struct{}

func New() Server {
	return &server{}
}

func (*server) Register(svr interface{}) error {
	return nil
}

func (*server) ListenAndServe() error {
	return nil
}

func (*server) Close(chan struct{}) error {
	return nil
}
