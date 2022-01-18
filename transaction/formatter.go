package transaction

import "time"

type TransactionByCampaignIdFormatter struct {
	ID        uint      `json:"id"`
	Name      string    `json:"name"`
	Amount    int       `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
}

func FormatTransactionByCampaignId(transaction Transaction) TransactionByCampaignIdFormatter {
	formatter := TransactionByCampaignIdFormatter{
		ID:        transaction.ID,
		Name:      transaction.User.Name,
		Amount:    transaction.Amount,
		CreatedAt: transaction.CreatedAt,
	}
	return formatter
}

func FormatTransactionsByCampaignId(transactions []Transaction) []TransactionByCampaignIdFormatter {
	if len(transactions) == 0 {
		return []TransactionByCampaignIdFormatter{}
	}

	var transactionByCampaignIdFormatter []TransactionByCampaignIdFormatter
	for _, transaction := range transactions {
		formatter := FormatTransactionByCampaignId(transaction)
		transactionByCampaignIdFormatter = append(transactionByCampaignIdFormatter, formatter)
	}

	return transactionByCampaignIdFormatter
}
