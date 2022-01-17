package campaign

import (
	"bwastartup/user"

	"gorm.io/gorm"
)

type Campaign struct {
	gorm.Model
	UserId           uint
	Name             string
	ShortDescription string
	Description      string
	Perks            string
	BackerCount      int
	GoalAmount       int
	CurrentAmount    int
	Slug             string
	CampaignImages   []CampaignImage
	User             user.User
}

type CampaignImage struct {
	gorm.Model
	CampaignId uint
	FileName   string
	IsPrimary  int
}

type CampaignDetailInput struct {
	Id int `uri:"id" binding:"required"`
}

type CampaignInput struct {
	Name             string `json:"name" binding:"required"`
	ShortDescription string `json:"short_description" binding:"required"`
	Description      string `json:"description" binding:"required"`
	GoalAmount       int    `json:"goal_amount" binding:"required"`
	Perks            string `json:"perks" binding:"required"`
	User             user.User
}

type CampaignImageInput struct {
	CampaignId uint `form:"campaign_id" binding:"required"`
	IsPrimary  bool `form:"is_primary"`
	User       user.User
}
