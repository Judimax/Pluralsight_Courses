package main

import (
	"fmt"
)

func main() {

	itemPrice := 50
	itemDiscount = ItemDiscount
	//itemDiscount = defaults.randomNumber()
	//itemDiscount = defaults.RandomNumber()

	totalDiscount := Discount(itemPrice, itemDiscount)
	fmt.Println("Item Price:      ", itemPrice)
	fmt.Println("Item Discount %: ", itemDiscount)
	fmt.Println("Total Discount:  ", totalDiscount)
}
