package main

import "fmt"


type OnSaleItem struct {
	name string
	price_orig float32
	price_sale float32
	seller string
	link string
	img string
}

func (osi OnSaleItem) ToString() string {
	var str string
	str = fmt.Sprintf(
		"Name: %s\nOriginal price: $%f, discount price: $%f, sold by %s\nLink: %s\nImage: %s\n",
		osi.name, osi.price_orig, osi.price_sale, osi.seller, osi.link, osi.img,
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