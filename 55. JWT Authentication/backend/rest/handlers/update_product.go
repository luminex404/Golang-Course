package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func UpdatedProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")

	pId, err := strconv.Atoi(productID)
	if err != nil {
		http.Error(w, "Please provide a valid product ID", 400)
	}

	var newProduct database.Product

	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&newProduct)

	if err != nil {
		fmt.Println(err)
		http.Error(w, "Plz give me valid JSON", 400)
	}

	newProduct.ID = pId

	database.Update(newProduct)

	util.SendData(w,"Successfully updated product",201)

}
