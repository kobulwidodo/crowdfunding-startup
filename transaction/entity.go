package transaction

import (
	"bwastartup/campaign"
	"bwastartup/user"

	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	Amount     int
	Status     string
	Code       string
	CampaignId uint
	UserId     uint
	User       user.User
	Campaign   campaign.Campaign
}

type GetTransactionByCampaignIdInput struct {
	Id uint `uri:"id" binding:"required"`
}
