package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/afaferz/web-app/models"
)

var template_html = template.Must(template.ParseGlob("templates/*.html"))

func Index(wr http.ResponseWriter, req *http.Request) {
	allProducts := models.SearchAllProducts()
	template_html.ExecuteTemplate(wr, "Index", allProducts)
}
func New(wr http.ResponseWriter, req *http.Request) {
	// allProducts := models.SearchAllProducts()
	template_html.ExecuteTemplate(wr, "New", nil)
}

func Insert(wr http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		name := req.FormValue("name")
		description := req.FormValue("description")
		price := req.FormValue("price")
		quantity := req.FormValue("quantity")

		priceToFloat, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("An error occurred in price convert", err)
		}
		quantityToInt, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("An error occurred in quantity convert", err)
		}

		product := models.Product{
			Name:        name,
			Description: description,
			Price:       priceToFloat,
			Quantity:    quantityToInt,
		}

		models.CreateNewProduct(product)
	}
	http.Redirect(wr, req, "/", http.StatusMovedPermanently)
}

func Edit(wr http.ResponseWriter, req *http.Request) {
	productId := req.URL.Query().Get("id")
	productToEdit := models.EditProduct(productId)
	template_html.ExecuteTemplate(wr, "Edit", productToEdit)
}

func Update(wr http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		id := req.FormValue("id")
		name := req.FormValue("name")
		description := req.FormValue("description")
		price := req.FormValue("price")
		quantity := req.FormValue("quantity")

		idToInt, err := strconv.Atoi(id)
		if err != nil {
			log.Println("An error occurred in quantity convert", err)
		}
		priceToFloat, err := strconv.ParseFloat(price, 64)
		if err != nil {
			log.Println("An error occurred in price convert", err)
		}
		quantityToInt, err := strconv.Atoi(quantity)
		if err != nil {
			log.Println("An error occurred in quantity convert", err)
		}

		product := models.Product{
			Id:          idToInt,
			Name:        name,
			Description: description,
			Price:       priceToFloat,
			Quantity:    quantityToInt,
		}
		models.UpdateProduct(product)
	}
	http.Redirect(wr, req, "/", http.StatusMovedPermanently)
}

func Delete(wr http.ResponseWriter, req *http.Request) {
	productId := req.URL.Query().Get("id")
	models.DeleteProduct(productId)
	http.Redirect(wr, req, "/", http.StatusMovedPermanently)
}
