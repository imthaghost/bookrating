package notes

import (
	"context"
)

type Service interface {
	Setup()
	InsertReview(ctx context.Context, databaseID string, bookTitle string, rating *float64, favorites *float64) error
	UpdateReview(context.Context, string) error
	BookExists(context.Context, string, string) (string, bool)
}
