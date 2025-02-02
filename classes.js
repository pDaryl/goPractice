class Movies {
    constructor(title, director, duration, isfavorite) {
        this.title = title;
        this.director = director;
        this.duration = duration;
        this.isfavorite = isfavorite;
    }
    toggleFav(){
        this.isfavorite = !this.isfavorite;
    }
}

class MovieCollection {
    constructor() {
        this.movies = [];
    }

    addMovie(newMovie){
        this.movies.push(newMovie);
    }

    removeMovieByTitle(title){
     this.movies = this.movies.filter(movie => movie.title !== title);
    }
}

function displayMovies(collection){
    for(const movie of collection.movies){
        console.log(`Title: ${movie.title}, Director: ${movie.director}, Duration: ${movie.duration}minutes, Favorite: ${movie.isfavorite}`);
    };
}

function findLongestMovie(collection){
return collection.movies.reduce((longest, movie) => movie.duration > longest.duration ? movie : longest, collection.movies[0]);
}

function favMovieCount(collection){
  return collection.movies.filter(movie => movie.isfavorite).length;
}

function totalMoviesDuration(collection){
  return collection.movies.reduce((total, movie) => total + movie.duration, 0);
}

function findMoviesByDirector(collection, director){
return collection.movies.filter(movie => director === movie.director) ;
}

const movies = [
    new Movies("Ho Ho Hoe", "Naughty Santa", 90, false), 
    new Movies("The Runaway", "D. P.", 110, false),
    new Movies("Hold Those Horses!", "H. H Manny", 72, false), 
    new Movies("Don't Leave", "F Boy", 85, false),
]

var movieCollection = new MovieCollection();

const movie1 = new Movies("Through the Trenches", "T. Smith", 90, false);

movieCollection.addMovie(movie1);
for(const movie of movies){
    movieCollection.addMovie(movie);
};


movieCollection.movies[0].toggleFav();

//displayMovies(movieCollection);

movieCollection.removeMovieByTitle("Hold Those Horses!");

console.log("Movie Collection:");
displayMovies(movieCollection);

console.log("\n");
console.log("Longest movie in collection:");
console.log(findLongestMovie(movieCollection));

console.log("\n");
console.log("Number of fav movies:")
console.log(favMovieCount(movieCollection));

console.log("\n");
const totalDur = totalMoviesDuration(movieCollection);
console.log("Total duration of movies:")
console.log(`${totalDur} minutes`);

console.log("\n");
console.log("Movie found by director name:")
console.log(findMoviesByDirector(movieCollection, "D. P."));





