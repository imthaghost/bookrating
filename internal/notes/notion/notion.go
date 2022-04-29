package notion

import (
	"github.com/dstotijn/go-notion"
	"github.com/imthaghost/bookrating/config"
	"github.com/imthaghost/bookrating/internal/notes"
)

// Service ...
type Service struct {
	Client *notion.Client
	Config config.Config
}

// Setup performs any configuration for connecting to the Notion API
func (s *Service) Setup() {

}


// New creates a new notes service using the Notion API
func New(cfg config.Config) notes.Service {
	client := notion.NewClient(cfg.Notion.IntegrationToken)

	return &Service{
		Client: client,
		Config: cfg,
	}

}