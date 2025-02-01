package main

import "fmt"

type Songs struct{
	Title string
	Artist string
	Duration int
	IsFavorite bool
}

type Playlist struct{
	Songs []Songs
}


func (s *Songs) isFav(){
	s.IsFavorite = !s.IsFavorite
}


func (p *Playlist) addSong(newSong Songs){
	p.Songs = append(p.Songs, newSong)
}

func findLongestSong(songs []Songs) Songs{
	var longestSong Songs
	maxDur := 0

	for _, s := range songs{
		if s.Duration > maxDur{
			maxDur = s.Duration
			longestSong = s
		}
	}
	return longestSong
}

func countFavSongs(songs []Songs) int {
	favSongsCount := 0

	for _, s := range songs{
		if s.IsFavorite == true{
			favSongsCount++
		} 
	}
	return favSongsCount
}

func totalPlaylistDuration(songs []Songs) int {
	totalDur := 0

	for _, s := range songs {
		totalDur += s.Duration
	}
	return totalDur
}

func findSongsByArtist(songs []Songs, Artist string) []Songs {
var songsByArtist []Songs

for _, s := range songs {
	if Artist == s.Artist {
		songsByArtist = append(songsByArtist, s)
	}
}
return songsByArtist
}

func findMostFrequentArtist(playlist []Songs ) string {
	artistCount := make(map[string] int)
	var freqArtist string
	var freqCount int

	for _, p := range playlist{
	artistCount[p.Artist]++
	}
	for artist, count := range artistCount{
		if count > freqCount{
			freqCount = count
			freqArtist = artist
		}
	}
	return freqArtist
}

func displayPlaylist(playlist []Songs) {
	for _, song := range playlist{
fmt.Printf("Title: %s, Artist: %s, Duration: %d, Favorite: %v\n", 
song.Title, 
song.Artist, 
song.Duration, 
song.IsFavorite )
	}
}

func removeSongByTitle(playlist []Songs, title string) []Songs {
	var editedPlaylist []Songs

	for _, song := range playlist{
		if song.Title != title{
			editedPlaylist = append(editedPlaylist, song)
		}
	}
	return editedPlaylist
}

func main(){

songs := []Songs {
	{
		Title:      "New Slaves",
		Artist:     "Kanye West",
		Duration:   4,
		IsFavorite: true,
	}, 
	{
		Title:      "Child's Play",
		Artist:     "Drake",
		Duration:   3,
		IsFavorite: true,
	}, 
	{
		Title:      "One Man Can Change The World",
		Artist:     "Big Sean",
		Duration:   4,
		IsFavorite: true,
	}, 
	{
		Title:      "Crooked Smile",
		Artist:     "J. Cole",
		Duration:   5,
		IsFavorite: true,
	},
	{
		Title:      "Take Care",
		Artist:     "Drake",
		Duration:   4,
		IsFavorite: true,
	},
}

playList := Playlist{}

newSong1 := Songs{
	Title:      "Crocodile Tears",
	Artist:     "J. Cole",
	Duration:   3,
	IsFavorite: false,
}

songs[0].isFav()
playList.addSong(newSong1)
playList.addSong(songs[0])
playList.addSong(songs[3])
playList.addSong(songs[1])
playList.addSong(songs[2])

mostListenedArtist := findMostFrequentArtist(playList.Songs)

fmt.Println("All songs:")
displayPlaylist(songs)

fmt.Println()
fmt.Println("Duration of all songs:")
songDur := totalPlaylistDuration(songs)
fmt.Printf("%d mins\n", songDur)

fmt.Println()
fmt.Println("Number of favorite Songs:")
fmt.Println(countFavSongs(songs))

fmt.Println()
fmt.Println("Longest song in playlist:")
fmt.Println(findLongestSong(songs))

fmt.Println()
fmt.Println("Return songs by Artist:")
fmt.Println(findSongsByArtist(songs, "Drake"))

fmt.Println()
fmt.Println("Songs in Playlist:")
displayPlaylist(playList.Songs)

fmt.Println()
fmt.Print("Most Listned to Artist in playlist:\n")
fmt.Println(mostListenedArtist)

fmt.Println()
fmt.Println("Edited Playlist:")
playList.Songs = removeSongByTitle(playList.Songs, "Crooked Smile")
displayPlaylist(playList.Songs)


}




