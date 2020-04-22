package ServiceImpl

import (
	"GoMicro-Project/Service/Model"
	"context"
	"strconv"
	"time"
)

type ProdService1 struct{}

func NewProd(id int32, name string) *Model.ProdModel {
	return &Model.ProdModel{
		ProdId:   id,
		ProdName: name,
	}
}

func (*ProdService1) GetProdList(ctx context.Context, req *Model.ProdRequest1, res *Model.ProdResponse1) error {
	time.Sleep(time.Second * 4)
	ret := make([]*Model.ProdModel, 0)
	var i int32
	for i = 0; i < req.Size; i++ {
		ret = append(ret, NewProd(100+i, "prod"+strconv.Itoa(int(i))))
	}
	res.Data = ret
	return nil
}
