package main

import (
	"os"
	"regexp"

	"github.com/joho/godotenv"
	"github.com/tusfendi/sales-go/cmd/api/server"
)

func init() {
	loadEnv()
}

func main() {
	server.Start()
}

func loadEnv() {
	projectName := regexp.MustCompile(`^(.*` + server.ProjectDirName + `)`)
	currentWorkDirectory, _ := os.Getwd()
	rootPath := projectName.Find([]byte(currentWorkDirectory))

	err := godotenv.Load(string(rootPath) + `/.env`)
	if err != nil {
		panic(".env is not loaded properly")
	}
}
