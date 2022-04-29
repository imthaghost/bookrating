package notes

import (
	"context"
	"github.com/dstotijn/go-notion"
	"github.com/imthaghost/bookrating/internal/reviews"
)

type Service interface {
	Setup()
	GetDatabase(context.Context, string) (*notion.DatabaseQueryResponse, error)
	InsertReview(ctx context.Context, databaseID string, bookTitle string, rating *float64, favorites *float64) error
	UpdateReview(context.Context, *reviews.Review, string) error
	BookExists(context.Context, string, string) (string, bool)
}