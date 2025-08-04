package main

import (
	"fmt"
	"math/rand"
	"slices"
	"time"
)

type Product struct {
	Name      string
	Quantity  int
	CreatedAt time.Time
}
type ProductList struct {
	products []Product
	sortFn   SortFunc
}
type SortFunc func(a, b Product) int

func (pl *ProductList) Generate() {
	n := rand.Intn(10) + 1
	products := make([]Product, n)
	for i := range products {
		products[i] = Product{
			Name:      fmt.Sprintf("Product %d", i+1),
			Quantity:  rand.Intn(100),
			CreatedAt: time.Now().Add(time.Duration(-rand.Intn(1000)) * time.Hour),
		}
	}
	pl.products = products
}
func (pl *ProductList) SetSortFn(fn SortFunc) {
	pl.sortFn = fn
}
func (p Product) Print() {
	fmt.Printf("Name: %s, Quantity: %d, CreatedAt: %s\n", p.Name, p.Quantity, p.CreatedAt.Format(time.RFC822Z))
}
func (pl ProductList) Print() {
	for _, product := range pl.products {
		product.Print()
	}
}
func (pl ProductList) Sort() {
	slices.SortFunc(pl.products, pl.sortFn)
}

func SortNameAsc(a, b Product) int {
	if a.Name < b.Name {
		return -1
	}
	return 1
}
func SortNameDesc(a, b Product) int {
	if a.Name > b.Name {
		return -1
	}
	return 1
}
func SortQuantityAsc(a, b Product) int {
	if a.Quantity < b.Quantity {
		return -1
	}
	return 1
}
func SortQuantityDesc(a, b Product) int {
	if a.Quantity > b.Quantity {
		return -1
	}
	return 1
}
func SortCreatedAtAsc(a, b Product) int {
	if a.CreatedAt.Before(b.CreatedAt) {
		return -1
	}
	return 1
}
func SortCreatedAtDesc(a, b Product) int {
	if a.CreatedAt.After(b.CreatedAt) {
		return -1
	}
	return 1
}

func main() {
	var products ProductList
	products.Generate()
	for {
		fmt.Println("\nSorting Options:")
		fmt.Println("1. Sort by Name Ascending")
		fmt.Println("2. Sort by Name Descending")
		fmt.Println("3. Sort by Quantity Ascending")
		fmt.Println("4. Sort by Quantity Descending")
		fmt.Println("5. Sort by CreatedAt Ascending")
		fmt.Println("6. Sort by CreatedAt Descending")
		fmt.Println("7. Exit")
		var choice int
		fmt.Print("Enter your choice: ")
		fmt.Scan(&choice)
		switch choice {
		case 1:
			products.SetSortFn(SortNameAsc)
		case 2:
			products.SetSortFn(SortNameDesc)
		case 3:
			products.SetSortFn(SortQuantityAsc)
		case 4:
			products.SetSortFn(SortQuantityDesc)
		case 5:
			products.SetSortFn(SortCreatedAtAsc)
		case 6:
			products.SetSortFn(SortCreatedAtDesc)
		case 7:
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
			continue
		}
		products.Sort()
		fmt.Println("\nSorted Products:")
		products.Print()
		fmt.Println()
	}

}
