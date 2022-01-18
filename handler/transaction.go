package handler

import (
	"bwastartup/helper"
	"bwastartup/transaction"
	"net/http"

	"github.com/gin-gonic/gin"
)

type transactionHandler struct {
	service transaction.Service
}

func NewTransactionHandler(service transaction.Service) *transactionHandler {
	return &transactionHandler{service}
}

func (h *transactionHandler) GetTransactionByCampaignId(c *gin.Context) {
	var input transaction.GetTransactionByCampaignIdInput

	if err := c.ShouldBindUri(&input); err != nil {
		response := helper.ApiResponse("Parameter Id tidak ditemukan", http.StatusBadRequest, "gagal", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	transactions, err := h.service.GetTransactionByCampaignId(input)
	if err != nil {
		response := helper.ApiResponse("Gagal mendapatkan data transaksi", http.StatusBadRequest, "gagal", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("Sukses mengambil data", http.StatusOK, "sukses", transaction.FormatTransactionsByCampaignId(transactions))
	c.JSON(http.StatusOK, response)
	return
}
