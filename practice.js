function countFrequencies (array) {
    let freqCount = new Map();

    for (const num of array){
        if(!freqCount.has(num)){
            freqCount.set(num, 1);
        }else {
            freqCount.set(num, freqCount.get(num) + 1);
        }
    }

    return freqCount;
}

const array = [1, 2, 3, 2, 1, 4, 2];

console.log(countFrequencies(array));

