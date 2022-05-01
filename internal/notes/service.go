// Copyright 2022 imthaghost. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package notes

import (
	"context"
)

type Service interface {
	Setup()
	InsertReview(ctx context.Context, databaseID string, bookTitle string, rating *float64, favorites *float64) error
	UpdateReview(context.Context, string) error
	BookExists(ctx context.Context, databaseID string, bookTitle string) (string, bool)
}
