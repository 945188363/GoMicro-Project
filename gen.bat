cd Models/Protos
protoc --micro_out=../ --go_out=../ Prods.proto
cd.. && cd..