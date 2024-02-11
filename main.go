package main

import (
	"fmt"
	"log"

	"github.com/mustafakemalgordesli/go-commerce/config"
	"github.com/mustafakemalgordesli/go-commerce/pkg/database"
	"github.com/mustafakemalgordesli/go-commerce/pkg/router"
	"github.com/spf13/viper"
)

func main() {

	if err := config.Setup(); err != nil {
		log.Fatalf("config.Setup() error: %s", err)
	}

	if err := database.Setup(); err != nil {
		log.Fatalf("database.Setup() error: %s", err)
		fmt.Println("database" + err.Error())
	}

	db := database.GetDB()
	r := router.Setup(db)

	host := "localhost"

	if h := viper.GetString("server.host"); h != "" {
		host = h
	}

	port := "3333"
	if p := viper.GetString("server.port"); p != "" {
		port = p
	}
	addr := host + ":" + port

	err := r.Run(addr)
	fmt.Println("Error:", err)
}
