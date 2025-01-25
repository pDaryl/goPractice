class Author {
    constructor(name, nationality) {
        this.name = name;
        this.nationality = nationality;
    }
}

class Books {
    constructor(title, pages, author) {
        this.title = title;
        this.pages = pages;
        this.author = author;
    }
}

function createBook(title, pages, name, nationality){
    const author = new Author(name, nationality);
    const book = new Books(title, pages, author);
    return book;
}

function printBook(book){
    console.log(`Book(Title: ${book.title}, Pages: ${book.pages}, Author: [Name: ${book.author.name}, Nationality: ${book.author.nationality}])`);
}


const b = createBook("Daryl's World", 369, "Daryl Peterson", "American");
printBook(b);
