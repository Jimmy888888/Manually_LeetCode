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