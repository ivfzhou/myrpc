package example_test

import (
	"testing"

	"github.com/ivfzhou/myrpc/example"
	"github.com/ivfzhou/myrpc/server"
)

func TestExample(t *testing.T) {
	svr := server.New()
	svr.Register(&example.BookService{})
}
