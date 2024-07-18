package jsonparser

import (
	"encoding/json"
	
	"os"
	"log"
	"fmt"
)

type Product struct {
	ID                 int     `json:"id"`
	Name               string  `json:"name"`
	ProductValue       float64 `json:"productValue"`
	ProductOriginalValue float64 `json:"productOriginalValue"`
	ProductDiscount    string  `json:"productDiscount"`
	ProductLink        string  `json:"productLink"`
	ProductImageLink   string  `json:"productImageLink"`
}

type Products struct {
	Products []Product `json:"products"`
}

func Jsonparser() {
	filePath := "./jsonparser/products.json"
	products := readFile(filePath)
	iterateProducts(products)
}

func readFile(filePath string) Products {	
	byteValue, err := os.ReadFile(filePath)
	if err != nil {
		log.Fatalf("failed to read file: %s", err)
	}

	
	var products Products
	err = json.Unmarshal(byteValue, &products)
	if err != nil {
		log.Fatalf("failed to unmarshal JSON: %s", err)
	}
	return products
}

func iterateProducts(products Products) {
	
	for _, product := range products.Products {
		fmt.Printf("ID: %d\n", product.ID)
		fmt.Printf("Name: %s\n", product.Name)
		fmt.Printf("Value: %.2f\n", product.ProductValue)
		fmt.Printf("Original Value: %.2f\n", product.ProductOriginalValue)
		fmt.Printf("Discount: %s\n", product.ProductDiscount)
		fmt.Printf("Link: %s\n", product.ProductLink)
		fmt.Printf("Image Link: %s\n", product.ProductImageLink)
		fmt.Println()
	}
}