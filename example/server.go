package example

import (
	"context"
	"fmt"
)

type BookService struct{}

func (*BookService) Name() string {
	return "BookService"
}

type BuyReq struct {
	Price int
}
type BuyRsp struct {
	Res int
}

func (*BookService) Buy(ctx context.Context, req *BuyReq) (*BuyRsp, error) {
	fmt.Println(req.Price)
	return &BuyRsp{Res: 1}, nil
}
