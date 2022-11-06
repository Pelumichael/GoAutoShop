package main

import (
	"fmt"
)

// Car declare a car struct
type Car struct {
	name            string
	color           string
	description     string
	manufactureYear int
	model           string
	price           float64
	quantityInStock int
}
type product struct {
	car                []Car
	Name               string
	quantityAvailable  int
	Price              int
	Model              string
	productDescription string
}

type store struct {
	products     []product
	productsSold []product
	allProduct   int
	totalSold    int
	salesMade    int
	makeEntry    string
}

func (c *Car) Name() string {
	return c.name
}

func (c *Car) Description() string {
	return c.description
}

func (c *Car) Model() string {
	return c.model
}

func (c *Car) Color() string {
	return c.color
}

func (c *Car) Price() float64 {
	return c.price
}

func (c *Car) Quantity() int {
	return c.quantityInStock
}

func (c *Car) Year() int {
	return c.manufactureYear
}

func (p product) showProduct() {
	fmt.Println(p)
}

func (p product) showStatusOfProduct() {
	if p.quantityAvailable > 0 {
		fmt.Println("In stock")
	} else {
		fmt.Println("Out of Stock")
	}
}

func (s store) allProducts() {
	fmt.Println("List of all Products: ", s.products)
}

func (s store) display() {
	fmt.Println("Number of Products: ", s.allProduct)
	fmt.Println("List of all Products: ", s.products)
	fmt.Println("Sold Products + (total amount): ", s.totalSold)
}

func (s store) addProduct(p product) {
	s.products = append(s.products, p)
	s.allProduct++
}

func (s store) sellProduct(p product) {
	s.productsSold = append(s.productsSold, p)
	s.salesMade++
	s.totalSold += p.Price
}

// Products Sold. Declare a method for the store to show all products sold at the store.
func (s store) showSoldProducts() {
	for _, p := range s.productsSold {
		fmt.Println(p)
	}
}

// Total Sales. Declare a method for the store struct to show the total sales made by the store.
func (s store) totalSalesMade() {
	fmt.Println(s.salesMade)
}

func main() {
	carr := Car{name: "Honda", quantityInStock: 3, price: 500000, model: "accord"}
	car2 := Car{name: "Toyota", quantityInStock: 4, price: 100000, model: "camry"}
	car3 := Car{name: "Mercedes Benz", quantityInStock: 3, price: 1000000, model: "c-class"}

	prod := product{
		car: []Car{carr, car2, car3},
		// prod1: 	   			car {"carr", "car2", "car3"}

	}
	//Create a store
	store := store{
		products:     []product{prod},
		productsSold: []product{prod},
		allProduct:   10,
		totalSold:    2,
		salesMade:    8,
		makeEntry:    "Add item",
	}

	store.addProduct(prod)
	//Create a product

	//Add the product to the store

	//Sell the product
	store.sellProduct(prod)

	//List all products in the store
	store.allProducts()

	//List all sold products
	store.showSoldProducts()
}
