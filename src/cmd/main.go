package main

import (
	"log"

	"github.com/FathiMohammadDev/car-selling/api"
	"github.com/FathiMohammadDev/car-selling/config"
	"github.com/FathiMohammadDev/car-selling/data/cache"
	"github.com/FathiMohammadDev/car-selling/data/db"
)

func main() {
	cfg := config.GetConfig()

	defer cache.CloseRedis()
	err := cache.InitRedis(cfg)
	if err != nil {
		log.Fatal(err)
	}

	defer db.CloseDb()
	err = db.InitDb(cfg)
	if err != nil {
		log.Fatal(err)
	}

	api.InitServer(cfg)

}
