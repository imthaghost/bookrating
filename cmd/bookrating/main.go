package main

import (
	"context"
	"github.com/imthaghost/bookrating/config"
	"github.com/imthaghost/bookrating/internal/notes/notion"
	"github.com/imthaghost/bookrating/internal/reviews"
	"github.com/k0kubun/pp"
	"log"
)

func main() {
	csvPath := "data/ratings.csv"
	// config service
	configService := config.New{}
	configService.Load()
	cfg := configService.Get()

	// database ID
	databaseID := cfg.Notion.DatabaseID

	// init Notion client
	notionClient := notion.New(cfg)

	// create book reviews
	reviewedBooks , err := reviews.CreateBookClubReview(csvPath)
	if err != nil {
		log.Fatal(err)
	}

	// cross-check with Notion to see if a review already exists
	pp.Println("Loading books in Notion...")
	for _, bookReview := range reviewedBooks {

		_, exist := notionClient.BookExists(context.Background(), bookReview.Title, databaseID)
		if exist {
			// update or do nothing :) technically
			} else {
				// insert
				var favorites float64
				favorites = float64(bookReview.Favorites)
				err := notionClient.InsertReview(context.Background(), databaseID, bookReview.Title, &bookReview.AverageRating, &favorites)
				if err != nil {
					log.Fatal(err)
				}
			}
		}
	}