package data

import (
    "fmt"
)

//ErrProductNotFound is an error raised when a product cannot be found in th database
var ErrProductNotFound = fmt.Errorf("Product not found")

//Product defines the structure for an API product
type Product struct {
    ID int `json:"id"`   //Unique identifier for the product
    Name string `json:"name" validate:"required"`
    Description string `json: "description"`
    Price float32 `json:"price" validate:"required, gt=0"`
    //The sku for the product
    //required: true
    //pattern: [a-z]+-[a-z]+-[a-z]+
    SKU string `json: "sku" valiate:"sku"`
}

//Product defined a slice of Product
type Products []*Product

//GetProducts returns all products from the database
func GetProducts() Products {
    return productList
}

// GetProductByID returns a single product which matches the id from the database
// if a product is not found this function return a ProductNotFound error

func GetProductByID(id int) (*Product, error) {
    i := findIndexByProductID(id)
    if id == -1 {
        return nil, ErrProductNotFound
    }
    return productList[i], nil
}


//updateProduct replaces a product in the database with the given item
//If a product with the given id doesnt exist in the database
//this function returns a ProductNotFound error

func UpdateProduct(p Product) error {
    i := findIndexByProductID(p.ID)
    if i == -1 {
        return ErrProductNotFound
    }

    //update the product in the DB
    productList[i] = &p
    return nil

}

//Addproduct adds a new product to the database
func AddProduct(p Product) {
    // get the next id in the sequence
    maxID := productList[len(productList)-1].ID
    p.ID = maxID + 1
    productList = append(productList,   &p)

}


//Deleteproduct deletes a product from the database
func DeleteProduct(id int) error {
    i := findIndexByProductID(id)
    if i == -1 {
        return ErrProductNotFound
    }

    productList = append(productList[:1], productList[i+1])

    return nil
}


//findIndex finds the index of a product in the databas
//returns -1 when no product can be found
func findIndexByProductID(id int) int {
     for i, p := range productList {
          if p.ID == id {
              return i
          }
     }
     return -1

}



var productList = []*Product{
    &Product{
        ID:   1,
        Name:   "Latte",
        Description: "Frothy milky coffee",
        Price: 2.45,
        SKU:    "abc323",
    },
    &Product{
        ID:    2,
        Name:   "Esspresso",
        Description: "Short and Strong coffee with milk",
        Price:    1.99,
        SKU: "fjd34",


    },
}
