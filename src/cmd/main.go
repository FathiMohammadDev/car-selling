package main

import (
	"github.com/FathiMohammadDev/car-selling/config"
	"github.com/FathiMohammadDev/car-selling/data/cache"
)

func main() {
	cfg := config.GetConfig()
	defer cache.CloseRedis()
	cache.InitRedis(cfg)
}
