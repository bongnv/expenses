package server

import (
	"context"

	"github.com/bongnv/gwf"
)

// CreateTransactionRequest ...
type CreateTransactionRequest struct {
	Type          string  `json:"type"`
	Amount        float64 `json:"amount"`
	FromAccountID int     `json:"fromAccountID"`
	ToAccountID   int     `json:"toAccountID"`
	SubCategory   string  `json:"subCategory"`
}

func (s *Server) createTransaction(ctx context.Context, gwfReq gwf.Request) (interface{}, error) {
	req := &CreateTransactionRequest{}
	if err := gwfReq.Decode(req); err != nil {
		return nil, err
	}

	s.logger.Println("creating a new transaction", req)
	return nil, nil
}
