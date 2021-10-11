package campaign

type Service interface {
	GetCampaigns(userId uint) ([]Campaign, error)
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
