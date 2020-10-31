package server

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/bongnv/expenses/backend/storage"
	"github.com/bongnv/gwf"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

var validate = validator.New()

// CreateTransactionRequest ...
type CreateTransactionRequest struct {
	Type          storage.TransactionType `json:"type" validate:"required"`
	Amount        float64                 `json:"amount" validate:"required"`
	FromAccountID uint                    `json:"fromAccountID"`
	ToAccountID   uint                    `json:"toAccountID"`
	SubCategory   string                  `json:"subCategory" validate:"required"`
	TxTime        time.Time               `json:"txTime" validate:"required"`
}

func (s *Server) createTransaction(ctx context.Context, gwfReq gwf.Request) (interface{}, error) {
	req := &CreateTransactionRequest{}
	if err := gwfReq.Decode(req); err != nil {
		return nil, err
	}

	if err := validate.Struct(req); err != nil {
		s.logger.Println("Validation failed", err)
		return nil, &gwf.HTTPError{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		}
	}

	s.logger.Println("creating a new transaction", req)
	// get latest ledger
	switch req.Type {
	case storage.TTExpense:
		var record storage.Ledger
		result := s.db.Where("account_id = ?", req.FromAccountID).Order("version desc").First(&record)
		if result.Error == gorm.ErrRecordNotFound {
			result = s.db.Save(&storage.Ledger{
				AccountID:   req.FromAccountID,
				Version:     1,
				Balance:     -req.Amount,
				SubCategory: req.SubCategory,
				Amount:      -req.Amount,
				TxTime:      req.TxTime,
			})
			if result.Error != nil {
				s.logger.Println("Error while saving records to ledger", result.Error)
				return nil, result.Error
			}
		}

		if result.Error != nil {
			s.logger.Println("Error while finding records from the Ledger", result.Error)
			return nil, result.Error
		}
	default:
		return nil, errors.New("unsupported type")
	}
	return nil, nil
}
