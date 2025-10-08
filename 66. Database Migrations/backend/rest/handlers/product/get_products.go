package product

import (
	"ecommerce/util"
	"net/http"
)

func (h *Handler) GetProducts(w http.ResponseWriter, r *http.Request) {
	productList, err := h.productRepo.List()

	if err != nil {
		util.SendError(w, http.StatusBadRequest, "Internal Server Error")

	}
	util.SendData(w, http.StatusOK, productList)
}
