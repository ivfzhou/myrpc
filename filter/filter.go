package filter

import "context"

type Filter interface {
	Next(ctx context.Context, req, rsp interface{})
}
