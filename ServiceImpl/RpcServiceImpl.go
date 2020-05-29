package ServiceImpl

import (
	db "GoMicro-Project/DB"
	"GoMicro-Project/DBModels"
	"GoMicro-Project/Service/Model"
	"GoMicro-Project/Utils"
	"context"
	"errors"
	"fmt"
	"log"
	"time"
)

type RpcService struct{}

func (*RpcService) GetProdListRpc(ctx context.Context, req *Model.RpcRequest, resp *Model.RpcResponse) error {
	// 测试gorm插入数据
	prod := DBModels.ProdDB{Name: "test"}
	conn := db.DBConn()
	if err := conn.Create(&prod).Error; err != nil {
		log.Fatal(err)
		return err
	}
	time.Sleep(time.Second * 1)
	if req.Request == nil {
		return errors.New("request is empty")
	}
	fmt.Println(req.Request)
	var ret int
	// 反序列化请求参数map
	// paramMap :=Utils.BytesToMap(req.Request)
	// if len(paramMap) == 0 {
	// 	return errors.New("param is empty")
	// }
	// inde := paramMap["Size"]
	inde := Utils.BytesToInt(req.Request)

	for i := 0; i < inde; i++ {
		ret = i * i
	}
	// 序列化响应数据map
	dataMap := make(map[string]interface{})
	dataMap["Num"] = ret
	resp.Response = Utils.MapToBytes(dataMap)
	// resp.Response = Utils.IntToBytes(ret)
	return nil
}
