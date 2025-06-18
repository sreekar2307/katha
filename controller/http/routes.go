package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sreekar2307/katha/config"
	"github.com/sreekar2307/katha/controller/http/middleware"
	"net/http"
)

func NewServer(
	conf config.Server,
	middleware middleware.Middleware,
	controller Controller,
) (*http.Server, error) {
	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(gin.Logger())

	public := r.Group("/public", middleware.UserAuthMiddleware())
	{
		v1 := public.Group("/v1")
		{
			v1.POST("/expenses", controller.V1.NewExpense())
			v1.GET("/expenses", controller.V1.Expenses())
			v1.GET("/balances", controller.V1.Balances())

		}
	}

	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", conf.Host, conf.Port),
		Handler: r,
	}
	return srv, nil
}
