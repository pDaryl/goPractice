class Movies {
    constructor(title, director, duration, watched) {
        this.title = title;
        this.director = director;
        this.duration = duration;
        this.watched = watched;
    }

    toggleWatched(){
        this.watched = !this.watched;
    }
}

function findLongestMovie(movies){
    var longestMovie = null;
    var mostTime = 0;

    for(const m of movies){
        if(m.duration > mostTime){
            mostTime = m.duration;
            longestMovie = m;
        }
    }
    return longestMovie;
}

function countMoviesWatched(movies){
    var moviesWatched = 0;

    for(const m of movies){
        if(m.watched === true){
            moviesWatched++;
        }
    }
    return moviesWatched;
}

const movies = [
    new Movies("Go Go Go", "D man", 97, false), 
    new Movies("504 Hoods", "BlackMan Walker",  110, true), 
    new Movies("Whoa Nelly", "HorseMan Jack", 78, true),
];

movies[0].toggleWatched();
movies[1].toggleWatched();


console.log("All Movies:\n", movies);

var movieCount = countMoviesWatched(movies);
console.log("Number of movies watched:\n", movieCount);

