package main

import "fmt"


type Movies struct {
	Title string
	Director string
	Duration int 
	Watched bool
}

func (m *Movies) toggleWatched(){
	m.Watched = !m.Watched
}

func findLongestMovie(movies []Movies) Movies {
	maxDuration := 0
	var longestMovie Movies

	for _, m := range movies {
		if m.Duration > maxDuration {
			maxDuration = m.Duration
			longestMovie = m 
		}
	}
	return longestMovie
}

func countMoviesWatched(movies []Movies) int{
	watchedMovies := 0
	
for _, m := range movies {
	if m.Watched == true {
		watchedMovies++
	}
}
	return watchedMovies
}

func main(){

movies := []Movies{
	{
		Title:    "Go Go Go",
		Director: "D man",
		Duration: 97,
		Watched:  false,
	}, 

	{
		Title:    "504 Hoods",
		Director: "BlackMan Walker",
		Duration: 110,
		Watched:  true,
	}, 

	{
		Title:    "Whoa Nelly",
		Director: "HorseMan Jack",
		Duration: 78,
		Watched:  true,
	}, 
}

movies[0].toggleWatched()
longestMovie := findLongestMovie(movies)
fmt.Println("Movie with longest duration:")
fmt.Println(longestMovie)
fmt.Println("Number of Movies watched:")
fmt.Println(countMoviesWatched(movies))

}




