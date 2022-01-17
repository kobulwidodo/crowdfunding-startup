package handler

import (
	"bwastartup/campaign"
	"bwastartup/helper"
	"bwastartup/user"
	"fmt"
	"net/http"
	"os"
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

func (h *campaignHandler) CreateCampaign(c *gin.Context) {
	var cInput campaign.CampaignInput
	err := c.ShouldBindJSON(&cInput)
	if err != nil {
		response := helper.ApiResponse("Harap isi semua input", http.StatusUnprocessableEntity, "gagal", helper.FormatBindError(err))
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}

	userLoggedin := c.MustGet("userLoggedin").(user.User)
	cInput.User = userLoggedin

	var newCampaign campaign.Campaign
	newCampaign, err = h.campaignService.CreateCampaign(cInput)
	if err != nil {
		response := helper.ApiResponse("Gagal menambah campaign baru", http.StatusInternalServerError, "gagal", nil)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	response := helper.ApiResponse("Berhasil membuat campaign baru", http.StatusOK, "sukses", campaign.CampaignFormatter(newCampaign))
	c.JSON(http.StatusOK, response)
	return
}

func (h *campaignHandler) UpdateCampaign(c *gin.Context) {
	var input campaign.CampaignDetailInput
	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.ApiResponse("ID tidak tersedia", http.StatusBadRequest, "gagal", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	var cInput campaign.CampaignInput
	err = c.ShouldBindJSON(&cInput)
	if err != nil {
		response := helper.ApiResponse("Harap isi semua input", http.StatusUnprocessableEntity, "gagal", helper.FormatBindError(err))
		c.JSON(http.StatusUnprocessableEntity, response)
		return
	}
	userLoggedin := c.MustGet("userLoggedin").(user.User)
	cInput.User = userLoggedin

	updatedCampaign, err := h.campaignService.UpdateCampaign(input, cInput)
	if err != nil {
		response := helper.ApiResponse("Gagal mengupdate campaign", http.StatusBadRequest, "gagal", err.Error())
		c.JSON(http.StatusBadRequest, response)
		return
	}

	response := helper.ApiResponse("Berhasil mengubah campaign", http.StatusOK, "sukses", campaign.CampaignFormatter(updatedCampaign))
	c.JSON(http.StatusOK, response)
}

func (h *campaignHandler) UploadImage(c *gin.Context) {
	var input campaign.CampaignImageInput

	if err := c.ShouldBind(&input); err != nil {
		response := helper.ApiResponse("Gagal mengunggah gambar", http.StatusBadRequest, "gagal", helper.FormatBindError(err))
		c.JSON(http.StatusBadRequest, response)
		return
	}

	file, err := c.FormFile("file")
	if err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.ApiResponse("Gagal mendapatkan gambar", http.StatusBadRequest, "gagal", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	userLoggedin := c.MustGet("userLoggedin").(user.User)
	userId := userLoggedin.ID
	path := fmt.Sprintf("images/%d-%s", userId, file.Filename)

	if err := c.SaveUploadedFile(file, path); err != nil {
		data := gin.H{"is_uploaded": false}
		response := helper.ApiResponse("Gagal menyimpan gambar", http.StatusBadRequest, "gagal", data)
		c.JSON(http.StatusBadRequest, response)
		return
	}

	input.User = userLoggedin

	if _, err := h.campaignService.SaveCampaignImage(input, path); err != nil {
		os.Remove(path)
		data := gin.H{"is_uploaded": false}
		response := helper.ApiResponse("Gagal mengalokasikan gambar", http.StatusInternalServerError, "gagal", data)
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	data := gin.H{"is_uploaded": true}
	response := helper.ApiResponse("Berhasil mengunggah gambar", http.StatusOK, "sukses", data)
	c.JSON(http.StatusOK, response)
	return
}
