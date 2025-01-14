/*function findSum(array, target){
if (array.length === 0) return [];
 
let map = new Map();

for (let i = 0; i < array.length; i++){
    const num = array[i];
    let complement = target - num;

    if (map.has(complement)){
        return [complement, num];
    } 
    map.set(num, i);
}
    return "there are no pairs";
}

console.log(findSum([1, 2, 3, 4, 5, 6], 3));
*/

const array = [1, 2, 3, 4];

for (let i = 0; i < array.length; i++){
    const num = array[i];

  // console.log(num);
   console.log(i);
}


