package delivery

import (
	"github.com/gin-gonic/gin"
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

func (c *UserDelivery) Mount(group *gin.RouterGroup) {

	group.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	group.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "test",
		})
	})
}
