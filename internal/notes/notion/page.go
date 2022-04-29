package notion

import (
	"context"
	"github.com/dstotijn/go-notion"
	"github.com/imthaghost/bookrating/internal/reviews"
)

// BookExists will determine if a review already exists in a given Notion database
func (s *Service) BookExists(ctx context.Context, bookTitle string, databaseID string) (string, bool) {

	filter := notion.DatabaseQuery{
		Filter: &notion.DatabaseQueryFilter{
			Property: "Book Title",
			Text: &notion.TextDatabaseQueryFilter{
				Contains: bookTitle,
			},
		},
	}

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

// UpdateReview ...
func (s *Service) UpdateReview(ctx context.Context, update *reviews.Review, pageID string) error {

	params := notion.UpdatePageParams {
		DatabasePageProperties: notion.DatabasePageProperties{
			"Book Title": notion.DatabasePageProperty{
				Title: []notion.RichText{
					{
						Type: notion.RichTextTypeText,
						Text: &notion.Text{
							Content: "nothing",
						},
					},
				},
			},
		},
	}

	_, err := s.Client.UpdatePage(ctx, pageID, params)
	if err != nil  {
		return err
	}

	return nil
}

