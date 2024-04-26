package usecase

import (
	"bytes"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/tusfendi/sales-go/internal/constants"
	"github.com/tusfendi/sales-go/internal/entity"
	"github.com/tusfendi/sales-go/internal/presenter"
	"github.com/tusfendi/sales-go/internal/repository"
	"github.com/xuri/excelize/v2"
)

type SalesUsecase interface {
	CreateSales(params presenter.SalesRequest) (int, error)
	DownloadSales(c *gin.Context, params presenter.DownloadSalesRequest)
}

type salesCtx struct {
	salesRepo repository.SalesRepository
}

func NewSalesUsecase(salesRepo repository.SalesRepository) SalesUsecase {
	return &salesCtx{
		salesRepo: salesRepo,
	}
}

func (u *salesCtx) CreateSales(params presenter.SalesRequest) (int, error) {
	if errMessage, isError := params.Validate(); isError {
		return http.StatusBadRequest, errors.New(errMessage)
	}

	if err := u.salesRepo.CreateSales(&entity.Sales{
		UserID:          params.UserID,
		Type:            params.Type,
		TransactionDate: params.TransactionDate,
		Nominal:         params.Nominal,
	}); err != nil {
		return http.StatusInternalServerError, errors.New(constants.ErrDB)
	}

	return http.StatusCreated, nil
}

func (u *salesCtx) DownloadSales(c *gin.Context, params presenter.DownloadSalesRequest) {
	if errMessage, isError := params.Validate(); isError {
		c.JSON(http.StatusBadRequest, gin.H{"response": constants.Failed, "error": errMessage})
		return
	}

	data, err := u.salesRepo.GetData(params)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"response": constants.Failed, "error": constants.ErrDB})
	}

	today := time.Now().Unix()
	sheetName := "Sheet1"
	fileName := fmt.Sprintf("%s%d%s", "data-sales-", today, ".xlsx")
	f := excelize.NewFile()
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()
	var index int
	for idx, row := range [][]interface{}{
		{"Requestor", fmt.Sprintf("%s (%s)", params.Name, params.Email)},
		{nil},
		{"Parameter"},
		{"Start Date", params.StartDateIndoFormat},
		{"End Date", params.EndDateIndoFormat},
		{nil},
		{"User", "Jumlah Hari Kerja", "Jumlah Transaksi Barang", "Jumlah Transaksi Jasa", "Nominal Transaksi Barang", "Nominal Transaksi Jasa"},
	} {
		cell, err := excelize.CoordinatesToCellName(1, idx+1)
		index = idx
		if err != nil {
			fmt.Println(err)
			return
		}
		f.SetSheetRow(sheetName, cell, &row)
	}
	index++

	styleID, err := f.NewStyle(&excelize.Style{CustomNumFmt: &constants.FmtCode})
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, sales := range data {
		cell, err := excelize.CoordinatesToCellName(1, index+1)
		if err != nil {
			fmt.Println(err)
			return
		}
		exportedData := []interface{}{
			sales.Name,
			sales.DaysTotal,
			sales.GoodsTransactionTotal,
			sales.TransactionTotal - sales.GoodsTransactionTotal,
			sales.GoodsNominalTotal,
			sales.NominalTotal - sales.GoodsNominalTotal,
		}
		f.SetSheetRow(sheetName, cell, &exportedData)
		index++
	}

	if err := f.SetCellStyle(sheetName, "E8", fmt.Sprintf("F%d", index), styleID); err != nil {
		fmt.Println(err)
		return
	}

	if err := f.SetColWidth(sheetName, "B", "F", float64(25)); err != nil {
		fmt.Println(err)
		return
	}

	style, _ := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold: true,
		},
	})

	if err := f.SetCellStyle(sheetName, "A1", "A7", style); err != nil {
		fmt.Println(err)
		return
	}
	if err := f.SetCellStyle(sheetName, "B7", "F7", style); err != nil {
		fmt.Println(err)
		return
	}

	if err := f.SaveAs(fileName); err != nil {
		fmt.Println(err)
	}

	var b bytes.Buffer
	if err := f.Write(&b); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	downloadName := time.Now().UTC().Format(fileName)
	c.Header("Content-Description", "File Transfer")
	c.Header("Content-Disposition", "attachment; filename="+downloadName)
	c.Data(http.StatusOK, "application/octet-stream", b.Bytes())
}
