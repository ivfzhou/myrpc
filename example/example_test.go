package example_test

import (
	"testing"

	"github.com/ivfzhou/myrpc/example"
	"github.com/ivfzhou/myrpc/server"
)

func TestExample(t *testing.T) {
	svr := server.New()
	err := svr.Register(&example.BookService{})
	if err != nil {
		t.Fatal(err)
	}
	err = svr.ListenAndServe()
	if err != nil {
		t.Fatal(err)
	}
}
