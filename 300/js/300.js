/**
 * @param {number[]} nums
 * @return {number}
 */
var lengthOfLIS = function(nums) {
    let n = nums.length;
    let dp = new Array(n).fill(1);

    for (let i = 1; i < n; i++){
        let curLong = dp[i];
        for (let j = i-1; j >= 0; j--){
            if (nums[i] > nums[j]){
                if (curLong < dp[j] + 1){
                    curLong = dp[j] + 1;
                }
            }
        }
        dp[i] = curLong;
    }

    let Longest = 1;
    for (const Long of dp) {
        if (Longest < Long){
            Longest = Long;
        }
    }

    console.log(dp);
    return Longest;
};