function findLengthOfLongestSubArrayOfKElements(array, k){
if (array.length === 0 || k > array.length) return array.length;

if (k === 0) return "number of distinct elements is 0", k;


var maxLength = 0;
var distinctNums = new Map();
var start = 0;

for (let end = 0; end < array.length; end++){
  var currentNum = array[end];
  if (!distinctNums.has(currentNum)){
    distinctNums.set(currentNum, 1)
  } else{
    distinctNums.set(currentNum, distinctNums.get(currentNum) +1); 
  }

  while(distinctNums.size > k){
    distinctNums.set(array[start], distinctNums.get(array[start]) - 1);
    if(distinctNums.get(array[start]) === 0){
      distinctNums.delete(array[start]);
    }
    start++;
  }
if (end - start + 1 > maxLength){
  maxLength = end - start + 1;
}
}
return maxLength;
}


const array = [1, 1, 1, 2, 2, 2, 3, 4, 5, 6];
const k = 3;

console.log(findLengthOfLongestSubArrayOfKElements(array, k));


