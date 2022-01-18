package transaction

type service struct {
	repository Repository
}

type Service interface {
	GetTransactionByCampaignId(input GetTransactionByCampaignIdInput) ([]Transaction, error)
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetTransactionByCampaignId(input GetTransactionByCampaignIdInput) ([]Transaction, error) {
	transactions, err := s.repository.FindByCampaignId(input.Id)
	if err != nil {
		return transactions, err
	}

	return transactions, nil
}
