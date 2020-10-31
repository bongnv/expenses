package server

import (
	"context"

	"github.com/bongnv/gwf"
)

type ListTransactionsResponse struct {
	Items []Transaction `json:"items"`
}

// Transaction ...
type Transaction struct {
	Type        string  `json:"type"`
	Amount      float64 `json:"amount"`
	SubCategory string  `json:"subCategory"`
	Category    string  `json:"category"`
	Account     Account `json:"account"`
}

// Account ...
type Account struct {
	ID   int    `json:"id"`
	Type string `json:"type"`
	Name string `json:"name"`
}

func (s *Server) listTransactions(ctx context.Context, req gwf.Request) (interface{}, error) {
	s.logger.Println("list all transactions")
	return &ListTransactionsResponse{
		Items: []Transaction{
			{
				Type: "EXPENSE",
			},
		},
	}, nil
}
