package main

import (
	"fmt"
)

func main() {
	generateID := GenerateID()

	for {
		fmt.Println("\n Inventory Management System ....")
		fmt.Println("1️ Add Product")
		fmt.Println("2️ View Product")
		fmt.Println("3️ Update Stock")
		fmt.Println("4️ Delete Product")
		fmt.Println("5️ View All Products")
		fmt.Println("6️ Exit")
		fmt.Print(" Enter your choice: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var name string
			var price float64
			var stock int
			var category int

			fmt.Print("Enter Product Name: ")
			fmt.Scan(&name)
			fmt.Print("Enter Product Price: ")
			fmt.Scan(&price)
			fmt.Print("Enter Stock Quantity: ")
			fmt.Scan(&stock)
			fmt.Print("Enter Category (0: Electronics, 1: Groseries, 2: Clothing): ")
			fmt.Scan(&category)

			id := generateID()
			AdProduct(id, name, price, stock, Category(category))

		case 2:
			var id string
			fmt.Print("Enter Product ID: ")
			fmt.Scan(&id)

			product, err := GetProduct(id)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Printf(" ID: %s | Name: %s | Price: ₹%.2f | Stock: %d | Category: %d\n",
					product.ID, product.Name, product.Price, product.Stock, product.Category)
			}

		case 3:
			var id string
			var quantity int
			fmt.Print("Enter Product ID: ")
			fmt.Scan(&id)
			fmt.Print("Enter Stock Quantity to Add: ")
			fmt.Scan(&quantity)

			err := UpdateStock(id, quantity)
			if err != nil {
				fmt.Println(err)
			}

		case 4:
			var id string
			fmt.Print("Enter Product ID: ")
			fmt.Scan(&id)

			err := DeleteProduct(id)
			if err != nil {
				fmt.Println(err)
			}

		case 5:
			DisplayInventory()

		case 6:
			fmt.Println(" Exiting")
			return

		default:
			fmt.Println(" Wrong choise Please try again.")
		}
	}
}
