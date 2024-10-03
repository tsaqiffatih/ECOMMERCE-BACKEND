package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/tsaqiffatih/ecommerce-bakcend/config"
	"github.com/tsaqiffatih/ecommerce-bakcend/models"
)

func CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	json.NewDecoder(r.Body).Decode(&product)

	config.DB.Create(&product)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(product)
}

func GetProducts(w http.ResponseWriter, r *http.Request) {
	var products []models.Product
	config.DB.Find(&products)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(products)
}
