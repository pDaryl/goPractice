package main

import (
	"fmt"
)


type Product struct{
	ProductName string
	Price float64
	Supplier Supplier
 }

 type Supplier struct{
	SupplierName string 
	Country string
 }

 func createProduct (name string, price float64, supplierName string, country string ) Product {
	product := Product{
		ProductName: name,
		Price:       price,
		Supplier:    Supplier{
			SupplierName: supplierName,
			Country:      country,
		},
	}
	return product
 }

 func printProduct(product Product) {
	fmt.Printf(
		"Product(Name:%s, Price:%f, Supplier: [Name:%s, Country:%s])\n",
		product.ProductName,
		product.Price,
		product.Supplier.SupplierName, 
		product.Supplier.Country, 
	)
 }

func main(){
p := createProduct("Laptop", 99.99, "AmeraTech", "USA")
printProduct(p)
}