package reviews

import (
	"encoding/csv"
	"math"
	"os"
	"strconv"
	"strings"
)

// TODO Switch our bookratings map to a concurrent safe one, we can allocate a slice of strings to buffer a chunk of lines and then process them using a goroutine.
// We can try to find a good chunk-size for this. Every chunk would have its own go routine.
// We can then parse a line and send over to a collecting goroutine over 1 channel that will fill our map. This would reduce processing time by 40%

// CreateBookClubReview creates an object that contains all member reviews of a given book
func CreateBookClubReview(path string) ([]*ReviewedBook, error) {
	// open file
	csvData, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	// close the file at the end of the program
	defer csvData.Close()
	// read file line by line
	lines, err := csv.NewReader(csvData).ReadAll()
	if err != nil {
		return nil, err
	}

	// stores all books and member reviews
	reviews := make(map[string][]int)
	var bookReview []*ReviewedBook

	// create a histogram of books and append member reviews
	for _, line := range lines {
		// assign each line
		bookName := strings.TrimSpace(line[0])
		bookTitle := strings.ToLower(bookName)
		memberName := line[0]
		memberRating := line[2]
		rating, _ := strconv.ParseFloat(memberRating, 64)

		// if the book has been seen then we can just add the review into the map
		// since the keys are unique we are always going to assign the member last review
		if book, ok := reviews[bookTitle]; ok {
			book[memberName] = rating
		} else {
			reviews[bookTitle] = map[string]float64{memberName: rating}
		}
	}

	for bookName, _ := range reviews {
		sum := 0.0
		favorite := 0
		for memberName, _ := range reviews[bookName] {
			sum += reviews[bookName][memberName]

			if reviews[bookName][memberName] == 5 {
				favorite += 1
			}

		}
		average := sum / float64(len(reviews[bookName]))
		averageRating := roundFloat(average, 1)
		bookReview = append(bookReview, &ReviewedBook{
			Title:         bookName,
			Favorites:     favorite,
			AverageRating: averageRating,
		})
	}

	return bookReview, nil
}

func roundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
