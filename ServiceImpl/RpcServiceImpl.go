package ServiceImpl

import (
	db "GoMicro-Project/DB"
	"GoMicro-Project/DBModels"
	"GoMicro-Project/Service/Model"
	"GoMicro-Project/Utils"
	"context"
	"errors"
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
	// ret := make([]int, 0)
	var ret int
	// size := Utils.BytesToMessage(req.Request).Data["Size"]
	size := req.Request
	if size == nil {
		return errors.New("empty data")
	}
	inde := Utils.BytesToInt(size)
	for i := 0; i < inde; i++ {
		ret = i * i
	}
	// respData := make(map[string]interface{})
	// respData["Data"] = ret
	// msg := Core.NewMessage(203,"rpc service dial success",respData)
	// resp.Response = Utils.MessageToBytes(msg)
	resp.Response = Utils.IntToBytes(ret)
	return nil
}
