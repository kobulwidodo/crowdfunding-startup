package campaign

type CampaignFormat struct {
	Id               uint   `json:"id"`
	UserId           uint   `json:"user_id"`
	Name             string `json:"name"`
	ShortDescription string `json:"short_description"`
	ImageUrl         string `json:"image_url"`
	GoalAmount       int    `json:"goal_amount"`
	CurrentAmount    int    `json:"current_amount"`
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
