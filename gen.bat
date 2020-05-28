cd Service/Protos
protoc --micro_out=../ --go_out=../ Rpc.proto
cd.. && cd..