package server

import (
	"github.com/tusfendi/sales-go/cmd/api/middleware"
	"github.com/tusfendi/sales-go/internal/delivery"
)

func (s *Server) GroupRouter() {
	UserDelivery := delivery.NewUserDelivery(s.uc.User())
	userGroup := s.r.Group("/user")
	UserDelivery.Mount(userGroup)

	SalesDelivery := delivery.NewSalesDelivery(s.uc.Sales())
	salesGroup := s.r.Group("/sales")
	salesGroup.Use(middleware.JwtAuthMiddleware(s.cfg.SecretKey))
	SalesDelivery.Mount(salesGroup)
}
