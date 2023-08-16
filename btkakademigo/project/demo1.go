package project

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Product struct {
	Id          int     `json:"id"`
	ProductName string  `json:"productName"`
	CategoryId  int     `json:"categoryId"`
	UnitPrice   float64 `json:"unitPrice"`
}

type Category struct {
	Id           int    `json:"id"`
	CategoryName string `json:"categoryName"`
}

func GetAllProducts() ([]Product, error) {
	response, err := http.Get("http://localhost:3000/products")
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	bodyBytes, _ := io.ReadAll(response.Body)
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

	var products []Product
	json.Unmarshal(bodyBytes, &products)
	return products, nil

}

func AddProduct() (Product, error) {
	product := Product{Id: 10, ProductName: "DellSonifDeneme", CategoryId: 2, UnitPrice: 22323.56}
	productJson, _ := json.Marshal(product)
	response, err := http.Post("http://localhost:3000/products",
		"application/json;charset=utf-8", bytes.NewBuffer(productJson))
	if err != nil {
		return product, err
	}
	defer response.Body.Close()

	bodyBytes, _ := io.ReadAll(response.Body)

	var productResponse Product
	json.Unmarshal(bodyBytes, &productResponse)

	return productResponse, nil

	// if product.Id != productResponse.Id &&
	// 	product.ProductName ==
	// 		productResponse.ProductName && product.CategoryId ==
	// 	productResponse.CategoryId && product.UnitPrice == productResponse.UnitPrice {
	// 	fmt.Println("Farklı id ile kaydedildi.")
	// } else if product.Id == productResponse.Id {
	// 	fmt.Println("Kaydedildi")
	// } else {
	// 	fmt.Println("Ürün kaydedilmedi.")
	// }

}
