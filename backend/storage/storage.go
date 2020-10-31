package storage

import (
	"encoding/json"
	"errors"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// AccountType ...
type AccountType int

// Account types
const (
	_ AccountType = iota
	ATBank
	ATCreditCard
	ATCash
	ATInvestment
)

// TransactionType ...
type TransactionType int

// Transaction types
const (
	_ TransactionType = iota
	TTExpense
	TTIncome
	TTTransfer
	TTInvestment
)

var (
	transactionTypeMap = map[TransactionType]string{
		TTExpense:    "EXPENSE",
		TTIncome:     "INCOME",
		TTTransfer:   "TRANSFER",
		TTInvestment: "INVESTMENT",
	}
)

func (t TransactionType) String() string {
	if v, ok := transactionTypeMap[t]; ok {
		return v
	}

	return "UNKNOWN"
}

func (t *TransactionType) UnmarshalJSON(data []byte) error {
	var val string
	err := json.Unmarshal(data, &val)
	if err != nil {
		return err
	}

	for k, v := range transactionTypeMap {
		if v == val {
			*t = k
			return nil
		}
	}

	return errors.New("invalid value")
}

// Account ...
type Account struct {
	ID   uint        `gorm:"primaryKey"`
	Name string      `gorm:"type:varchar(255)"`
	Type AccountType `gorm:"type:tinyint"`
}

// Category ...
type Category struct {
	ID          uint            `gorm:"primaryKey"`
	Category    string          `gorm:"type:varchar(128)"`
	SubCategory string          `gorm:"type:varchar(128)"`
	TxType      TransactionType `gorm:"type:tinyint"`
}

// Ledger ...
type Ledger struct {
	ID          uint    `gorm:"primaryKey"`
	AccountID   uint    `gorm:"uniqueIndex:idx_account_version,priority:1"`
	Version     uint    `gorm:"uniqueIndex:idx_account_version,priority:2"`
	Balance     float64 `gorm:"type:decimal(18,3)"`
	SubCategory string  `gorm:"type:varchar(128)"`
	Amount      float64 `gorm:"type:decimal(18,3)"`
	TxTime      time.Time
}

// Init initializes DB connections..
func Init() (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open("root:secret@tcp(127.0.0.1:3306)/expenses?parseTime=True"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err = db.AutoMigrate(&Ledger{}, &Account{}, &Category{}); err != nil {
		return nil, err
	}

	return db, nil
}
