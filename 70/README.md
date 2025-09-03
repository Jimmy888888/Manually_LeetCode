# DP

## State Definition

assumme dp[i] = the number of ways to reach i

at i :

if use 1 -> last step is dp[i-1]

if use 2 -> last step is dp[i-2]

and the ways to reach i = dp[i-1] + dp[i-2]

so dp[i] = dp[i-1] + dp[i-2]

## State Transition Equation

dp[i] = dp[i-1] + dp[i-2]

## base case
dp[0] = 1

dp[1] = 2



```js
/**
 * @param {number} n
 * @return {number}
 */
var climbStairs = function(n) {
    let dp = new Array(n).fill(0);

    // base cases
    dp[0] = 1;
    dp[1] = 2;
    //state transition equation
    for (let i = 2; i < n; i++){
        dp[i] = dp[i-1] + dp[i-2];
    }

    return dp[n-1];
};
```