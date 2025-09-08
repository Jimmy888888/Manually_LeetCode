/**
 * @param {number[]} coins
 * @param {number} amount
 * @return {number}
 */
var coinChange = function(coins, amount) {
    // assume dp[i] is the number that can formed i , 0 <= i <= amount
    // dp[0] = 0;
    // dp[i] = dp[i - coin] + 1, if dp[i - coin] > -1

    // make array
    let dp = new Array(amount+1).fill(Number.MAX_SAFE_INTEGER);
    dp[0] = 0;

    for (let i = 1; i <= amount; i++){
        for (const c of coins){
            if (i < c){
                continue;
            }
            if (dp[i - c] < Number.MAX_SAFE_INTEGER){
                if (dp[i] > dp[i-c] + 1){
                    dp[i] = dp[i-c] + 1;
                }
            }
        }
    }

    if (dp[amount] === Number.MAX_SAFE_INTEGER){
        return -1;
    }

    return dp[amount];
};


