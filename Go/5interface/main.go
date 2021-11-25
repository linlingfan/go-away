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
	//baoma := new(BaoMa)
	//var carPrice CarPrice = baoma
	//var carColor CarColor = baoma
	//var priceAndColor PriceAndColor = baoma
	//carPrice.Price()
	//carColor.Color()
	//fmt.Println(baoma)
	//fmt.Println(carPrice)
	//fmt.Printf("value: %+v,and type: %T \n", carPrice, carColor)
	//priceAndColor.Price()
	//priceAndColor.Color()
	//
	//fmt.Printf("value: %+v,and type: %T \n", baoma, baoma)
	//println(baoma.Price())
	//println(baoma.Color())

	//
	var car1 Car = new(BMW)
	car1.Name()
	car1.Run()
	var car2 Car = new(BenChi)
	car2.Name()
	car2.Run()
	var f1 Factory = new(BMW)
	f1.FName()
}

type CarPrice interface {
	Price() int
}

type CarColor interface {
	Color() string
}

// 组合接口
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

// test interface 实现接口

type Car interface {
	Run()
	Name()
}

type Factory interface {
	//Name(n string) // 无法重载
	FName()
}

type BMW struct {
	Car
	Factory // 实现两个接口 方法名要不一样
}

func (b *BMW) Run() {
	println("bmw is running")
}

func (b *BMW) Name() {
	println("this is bmw")
}

//func (b *BMW) Name(n string) {
//	println("this is bmw")
//}

func (b *BMW)FName()  {
	println("bmw factory")
}

type BenChi struct {
	Car
}

func (b *BenChi) Run() {
	println("benchi is running")
}

func (b *BenChi) Name() {
	println("this is benchi")
}
