package usecase

import (
	"net/http"

	"github.com/pkg/errors"
	"github.com/tusfendi/sales-go/internal/constants"
	"github.com/tusfendi/sales-go/internal/entity"
	"github.com/tusfendi/sales-go/internal/presenter"
	"github.com/tusfendi/sales-go/internal/repository"
)

type SalesUsecase interface {
	CreateSales(params presenter.SalesRequest) (int, error)
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
