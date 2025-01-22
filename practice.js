function lengthLongestSubArrayOfSum (array, k) {
  if (array.length === 0) return array;

  if (k === 0) return "target sum is 0";

  var maxLength = 0;
  var currentSum = 0;
  var start = 0;

  for (let end = 0; end < array.length; end++){
    currentSum += array[end];

    if (currentSum > k && start <= end){
      currentSum -= array[start];
      start++;
    }
    if (array.slice(start, end+1).length > maxLength){
      maxLength = array.slice(start, end+1).length;
    }
  }
return maxLength;
}

const array = [2, 1, 5, 2, 3, 2]
const k = 5

console.log(lengthLongestSubArrayOfSum(array, k));



