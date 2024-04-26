package delivery

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/tusfendi/sales-go/internal/constants"
	"github.com/tusfendi/sales-go/internal/presenter"
	"github.com/tusfendi/sales-go/internal/usecase"
)

type SalesDelivery struct {
	salesUsecase usecase.SalesUsecase
}

func NewSalesDelivery(salesUsecase usecase.SalesUsecase) SalesDelivery {
	return SalesDelivery{
		salesUsecase: salesUsecase,
	}
}

func (c *SalesDelivery) Mount(group *gin.RouterGroup) {

	group.POST("", c.CreateSales)
	group.GET("/download", c.DownloadSales)
}

func (u *SalesDelivery) CreateSales(c *gin.Context) {
	var params presenter.SalesRequest
	err := c.ShouldBindJSON(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response": constants.Failed, "error": constants.ErrData})
		return
	}

	params.UserID = int64(c.MustGet("id").(float64))
	params.Email = c.GetString("email")
	if httpCode, err := u.salesUsecase.CreateSales(params); err != nil {
		c.JSON(httpCode, gin.H{"response": constants.Failed, "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"response": constants.Success})
}

func (u *SalesDelivery) DownloadSales(c *gin.Context) {
	var params presenter.DownloadSalesRequest
	today := time.Now().Format(constants.FormatDate)
	params.UserID = int64(c.MustGet("id").(float64))
	params.Email = c.GetString("email")
	params.Name = c.GetString("name")
	params.StartDate = c.DefaultQuery("start_date", today)
	params.EndDate = c.DefaultQuery("end_date", today)
	u.salesUsecase.DownloadSales(c, params)
}
