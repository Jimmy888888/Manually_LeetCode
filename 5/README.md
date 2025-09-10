# 5

## DP

### State Definition

how to check a palindromic :

-> two flag(f & b) go through forward and backward, and the two 

flag must be the same always, 

-> until f >= b

### assume :
dp[l][r] means s[l...r] is a palindromic

-> dp[l][r] = ( s[l] === s[r] ) && (dp[l-1][r-1])

### how to go through :

-> start from the 0 point, and increase untill end point


```js
/**
 * @param {string} s
 * @return {string}
 */
var longestPalindrome = function(s) {
    // check
    let n = s.length;
    if (n === 1){
        return s;
    }

    // make dp
    let dp = new Array(n).fill(null).map(() => Array(n).fill(false));
    // record max
    let maxL = 0;
    let maxSub = "";

    // state transation equation
    for (let i = 0; i < n; i++){
        for (let j = 0; j <= i; j++){
            if(s[i] === s[j]){
                if (i-j <= 2 || dp[i-1][j+1]){
                    dp[i][j] = true;
                }
                if (dp[i][j]){
                    if (i-j+1 > maxL){
                        maxL = i-j+1;
                        maxSub = s.substring(j, i+1);
                        console.log(maxSub);
                    }
                }

            }
        }
    }

    return maxSub;
};
```

#### result : pass, but spend too much time

use center expend method will use less time, and less space