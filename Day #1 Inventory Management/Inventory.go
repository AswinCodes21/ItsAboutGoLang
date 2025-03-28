package main

import (
	"errors"
	"fmt"
)

type Category int

const (
	Electronics Category = iota
	Groceries
	Clothing
)

type Product struct {
	ID       string
	Name     string
	Price    float64
	Stock    int
	Category Category
}

var inventory = make(map[string]*Product)

func AdProduct(id, name string, price float64, stock int, category Category) {
	inventory[id] = &Product{id, name, price, stock, category}
	fmt.Println("Product added successFully!")
}

func GetProduct(id string) (*Product, error) {
	product, exists := inventory[id]
	if !exists {
		return nil, errors.New("Product not found")
	}
	return product, nil
}
func UpdateStock(id string, quantity int) error {
	product, exists := inventory[id]
	if !exists {
		return errors.New("Product not found")
	}

	product.Stock += quantity
	inventory[id] = product
	fmt.Println(" Stock updated successfully!")
	return nil
}

func DeleteProduct(id string) error {
	if _, exists := inventory[id]; !exists {
		return errors.New("Product not found")
	}
	delete(inventory, id)
	fmt.Println("Product deleted successfully!")
	return nil
}
func DisplayInventory() {
	if len(inventory) == 0 {
		fmt.Println("Inventory is empty.")
		return
	}

	fmt.Println("\nInventory List:")
	for _, product := range inventory {
		fmt.Println("-------------------------------------------------")
		fmt.Println("ID       :", product.ID)
		fmt.Println("Name     :", product.Name)
		fmt.Println("Price    : ₹", product.Price)
		fmt.Println("Stock    :", product.Stock)
		fmt.Println("Category :", product.Category)
	}
	fmt.Println("-------------------------------------------------")
}
