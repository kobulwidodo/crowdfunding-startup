package campaign

import "strings"

type CampaignFormat struct {
	Id               uint   `json:"id"`
	UserId           uint   `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageUrl         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
	Slug             string `json:"slug"`
}

type CampaignDetailFormat struct {
	Id               uint                   `json:"id"`
	Name             string                 `json:"name"`
	ShortDescription string                 `json:"ShortDescription"`
	Description      string                 `json:"Description"`
	ImageUrl         string                 `json:"image_url"`
	GoalAmount       int                    `json:"goal_amount"`
	CurrentAmount    int                    `json:"current_amount"`
	UserId           uint                   `json:"user_id"`
	Slug             string                 `json:"slug"`
	User             CampaignUserFormat     `json:"user"`
	Perks            []string               `json:"perks"`
	Images           []CampaignImagesFormat `json:"images"`
}

type CampaignUserFormat struct {
	Name      string `json:"name"`
	AvatarUrl string `json:"avata_url"`
}

type CampaignImagesFormat struct {
	ImageUrl  string `json:"image_url"`
	IsPrimary bool   `json:"is_primary"`
}

func CampaignDetailFormatter(campaign Campaign) CampaignDetailFormat {
	campaignDetail := CampaignDetailFormat{
		Id:               campaign.ID,
		Name:             campaign.Name,
		ShortDescription: campaign.ShortDescription,
		Description:      campaign.Description,
		GoalAmount:       campaign.GoalAmount,
		CurrentAmount:    campaign.CurrentAmount,
		UserId:           campaign.UserId,
		Slug:             campaign.Slug,
		ImageUrl:         "",
	}

	if len(campaign.CampaignImages) > 0 {
		campaignDetail.ImageUrl = campaign.CampaignImages[0].FileName
	}

	user := CampaignUserFormat{
		Name:      campaign.User.Name,
		AvatarUrl: campaign.User.AvatarFileName,
	}
	campaignDetail.User = user

	images := []CampaignImagesFormat{}
	for _, i := range campaign.CampaignImages {
		isPrimary := false
		if i.IsPrimary == 1 {
			isPrimary = true
		}
		images = append(images, CampaignImagesFormat{ImageUrl: i.FileName, IsPrimary: isPrimary})
	}
	campaignDetail.Images = images

	var perks []string
	for _, p := range strings.Split(campaign.Perks, ",") {
		perks = append(perks, strings.TrimSpace(p))
	}
	campaignDetail.Perks = perks

	return campaignDetail
}

func CampaignFormatter(campaign Campaign) CampaignFormat {
	campaignRes := CampaignFormat{
		Id:               campaign.ID,
		UserId:           campaign.UserId,
		Name:             campaign.Name,
		ShortDescription: campaign.ShortDescription,
		ImageUrl:         "",
		GoalAmount:       campaign.GoalAmount,
		CurrentAmount:    campaign.CurrentAmount,
		Slug:             campaign.Slug,
	}

	if len(campaign.CampaignImages) > 0 {
		campaignRes.ImageUrl = campaign.CampaignImages[0].FileName
	}

	return campaignRes
}

func CampaignsFormatter(campaigns []Campaign) []CampaignFormat {
	campaignsRes := []CampaignFormat{}
	for _, campaign := range campaigns {
		campaignFormat := CampaignFormatter(campaign)
		campaignsRes = append(campaignsRes, campaignFormat)
	}

	return campaignsRes
}
