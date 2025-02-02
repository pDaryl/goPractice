package main

import "fmt"

type Movies struct {
	Title string
	Director string 
	Duration int
	isFavorite bool
}

type MovieCollection struct{
	Movies []Movies
}

func DisplayMovies(movies []Movies){
	for _, m := range movies {
		fmt.Printf("Title: %s, Director: %s, Duration: %d, Favorites: %v\n", m.Title, m.Director, m.Duration, m.isFavorite)
	}
}

func (m *Movies) toggleFav(){
	m.isFavorite = !m.isFavorite
}

func (mc *MovieCollection) addMovie(NewMovie Movies){
	mc.Movies = append(mc.Movies, NewMovie)
}

func (mc *MovieCollection) removeMovieByTitle(title string) []Movies {
var editedMovies []Movies

for _, movie := range mc.Movies {
if movie.Title != title {
	editedMovies = append(editedMovies, movie)
}
}
return editedMovies
}

func findLongestDuration(collection []Movies) Movies {
var longestMovie Movies
maxDur := 0

for _, c := range collection{
	if c.Duration > maxDur {
		maxDur = c.Duration
		longestMovie = c
	}
}
return longestMovie
}

func favMovieCount(collection []Movies) int {
	favMovies := 0

	for _, c := range collection{
		if c.isFavorite == true {
			favMovies++
		}
	}
	return favMovies
}

func findMoviesByDirector(collection []Movies, director string) []Movies {
var movieByDirector []Movies

for _, c := range collection {
	if c.Director == director {
		movieByDirector = append(movieByDirector, c)
	}
}
return movieByDirector
}

func totalMoviesDuration(collection []Movies) int {
	totalDur := 0

	for _, c := range collection {
		totalDur += c.Duration
	}
	return totalDur
}

func main(){

	movies := []Movies{
		{Title:      "The Runaway",
		Director:   "D. P.",
		Duration:   110,
		isFavorite: false,
	}, 
	{
		Title:      "Hold Those Horses!",
		Director:   "H. H. Manny",
		Duration:   72,
		isFavorite: false,
	}, 
	{
		Title:      "Don't Leave",
		Director:   "F Boy",
		Duration:   85,
		isFavorite: false,
	},
}

movieCollection := MovieCollection{}

movies[0].toggleFav()

movie1 := Movies{
	Title:      "Ho Ho Hoe",
	Director:   "Naughty Santa",
	Duration:   90,
	isFavorite: false,
}

movieCollection.addMovie(movie1)
movieCollection.addMovie(movies[0])
movieCollection.addMovie(movies[1])
movieCollection.addMovie(movies[2])

//DisplayMovies(movies)
fmt.Println("All movies in collectin:")
DisplayMovies(movieCollection.Movies)

fmt.Println()
fmt.Println("All Movies AFTER removing:")
movieCollection.Movies = movieCollection.removeMovieByTitle("Hold Those Horses!")
DisplayMovies(movieCollection.Movies)

fmt.Println()
longestMovie := findLongestDuration(movieCollection.Movies)
fmt.Printf("the longest movie: %v\n", longestMovie)

fmt.Println()
fmt.Println("Number of favorite movies in collection:", favMovieCount(movieCollection.Movies))

fmt.Println()
moviesByDirector := findMoviesByDirector(movieCollection.Movies, "D. P.")
fmt.Println("Movies by Director name:\n", moviesByDirector)

fmt.Println()
fmt.Printf("Total Duration of all movies is %dmins:\n", totalMoviesDuration(movieCollection.Movies))
}