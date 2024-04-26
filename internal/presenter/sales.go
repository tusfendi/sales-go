package presenter

import (
	"fmt"
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

type DownloadSalesRequest struct {
	UserID              int64
	Email               string
	Name                string
	StartDate           string `json:"start_date" binding:"required"`
	EndDate             string `json:"end_date" binding:"required"`
	StartDateIndoFormat string
	EndDateIndoFormat   string
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

func (r *DownloadSalesRequest) Validate() (string, bool) {
	var startDateTime, endDateTime time.Time
	var err error
	if startDateTime, err = time.Parse(constants.FormatDate, r.StartDate); err != nil {
		return "format tanggal mulai harus yyyy-mm-dd", true
	}
	r.StartDate = r.StartDate + " 00:00:00"

	if endDateTime, err = time.Parse(constants.FormatDate, r.EndDate); err != nil {
		return "format tanggal berakhir harus yyyy-mm-dd", true
	}
	r.EndDate = r.EndDate + " 23:59:59"

	r.StartDateIndoFormat = fmt.Sprintf("%v %v %v", startDateTime.Day(), constants.MappingMonthToBulan[startDateTime.Month().String()], startDateTime.Year())
	r.EndDateIndoFormat = fmt.Sprintf("%v %v %v", endDateTime.Day(), constants.MappingMonthToBulan[endDateTime.Month().String()], endDateTime.Year())

	return "", false
}
