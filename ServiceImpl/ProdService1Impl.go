package ServiceImpl

import (
	db "GoMicro-Project/DB"
	"GoMicro-Project/DBModels"
	"GoMicro-Project/Service/Model"
	"context"
	"log"
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
	// 测试gorm插入数据
	prod := DBModels.ProdDB{Name: "test"}
	conn := db.DBConn()
	if err := conn.Create(&prod).Error; err != nil {
		log.Fatal(err)
		return err
	}
	time.Sleep(time.Second * 4)
	ret := make([]*Model.ProdModel, 0)
	var i int32
	for i = 0; i < req.Size; i++ {
		ret = append(ret, NewProd(100+i, "prod"+strconv.Itoa(int(i))))
	}
	res.Data = ret
	return nil
}
