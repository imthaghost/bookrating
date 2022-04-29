package config

// Service is an interface that defines the functions needed to implement a Config Service.
type Service interface {
// Load will do any config setup (like load env vars)
Load()
// Get will get the config
Get() Config
}

// Config is a service that is designed to provide various configuration to the rest of the application.
type Config struct {
	General	GeneralConfig
	Notion 	NotionConfig
}

// GeneralConfig contains general information the application
type GeneralConfig struct {
	AppEnv	string // the environment that the application is running in (dev, prod, etc)
}

// NotionConfig contains information that the Notion API needs to run
type NotionConfig struct {
	IntegrationToken	string 	// integration token
	DatabaseID 			string	// database ID
}