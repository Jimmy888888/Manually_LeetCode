# 198

## State Definition

1. dp[i] = the biggest sum of point i in nums

2. can't connect

## State transition Equation

at i:

assume take i -> dp[i] = dp[i-2] + nums[i]

assume didn't take i -> dp[i] = dp[i-1]

--> dp[i] = max((dp[i-2] + nums[i]), dp[i-1])

## base cases
dp[0] = nums[0]

dp[1] = max(nums[0], nums[1])

```js
/**
 * @param {number[]} nums
 * @return {number}
 */
var rob = function(nums) {
    // check len
    let n = nums.length;
    if (n === 1) {
        return nums[0];
    }
    if (n === 2) {
        return Math.max(nums[0], nums[1]);
    }

    // make dp
    let dp = new Array(n).fill(0);

    // base cases
    dp[0] = nums[0];
    dp[1] = Math.max(nums[0], nums[1]);

    // state transition equation
    for (let i = 2; i < n; i++){
        dp[i] = Math.max((dp[i-2] + nums[i]), dp[i-1]);
    }

    return dp[n-1];
};
```
