package cmd

import (
	"context"

	"github.com/imthaghost/bookrating/config"
	"github.com/imthaghost/bookrating/internal/notes/notion"
	"github.com/imthaghost/bookrating/internal/reviews"

	"github.com/k0kubun/pp"
)

func generateReview(ctx context.Context, args []string) error {
	csvPath := args[0]

	// config service
	configService := config.New{}
	configService.Load()
	cfg := configService.Get()

	// database ID
	databaseID := cfg.Notion.DatabaseID

	// init Notion client
	notionClient := notion.New(cfg)

	// create book reviews
	reviewedBooks, err := reviews.CreateBookClubReview(csvPath)
	if err != nil {
		return err
	}

	pp.Println("Loading books in Notion...")
	// cross-check with Notion to see if a review already exists
	for _, bookReview := range reviewedBooks {
		// if the book already exists we don't really have to do anything
		_, exist := notionClient.BookExists(ctx, bookReview.Title, databaseID)
		if exist {
			// nothing
		} else {
			// insert
			var favorites float64
			favorites = float64(bookReview.Favorites)
			err := notionClient.InsertReview(ctx, databaseID, bookReview.Title, &bookReview.AverageRating, &favorites)
			if err != nil {
				return err
			}
		}
	}
	return nil
}
