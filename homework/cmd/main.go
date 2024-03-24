package main

import (
	"flag"
	"fmt"
	"homework/internal/adapters/hashmap"
	"homework/internal/ports/http"
	"os"

	"github.com/dubter/config"
	"homework/internal/app"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "config_path", "", "path to yaml config for server settings")

	flag.Parse()

	if configPath == "" {
		flag.Usage()
		os.Exit(1)
	}

	yaml, err := config.LoadConfig(configPath)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	hash := hashmap.NewHash()
	deviceService := app.NewService(hash)
	handler := http.NewHandler(
		&http.Config{
			Service:      deviceService,
			Port:         yaml.Port,
			Host:         yaml.Host,
			ReadTimeout:  yaml.ReadTimeout,
			WriteTimeout: yaml.WriteTimeout,
		})

	server := handler.NewServer()

	if err = server.ListenAndServe(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
