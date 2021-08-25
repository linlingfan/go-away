package main

import (
	"bytes"
	"fmt"
	"io"
)

func main() {
	var w io.Writer = new(bytes.Buffer)
	w.Write([]byte{'a', 97})
	fmt.Printf("%+v \n", w)
	buf := bytes.NewBuffer([]byte{'a', '1'})

	buf.Write([]byte{'1', 65})
	fmt.Printf("%+v \n", buf.String())

	//
	baoma := new(BaoMa)
	var carPrice CarPrice = baoma
	var carColor CarColor = baoma
	var priceAndColor PriceAndColor = baoma
	carPrice.Price()
	carColor.Color()
	fmt.Println(baoma)
	fmt.Println(carPrice)
	fmt.Printf("value: %+v,and type: %T \n", carPrice, carColor)
	priceAndColor.Price()
	priceAndColor.Color()

	fmt.Printf("value: %+v,and type: %T \n", baoma, baoma)
	println(baoma.Price())
	println(baoma.Color())
}

type CarPrice interface {
	Price() int
}

type CarColor interface {
	Color() string
}

type PriceAndColor interface {
	CarPrice
	CarColor
}

type BaoMa struct {
	ID int
}

func (b *BaoMa) Price() int {
	return 100
}

func (b *BaoMa) Color() string {
	return "color"
}
