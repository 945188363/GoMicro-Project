cd Models/Protos
protoc --micro_out=../ --go_out=../ ProdService
cd.. && cd..