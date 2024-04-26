package server

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/tusfendi/sales-go/config"
	"github.com/tusfendi/sales-go/internal/delivery"
)

const projectDirName = "sales-go"

type Server struct {
	r    *gin.Engine
	cfg  config.Config
	uc   UsecaseRegistry
	repo RepositoryRegistry
}

func LoadEnv() {
	projectName := regexp.MustCompile(`^(.*` + projectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	err := godotenv.Load(string(rootPath) + `/.env`)
	if err != nil {
		panic(".env is not loaded properly")
	}
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

func (s *Server) GroupRouter() {
	UserDelivery := delivery.NewUserDelivery(s.uc.User())
	userGroup := s.r.Group("/user")
	UserDelivery.Mount(userGroup)
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

	repo := NewRepositoryRegistry(cfg, mySqlCon)
	us := NewUsecaseRegistry(repo)

	return &Server{
		r:    r,
		cfg:  *cfg,
		uc:   us,
		repo: repo,
	}
}
