class Books {
    constructor(title, author, pages, isCheckedOut) {
        this.title = title;
        this.author = author;
        this.pages = pages;
        this.isCheckedOut = isCheckedOut;
    }

    toggleCheckOut(){
        this.isCheckedOut = !this.isCheckedOut;
    }
}

function findLongestBook(books){
    var mostPages = 0;
    var longestBook = null;

    for(const b of books) {
        if (b.pages > mostPages){
            mostPages = b.pages;
            longestBook = b;
        }
    }
    return longestBook;
}

const books = [
    new Books("Join the War",
		"PP Wario",
		269,
		false,), 

    new Books(
      "Winner Winner Chicken Dinner",
	    "W.W Chic",
		500,
	    true,
    ), 
    new Books(
       "Here me, here ye!",
		"Ching Ming",
	    304,
		true,
    )
]

books[0].toggleCheckOut();
books[1].toggleCheckOut();

var longestBook = findLongestBook(books);
console.log(longestBook);
