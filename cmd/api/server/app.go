package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tusfendi/sales-go/cmd/api/registry"
	"github.com/tusfendi/sales-go/config"
)

const ProjectDirName = "sales-go"

type Server struct {
	r    *gin.Engine
	cfg  config.Config
	uc   registry.UsecaseRegistry
	repo registry.RepositoryRegistry
}

func Start() {
	server := initServer()

	server.r.GET("/foo", func(c *gin.Context) {
		fmt.Println("The URL: ", c.Request.Host+c.Request.URL.Path)
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"response": c.Request.Host})
	})

	server.GroupRouter()
	server.r.Run(":" + fmt.Sprint(server.cfg.ApiPort))
}

func initServer() *Server {
	cfg := config.NewConfig()
	mySqlCon, err := config.NewMysql(cfg.AppEnv, &cfg.MysqlOption)
	if err != nil {
		log.Fatal(err)
		println("error mysql")
	}

	r := gin.Default()
	r.SetTrustedProxies(nil)

	repo := registry.NewRepositoryRegistry(cfg, mySqlCon)
	us := registry.NewUsecaseRegistry(cfg, repo)

	return &Server{
		r:    r,
		cfg:  *cfg,
		uc:   us,
		repo: repo,
	}
}
