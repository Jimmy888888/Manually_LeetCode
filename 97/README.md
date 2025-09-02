# DP 

## State definition:

build a 2D bool{(s1.length+1) * (s2.length+1)} array DP[i][j], init with 0

DP[i][j] = whether first i chrs in s1 and first j chrs in s2 can build first i+j chrs in s3

## State Equation:

DP[i][j] = (DP[i-1][j] && s1[i-1] === s3[i+j-1]) || (DP[i][j-1] && s2[j-1] === s3[i+j-1])

## Base cases:

DP[i][0] = 1

DP[0][j] = 1

finally check DP[s1.length][s2.length]

```js
/**
 * @param {string} s1
 * @param {string} s2
 * @param {string} s3
 * @return {boolean}
 */

var isInterleave = function(s1, s2, s3) {
    let n1 = s1.length;
    let n2 = s2.length;
    let n3 = s3.length;

    if (n1 + n2 !== n3){
        return false;
    }

    let DP = Array(n1 + 1).fill(null).map(() => Array(n2 + 1).fill(false));

    // base cases
    DP[0][0] = true;

    for (let i = 1; i < n1 + 1; i++){
        if (s1[i-1] === s3[i-1]){
            DP[i][0] = DP[i-1][0];
        }
    }

    for (let j = 1; j < n2 + 1; j++){
        if (s2[j-1] == s3[j-1]){
            DP[0][j] = DP[0][j-1];
        }
    }

    // state transition
    for (let i = 1; i < n1 + 1; i++){
        for (let j = 1; j < n2 + 1; j++){
            chr3 = s3[i+j-1];
            DP[i][j] = (DP[i-1][j] && s1[i-1] === chr3) || (DP[i][j-1] && s2[j-1] === chr3);
        }
    }

    console.log(DP);

    return DP[n1][n2];
};
```