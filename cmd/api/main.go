package main

import (
	"github.com/tusfendi/sales-go/cmd/api/server"
)

func init() {
	server.LoadEnv()
}

func main() {
	server.Start()
}
