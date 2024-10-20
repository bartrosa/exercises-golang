/////////////////////////////////////
// 1 We have a list of dictionaries representing product information 
// (keys: "name", "price", "quantity"). Write a function that generates 
// a dictionary from this list, in which the key is the product name and the 
// value is the total price of this product (price * quantity).
//

package main

import "fmt"

type Product struct {
	Name     string
	Price    float64
	Quantity int
}

func calculateTotalPrices(products []Product) map[string]float64 {
	result := make(map[string]float64)
	for _, product := range products {
		result[product.Name] = product.Price * float64(product.Quantity)
	}
	return result
}

func main() {
	products := []Product{
		{Name: "Apple", Price: 0.5, Quantity: 10},
		{Name: "Banana", Price: 0.3, Quantity: 20},
		{Name: "Orange", Price: 0.6, Quantity: 15},
	}

	result := calculateTotalPrices(products)
	fmt.Println(result)
}
