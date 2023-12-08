// // usage example:
// var a = ['a', 1, 'a', 2, '1'];
// var unique = a.filter(onlyUnique);

// console.log(unique); // ['a', 1, 2, '1']

import fs from 'fs';

fs.readFile("./nums-js.txt", (e, data) => {
  const nums = data.toString().split(" ")


for (let i=nums.length-1; i > 0; i--) {
    if(i <= nums.length && nums[i] === nums[i - 1]){
    nums.splice(i,1);
  }
}

console.log(nums.join(" "));  
})



