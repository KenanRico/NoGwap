package main

import "fmt"


type OnSaleItem struct {
	Name string
	Price_orig float32
	Price_sale float32
	Seller string
	Link string
	Img string
}

func (osi OnSaleItem) ToString() string {
	var str string
	str = fmt.Sprintf(
		"Name: %s\nOriginal price: $%f, discount price: $%f, sold by %s\nLink: %s\nImage: %s\n",
		osi.Name, osi.Price_orig, osi.Price_sale, osi.Seller, osi.Link, osi.Img,
	)
	return str
}

type OnSaleItemList []OnSaleItem
func (osis OnSaleItemList) ToString() string {
	var str string	
	for i:=0; i<len(osis); i++ {
		str += osis[i].ToString()+"--\n"
	}
	return str
}