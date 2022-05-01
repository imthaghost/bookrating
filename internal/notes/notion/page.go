package notion

import (
	"context"
	"github.com/dstotijn/go-notion"
)

// BookExists will determine if a reviewed book already exists in a given Notion database
func (s *Service) BookExists(ctx context.Context, bookTitle string, databaseID string) (string, bool) {
	// filter by book title
	filter := notion.DatabaseQuery{
		Filter: &notion.DatabaseQueryFilter{
			Property: "Book Title",
			Text: &notion.TextDatabaseQueryFilter{
				Contains: bookTitle,
			},
		},
	}
	// results
	result, err := s.Client.QueryDatabase(ctx, databaseID, &filter)
	if err != nil {
		return "", false
	}
	if len(result.Results) > 0 {
		return result.Results[0].ID, true
	}

	return "", false
}

// InsertReview will insert a review into a given Notion database
func (s *Service) InsertReview(ctx context.Context, databaseID string, bookTitle string, rating *float64, favorites *float64) error {
	// page properties ( book title, rating, favorites)
	pageParam := notion.CreatePageParams{
		ParentType: notion.ParentTypeDatabase,
		ParentID:   databaseID,
		DatabasePageProperties: &notion.DatabasePageProperties{

			"Book Title": notion.DatabasePageProperty{
				Title: []notion.RichText{
					{
						Type: notion.RichTextTypeText,
						Text: &notion.Text{
							Content: bookTitle,
						},
					},
				},
			},

			"Rating": notion.DatabasePageProperty{
				Number: rating,
			},

			"Favorites": notion.DatabasePageProperty{
				Number: favorites,
			},
		},
	}

	// insert review into Notion
	_, err := s.Client.CreatePage(ctx, pageParam)
	if err != nil {
		return err
	}

	return nil
}
