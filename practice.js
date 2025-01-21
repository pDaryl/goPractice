function findSumContinuousArray (array, k) {
  if (array.length === 0) return array; 

  if (k === 0) return "no target sum to find";

  var result = [];
  var start = 0;
  var currentSum = 0; 

  for (let reader = 0; reader < array.length; reader++){
    currentSum += array[reader];

   if (currentSum > k && start <= reader) {
    currentSum -= array[start];
    start++
   }
   if(currentSum === k) {
    var subArray = array.slice(start, reader + 1);
     result.push(subArray);
   }
  }
  return result;
}

const array = [1, 2, 3, 4, 5]; 
const k = 6;

console.log(findSumContinuousArray(array, k));



