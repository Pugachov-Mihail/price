package lib

import (
	"math/rand"
	"strconv"
)

type DataProduct struct {
	Count int32
	Mdate string
	Cdate string
}

func NewDataProduct() *DataProduct {
	return &DataProduct{
		Count: rand.Int31(),
		Mdate: strconv.Itoa(rand.Int()),
		Cdate: strconv.Itoa(rand.Int()),
	}
}
