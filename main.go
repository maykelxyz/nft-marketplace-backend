package main

import (
	"log"
	"nft-marketplace-backend/src/cmd"
	"nft-marketplace-backend/src/config"
)

func main() {
	cfg, err := config.InitConfig(".env")
	if err != nil {
		log.Fatalf("Failed to initialize config: %v", err)
	}
	cmd.InitializeServer(cfg)
}
