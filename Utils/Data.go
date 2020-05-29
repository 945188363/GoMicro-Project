package Utils

import (
	"GoMicro-Project/Core"
	"bytes"
	"encoding/binary"
	"unsafe"
)

func IntToBytes(n int) []byte {
	data := int64(n)
	bytebuf := bytes.NewBuffer([]byte{})
	binary.Write(bytebuf, binary.BigEndian, data)
	return bytebuf.Bytes()
}

func BytesToInt(bys []byte) int {
	bytebuff := bytes.NewBuffer(bys)
	var data int64
	binary.Read(bytebuff, binary.BigEndian, &data)
	return int(data)
}

// 与[]Byte对应的数据结果
type SliceMock struct {
	addr uintptr
	len  int
	cap  int
}

func MessageToBytes(t Core.Message) []byte {
	// 因[]byte底层数据结构与slicemock一致，需要构造一个完全一样的数据结构进行转换
	len := unsafe.Sizeof(t)
	sliceMockTest := SliceMock{
		addr: uintptr(unsafe.Pointer(&t)),
		len:  int(len),
		cap:  int(len),
	}
	return *(*[]byte)(unsafe.Pointer(&sliceMockTest))
}

func BytesToMessage(t []byte) Core.Message {
	// []byte转换成数据结构，只需取出addr地址即可，然后转换成对应的数据结构类型即可
	return *(*Core.Message)(unsafe.Pointer(&t[0]))
}

func MapToBytes(t map[string]interface{}) []byte {
	// 因[]byte底层数据结构与slicemock一致，需要构造一个完全一样的数据结构进行转换
	len := unsafe.Sizeof(t)
	sliceMockTest := SliceMock{
		addr: uintptr(unsafe.Pointer(&t)),
		len:  int(len),
		cap:  int(len),
	}
	return *(*[]byte)(unsafe.Pointer(&sliceMockTest))
}

func BytesToMap(t []byte) map[string]interface{} {
	// []byte转换成数据结构，只需取出addr地址即可，然后转换成对应的数据结构类型即可
	return *(*map[string]interface{})(unsafe.Pointer(&t[0]))
}
