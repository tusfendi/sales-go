package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tusfendi/sales-go/internal/constants"
	"github.com/tusfendi/sales-go/internal/presenter"
	"github.com/tusfendi/sales-go/internal/usecase"
)

type UserDelivery struct {
	userUsecase usecase.UserUsecase
}

func NewUserDelivery(userUsecase usecase.UserUsecase) UserDelivery {
	return UserDelivery{
		userUsecase: userUsecase,
	}
}

func (u *UserDelivery) Mount(group *gin.RouterGroup) {
	group.POST("/registration", u.Registration)
	group.POST("/auth", u.AuthUser)
}

func (u *UserDelivery) Registration(c *gin.Context) {
	var params presenter.Registration
	err := c.ShouldBindJSON(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response": constants.Failed, "error": constants.ErrData})
		return
	}

	if httpCode, err := u.userUsecase.CreateUser(params); err != nil {
		c.JSON(httpCode, gin.H{"response": constants.Failed, "error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"response": constants.Success})
}

func (u *UserDelivery) AuthUser(c *gin.Context) {
	var params presenter.Auth
	err := c.ShouldBindJSON(&params)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"response": constants.Failed, "error": constants.ErrData})
		return
	}

	result, httpCode, err := u.userUsecase.AuthUser(params)
	if err != nil {
		c.JSON(httpCode, gin.H{"response": constants.Failed, "error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"response": constants.Success, "data": result})
}
