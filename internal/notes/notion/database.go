package notion

import (
	"context"
	"github.com/dstotijn/go-notion"
)

// CreateDatabase will create a database in notion
func (s *Service) CreateDatabase(ctx context.Context) (*notion.Database, error) {

	return nil, nil
}

// ListDatabases will list all the databases from Notion
func ListDatabases(ctx context.Context) ([]*notion.Database, error) {
	return nil, nil
}

// DeleteDatabase will delete a single database from Notion
func DeleteDatabase(ctx context.Context, id string) error {
	return nil
}

// GetDatabase gets our database object from Notion
func (s *Service) GetDatabase(ctx context.Context, id string) (*notion.DatabaseQueryResponse, error) {
	db, err := s.Client.QueryDatabase(ctx, id, &notion.DatabaseQuery{})
	if err != nil  {

		return nil, err
	}

	return &db, nil
}

