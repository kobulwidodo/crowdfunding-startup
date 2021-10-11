package handler

import (
	"bwastartup/campaign"
	"bwastartup/helper"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type campaignHandler struct {
	campaignService campaign.Service
}

func NewCampaignHandler(campaignService campaign.Service) *campaignHandler {
	return &campaignHandler{campaignService}
}

func (h *campaignHandler) GetCampaigns(c *gin.Context) {
	userId, _ := strconv.Atoi(c.Query("user_id"))

	campaigns, err := h.campaignService.GetCampaigns(uint(userId))
	if err != nil {
		response := helper.ApiResponse("Gagal mendapatkan campaigns", http.StatusInternalServerError, "gagal", nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := helper.ApiResponse("Sukses mendapatkan campaigns", http.StatusOK, "sukses", campaigns)
	c.JSON(http.StatusOK, response)
}
