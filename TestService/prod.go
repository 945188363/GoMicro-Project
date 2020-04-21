package TestService

import (
	"strconv"
)

type Prod struct {
	ProdId   int
	ProdName string
}

func NewProd(id int, name string) *Prod {
	return &Prod{
		ProdId:   id,
		ProdName: name,
	}
}

func NewProdList(n int) []*Prod {
	ret := make([]*Prod, 0)
	for i := 0; i < n; i++ {
		ret = append(ret, NewProd(100+i, "prod"+strconv.Itoa(i)))
	}
	return ret
}
