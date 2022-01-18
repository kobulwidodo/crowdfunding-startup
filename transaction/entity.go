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
	Campaign   campaign.Campaign
	User       user.User
}
