# 300

## DP

### State definition:
assume dp[i] is the longest invreasing subsquence if nums[i] is 

the biggest nums in sub

if nums[i] > nums[j], dp[i] = dp[j] + 1, else dp[i] = 1

long chain and a short chain inside  5 6 7 8 9 10 1 2 3 11 12 13

```js
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
```




