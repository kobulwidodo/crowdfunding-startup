package campaign

import "errors"

type Service interface {
	GetCampaigns(userId uint) ([]Campaign, error)
	GetCampaignById(input CampaignDetailInput) (Campaign, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetCampaigns(userId uint) ([]Campaign, error) {
	var campaings []Campaign
	if userId != 0 {
		campaings, err := s.repository.FindByUserId(int(userId))
		if err != nil {
			return campaings, err
		}

		return campaings, nil
	}

	campaings, err := s.repository.FindAll()
	if err != nil {
		return campaings, err
	}

	return campaings, nil
}

func (s *service) GetCampaignById(input CampaignDetailInput) (Campaign, error) {
	var campaign Campaign
	campaign, err := s.repository.FindById(uint(input.Id))
	if err != nil {
		return campaign, err
	}

	if campaign.ID == 0 {
		return campaign, errors.New("Tidak dapat mendapatkan dengan id tersebut")
	}

	return campaign, nil
}
