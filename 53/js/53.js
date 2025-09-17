/**
 * @param {number[]} nums
 * @return {number}
 */
var maxSubArray = function(nums) {
    let n = nums.length;
    if (n === 1) {
        return nums[0];
    }
    
    let curS = nums[0];
    let golS = nums[0];

    for (let i = 1; i < n; i++){
        let num = nums[i];
        if (curS + num < num) {
            curS = num;
        }else{
            curS += num;
        }

        if (curS > golS) {
            golS = curS;
        }
    }

    return golS;
};