package models

import "github.com/afaferz/web-app/db"

type Product struct {
	Id                int
	Name, Description string
	Price             float64
	Quantity          int
}

func SearchAllProducts() []Product {
	db := db.ConnectDB()
	defer db.Close()

	queryAllProducts := "SELECT * FROM products ORDER BY id ASC"

	selectAllProducts, err := db.Query(queryAllProducts)
	if err != nil {
		panic(err)
	}

	p := Product{}
	products := []Product{}

	for selectAllProducts.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err := selectAllProducts.Scan(&id, &name, &description, &price, &quantity)

		if err != nil {
			panic(err.Error())
		}
		p.Id = id
		p.Name = name
		p.Description = description
		p.Price = price
		p.Quantity = quantity

		products = append(products, p)
	}
	defer db.Close()
	return products
}

func CreateNewProduct(product Product) {
	db := db.ConnectDB()

	queryInsertNewProduct, err := db.Prepare("INSERT INTO products (name, description, price, quantity) values ($1, $2, $3, $4)")
	if err != nil {
		panic(err.Error())
	}
	queryInsertNewProduct.Exec(product.Name, product.Description, product.Price, product.Quantity)
	defer db.Close()
}

func EditProduct(productId string) Product {
	db := db.ConnectDB()

	queryToSearchProductToEdit, err := db.Query("SELECT * FROM products WHERE id=$1", productId)
	if err != nil {
		panic(err.Error())
	}
	productToEdit := Product{}

	for queryToSearchProductToEdit.Next() {
		var id, quantity int
		var name, description string
		var price float64

		err := queryToSearchProductToEdit.Scan(&id, &name, &description, &price, &quantity)

		if err != nil {
			panic(err.Error())
		}
		productToEdit.Id = id
		productToEdit.Name = name
		productToEdit.Description = description
		productToEdit.Price = price
		productToEdit.Quantity = quantity
	}
	defer db.Close()
	return productToEdit
}

func UpdateProduct(product Product) {
	db := db.ConnectDB()

	queryUpdateProduct, err := db.Prepare("UPDATE products SET name=$1, description=$2, price=$3, quantity=$4 WHERE id=$5")
	if err != nil {
		panic(err.Error())
	}

	queryUpdateProduct.Exec(product.Name, product.Description, product.Price, product.Quantity, product.Id)
	defer db.Close()

}

func DeleteProduct(productId string) {
	db := db.ConnectDB()
	queryToDelectProduct, err := db.Prepare("DELETE FROM products WHERE id=$1")

	if err != nil {
		panic(err.Error())
	}

	queryToDelectProduct.Exec(productId)
	defer db.Close()
}
