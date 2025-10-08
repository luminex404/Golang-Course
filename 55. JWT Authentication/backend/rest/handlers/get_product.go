package handlers

import (
	"ecommerce/database"
	"ecommerce/util"
	"net/http"
	"strconv"
)

func GetProduct(w http.ResponseWriter, r *http.Request) {
	productID := r.PathValue("id")

	pId, err := strconv.Atoi(productID)
	if err != nil {
		http.Error(w, "Please provide a valid product ID", 400)
	}
    
	 product := database.Get(pId)
	 
	 if product == nil{
		util.SendError(w,"Product not found",404)
		return
	 }
	 
    util.SendData(w, product, 200)
	 
}
