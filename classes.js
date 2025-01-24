class Rectangle {
    length;
    width;
    constructor(length, width){
        this.length = length;
        this.width = width; 
    }
}

function createRectangle (length, width) {
    var rectangle = new Rectangle(length, width);
    return rectangle;
}

function printRectangle (rectangle){
    console.log(`Rectangle(Length:${rectangle.length}, Width:${rectangle.width})`);
}

const rectangle = createRectangle(4.5, 5.4);
printRectangle(rectangle);


