package reviews



// Review represents a members review of a book
type Review struct {
	Member string
	Rating int
}

// ReviewedBook represents a complete review of a book
type ReviewedBook struct {
	Title string
	Favorites int
	AverageRating float64
}