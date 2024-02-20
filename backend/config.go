package main

import (
	"flag"
	"log"
)

type userServiceCfg struct {
	URL    string
	JWTKey string
}

type configData struct {
	port        int
	userService userServiceCfg
	devAuthUser string
	jwtKey      string
}

func getConfiguration() *configData {
	var config configData
	flag.IntVar(&config.port, "port", 8080, "Port to offer service on")
	flag.StringVar(&config.jwtKey, "jwtkey", "", "JWT signature key")

	// user service
	flag.StringVar(&config.userService.URL, "userws", "", "URL for the user service")
	flag.StringVar(&config.userService.JWTKey, "userjwtkey", "", "JWT for the user service")

	// dev user
	flag.StringVar(&config.devAuthUser, "devuser", "", "Authorized computing id for dev")

	flag.Parse()

	if config.jwtKey == "" {
		log.Fatal("Parameter jwtkey is required")
	}
	if config.userService.URL == "" {
		log.Fatal("Parameter userws is required")
	}
	if config.userService.JWTKey == "" {
		log.Fatal("Parameter userjwt is required")
	}

	log.Printf("[CONFIG] port          = [%d]", config.port)
	log.Printf("[CONFIG] userws        = [%s]", config.userService.URL)
	if config.devAuthUser != "" {
		log.Printf("[CONFIG] devuser       = [%s]", config.devAuthUser)
	}

	return &config
}
