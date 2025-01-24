class Point {
    X;
    Y;

    constructor(X, Y) {
        this.X = X;
        this.Y = Y;
    }
}

function createPoint (X, Y){
    var point = new Point(X, Y);
    return point;
}

function printPoint (point) {
    console.log(`Point(X:${point.X}, Y:${point.Y})`);
}

const p = createPoint(3.5, 5.3);
printPoint(p)


