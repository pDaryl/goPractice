class Songs {
    constructor(title, artist, duration, isFavorite) {
        this.title = title;
        this.artist = artist;
        this.duration = duration;
        this.isFavorite = isFavorite;
    }
    toggleFavorite(){
        this.isFavorite = !this.isFavorite
    }
}

class Playlist {
    constructor() {
        this.songs = [];
    }

    addSong(song){
        this.songs.push(song);
    }
}

function displayPlaylist(playlist){
    for(const song of playlist.songs){
        console.log(`Title: ${song.title}, Artist: ${song.artist}, Duration: ${song.duration}, Fav Song: ${song.isFavorite}`);
    }
}

function findMostFrequentArtist(playlist){
    var artistCount = new Map();
    var mostSeenArtist = null;
    var freqCount = 0;

    for(const song of playlist.songs){
        var artist = song.artist
        artistCount.set(artist, (artistCount.get(artist) || 0) + 1)
    }
 
    for(const [artist, count] of artistCount.entries()){
        if(count > freqCount){
            freqCount = count;
            mostSeenArtist = artist;
        }
    }
    return mostSeenArtist;
}

function removeSongByTitle(playlist, title){
   var newPlaylist = playlist.songs.filter(song => song.title !== title);
   return newPlaylist;
}

function findLongestSong(songs){
    var longestSong = null;
    var songDur = 0;

    for(const s of songs){
        if(s.duration > songDur){
            songDur = s.duration;
            longestSong = s;
        }
    }
    return longestSong;
}

function favSongCount(songs){
   var favCount = 0;

   for(const s of songs){
    if(s.isFavorite === true){
        favCount++;
    }
   }
   return favCount;
}

function totalPlaylistDuration(songs){
    var totalDur = 0;

    for(const s of songs){
        totalDur += s.duration;
    }
    return totalDur;
}

function findSongsByArtist(songs, artist){
    var songsByArtist = [];
    
    for(const s of songs){
        if(artist === s.artist){
            songsByArtist.push(s);
        }
    }
    return songsByArtist;
}

const songs = [
    new Songs("New Slaves", "Kanye West", 4, true),
    new Songs("Child's Play", "Drake", 3, true), 
    new Songs("One Man Can Change The World", "Big Sean", 4, false),
    new Songs("Crooked Smile", "J. Cole", 5, true),
    new Songs("Take Care", "Drake", 4, true),
]

const playlist = new Playlist();

const song1 = new Songs("Crocodile Tears", "J. Cole", 4, false);
playlist.addSong(song1);
playlist.addSong(songs[0]);
playlist.addSong(songs[1]);
playlist.addSong(songs[2]);
playlist.addSong(songs[3]);
playlist.addSong(songs[4]);

songs[0].toggleFavorite();
const mostSeenArtist = findMostFrequentArtist(playlist); 

console.log("All Songs:\n");
console.log(songs);
console.log("Longest song:\n", findLongestSong(songs));
console.log("Number of fav songs:\n", favSongCount(songs));
console.log("Total duration of songs:\n", totalPlaylistDuration(songs));
console.log("Songs by Artists:\n", findSongsByArtist(songs, "Drake"));

console.log("the playlist:\n")
displayPlaylist(playlist);
console.log("most seen artist:\n")
console.log(mostSeenArtist);


var newPlaylist = removeSongByTitle(playlist, "Take Care");
console.log("new playlist:\n");
console.log(newPlaylist);





