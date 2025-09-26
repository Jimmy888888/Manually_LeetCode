/**
 * @param {number} n
 * @return {boolean}
 */
var divisorGame = function(n) {
    if (n === 1){
        return false;
    }

    // make dp
    let dp = new Array(n+1).fill(false);

    // dp[i] means must win or loss at i
    // dp[1] = false

    // go through, Alice go first
    for(let i = 2; i <= n; i++){
        for(let j = 1; j < i; j++){
            if (i % j === 0){
                if (dp[i-j] === false){
                    dp[i] = true;
                    break;
                }
            }
        }
    }

    return dp[n];
};