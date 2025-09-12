# 72

## DP

### State definition

assume dp[i][j] = the minimum step that word1[i] can convert to word2[j]

### base cases
dp[0][0] = 0

dp[i][0] = len(word1) ; delete all word1 range i to 0

dp[0][j] = len(word2) ; word1 is 0, so insert all to become word2

### State transition equation

1. if word1[i] == word2[j]

dp[i][j] == dp[i-1][j-1]

2. if word1[i] !== word2[j]

 if insert:

dp[i][j] = dp[i-1][j] + 1

 if delete:

dp[i][j] = dp[i][j-1] + 1

 if replace:

dp[i][j] = dp[i-1][j-1] + 1

```js
/**
 * @param {string} word1
 * @param {string} word2
 * @return {number}
 */
var minDistance = function(word1, word2) {
    // make dp
    const n1 = word1.length;
    const n2 = word2.length;
    let dp = new Array(n1+1).fill(null).map(() => new Array(n2+1).fill(0));
    // let dp = Array.from({ length: n1 + 1 }, () => new Uint32Array(n2 + 1));

    // base cases
    for (let i = 0; i <= n1; i++){
        dp[i][0] = i;
    }
    for (let j = 0; j <= n2; j++){
        dp[0][j] = j;
    }

    // transition equation
    for (let i = 1; i <= n1; i++){
        for (let j = 1; j <= n2; j++){
            if (word1[i-1] === word2[j-1]){
                dp[i][j] = dp[i-1][j-1];
            }else{
                dp[i][j] = Math.min(
                    dp[i-1][j],
                    dp[i][j-1],
                    dp[i-1][j-1]
                ) + 1;
            }
        }
    }

    return dp[n1][n2];

};
```
