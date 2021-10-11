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

	response := helper.ApiResponse("Sukses mendapatkan campaigns", http.StatusOK, "sukses", campaign.CampaignsFormatter(campaigns))
	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) GetCampaign(c *gin.Context) {
	var input campaign.CampaignDetailInput

	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.ApiResponse("ID tidak tersedia", http.StatusBadRequest, "gagal", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var dataCampaign campaign.Campaign
	dataCampaign, err = h.campaignService.GetCampaignById(input)
	if err != nil {
		response := helper.ApiResponse("Gagal mendapatkan Campaign", http.StatusBadRequest, "gagal", gin.H{"errors": err.Error()})
		c.JSON(http.StatusBadRequest, response)
		return
	}

	campaignRes := campaign.CampaignDetailFormatter(dataCampaign)

	response := helper.ApiResponse("Sukses mendapatkan Campaign", http.StatusOK, "sukses", campaignRes)
	c.JSON(http.StatusOK, response)
}
