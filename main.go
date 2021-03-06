package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/DonilZ/moviefan-rest-service/config"
	"github.com/DonilZ/moviefan-rest-service/daemon"

	_ "github.com/DonilZ/moviefan-rest-service/docs"
)

// @title Moviefan Swagger API
// @version 1.0
// @description Swagger API for Golang Project Moviefan
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.email info.donilz@gmail.com

var assetsPath string

func processFlags() *daemon.Config {
	cfg := &daemon.Config{}

	flag.StringVar(&cfg.ListenSpec, "listen", config.GetListenAddress(), "HTTP listen spec")

	flag.StringVar(&cfg.Db.ConnectString, "db-connect",
		fmt.Sprintf("host=%s dbname=%s sslmode=disable", config.GetDbHost(), config.GetDbName()),
		"DB Connect String")

	flag.StringVar(&assetsPath, "assets-path", "assets", "Path to assets dir")

	flag.Parse()
	return cfg
}

func setupHTTPAssets(cfg *daemon.Config) {
	log.Printf("Assets served from %q.", assetsPath)
	cfg.UI.Assets = http.Dir(assetsPath)
}

// @BasePath /api/v1
func main() {
	cfg := processFlags()

	setupHTTPAssets(cfg)

	if err := daemon.Run(cfg); err != nil {
		log.Printf("Error in main(): %v", err)
	}
}
