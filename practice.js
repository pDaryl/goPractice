function lengthSubArrayTargetSum(array, k){
  if (array.length === 0) return array.length;

  var maxLength = 0;
  var currentSum = 0;
  var start = 0;

  for (let end = 0; end < array.length; end++){
    currentSum += array[end];

    while(currentSum > k) {
      currentSum -= array[start];
      start++;
    }

    if (end - start + 1 > maxLength){
      maxLength = end - start + 1;
    }

  }
return maxLength;
}


const array = [1, 1, 1, 2, 2, 2, 3, 4, 5, 6];
const k = 6;

console.log(lengthSubArrayTargetSum(array, k));
