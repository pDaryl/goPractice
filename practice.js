function reverse(array, start, end) {
 while (start < end) {
   var temp = array[start]
   array[start] = array[end]
   array[end] = temp
    start++
    end--
 }
}

function rotateArrayInPlace(array, k){
    const len = array.length;

    if (len === 0 || k === 0) return array;

    var steps = k % len;

    if (steps === 0) return array; 

    reverse(array, 0, len-1); // reverse entire array
    reverse(array, 0, k-1); // grab first elements
    reverse(array, k, len-1); // grab the remaining elements 

    return array;
}

const array = [1, 2, 3, 4, 5];
const k = 2

console.log(rotateArrayInPlace(array, k));



