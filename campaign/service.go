package campaign

import (
	"errors"
	"fmt"

	"github.com/gosimple/slug"
)

type Service interface {
	GetCampaigns(userId uint) ([]Campaign, error)
	GetCampaignById(input CampaignDetailInput) (Campaign, error)
	CreateCampaign(input CampaignInput) (Campaign, error)
	UpdateCampaign(inputId CampaignDetailInput, inputData CampaignInput) (Campaign, error)
	SaveCampaignImage(input CampaignImageInput, fileLocation string) (CampaignImage, error)
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

func (s *service) CreateCampaign(input CampaignInput) (Campaign, error) {
	newCampaign := Campaign{
		UserId:           input.User.ID,
		Name:             input.Name,
		ShortDescription: input.ShortDescription,
		Description:      input.Description,
		Perks:            input.Perks,
		BackerCount:      0,
		GoalAmount:       input.GoalAmount,
		CurrentAmount:    0,
		Slug:             fmt.Sprintf("%s-%d", slug.Make(input.Name), input.User.ID),
	}

	campaign, err := s.repository.Save(newCampaign)
	if err != nil {
		return campaign, err
	}

	return campaign, nil
}

func (s *service) UpdateCampaign(inputId CampaignDetailInput, inputData CampaignInput) (Campaign, error) {
	campaign, err := s.repository.FindById(uint(inputId.Id))
	if err != nil {
		return campaign, err
	}

	if campaign.User.ID != inputData.User.ID {
		return campaign, errors.New("Tidak memiliki akses")
	}

	campaign.Name = inputData.Name
	campaign.ShortDescription = inputData.ShortDescription
	campaign.Description = inputData.Description
	campaign.GoalAmount = inputData.GoalAmount
	campaign.Perks = inputData.Perks

	newCampaign, err := s.repository.Update(campaign)
	if err != nil {
		return campaign, err
	}

	return newCampaign, nil
}

func (s *service) SaveCampaignImage(input CampaignImageInput, fileLocation string) (CampaignImage, error) {
	campaign, err := s.repository.FindById(input.CampaignId)
	if err != nil {
		return CampaignImage{}, err
	}
	if campaign.UserId != input.User.ID {
		return CampaignImage{}, errors.New("Tidak memiliki akses")
	}
	isPrimary := 0
	if input.IsPrimary {
		if _, err := s.repository.MarkAllImagesAsNonPrimary(input.CampaignId); err != nil {
			return CampaignImage{}, err
		}
		isPrimary = 1
	}

	campaignImage := CampaignImage{
		CampaignId: input.CampaignId,
		FileName:   fileLocation,
		IsPrimary:  isPrimary,
	}

	newImage, err := s.repository.CreateImage(campaignImage)
	if err != nil {
		return newImage, err
	}

	return newImage, nil
}
