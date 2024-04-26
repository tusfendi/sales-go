package presenter

import (
	"time"

	"github.com/tusfendi/sales-go/internal/constants"
)

type SalesRequest struct {
	UserID                int64
	Email                 string
	Type                  string `json:"type" binding:"required"`
	Nominal               int64  `json:"nominal" binding:"required"`
	TransactionDateString string `json:"transaction_date" binding:"required"`
	TransactionDate       time.Time
}

var IsValidType = map[string]bool{
	constants.SalesTypeGoods:    true,
	constants.SalesTypeServices: true,
}

func (r *SalesRequest) Validate() (string, bool) {
	if r.Type == "" || !IsValidType[r.Type] {
		return "Jenis transaksi tidak valid", true
	}

	date, err := time.Parse(constants.FormatDate, r.TransactionDateString)
	if err != nil {
		return "format tanggal yyyy-mm-dd", true
	}

	r.TransactionDate = date
	return "", false
}
