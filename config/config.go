// Copyright 2022 imthaghost. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

const (
	DEV  = "dev"
	PROD = "prod"
)

// New ...
type New struct{}

// Load will load the config
func (n *New) Load() {
	err := godotenv.Load()
	if err != nil {
		log.Println("did not load env vars from .env")
		log.Println(err)
	}
}

// Get ...
func (n *New) Get() Config {

	return Config{
		General: getGeneralConfig(),
		Notion:  getNotionConfig(),
	}
}

// getGeneralConfig ....
func getGeneralConfig() GeneralConfig {
	// dev -- default
	config := GeneralConfig{
		AppEnv: os.Getenv("APP_ENV"),
	}

	return config
}

// getNotionConfig ....
func getNotionConfig() NotionConfig {
	return NotionConfig{
		IntegrationToken: os.Getenv("NOTION_INTEGRATION_KEY"),
		DatabaseID:       os.Getenv("NOTION_DATABASE_ID"),
	}
}
