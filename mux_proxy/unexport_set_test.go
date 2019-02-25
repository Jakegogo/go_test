package sexpr

import (
	"fmt"
	utils "go_test/utils"
	"testing"
)

type Movie struct {
	Title, Subtitle string
	year            int
	Color           bool
	Actor           map[string]string
	Oscars          []string
	Sequel          *string
}

func TestSetUnexport(t *testing.T) {
	y := 1964
	var movie = Movie{
		Title:    "Dr. Strangelove",
		Subtitle: "How I Learned to Stop Worrying and Love the Bomb",
		year:     y,
		Color:    false,
	}

	//mf := &movie
	movie1 := utils.SetUnexportField(&movie, "year", 1990)

	fmt.Println("movie copy:", movie1.(*Movie).year)

	fmt.Println("movie orign:", movie.year)
}
