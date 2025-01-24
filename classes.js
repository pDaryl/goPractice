class Circle {
    radius;
    color;
    constructor(radius, color) {
        this.radius = radius;
        this.color = color;
    }
}

function createCircle(radius, color) {
    var circle = new Circle(radius, color);
    return circle;
}

function printCircle(circle) {
    console.log(`Circle(Radius:${circle.radius}, Color:${circle.color})`);
}

const circle = createCircle(7.25, "green");
printCircle(circle);


