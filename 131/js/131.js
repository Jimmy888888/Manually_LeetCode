/**
 * @param {string} s
 * @return {string[][]}
 */
var partition = function(s) {
    
    // dp[i][j] means s[i~j] is a palindrome or not
    let n = s.length;
    let dp = new Array(n).fill(null).map(() => Array(n).fill(false));

    // go through
    for (let i = n-1; i > -1; i--){
        for (let j = i; j < n; j++){
            if (s[i] === s[j]){
                dp[i][j] = ( j-i < 2 || dp[i+1][j-1]);
            }
        }
    }

    let result = [];
    let path = [];

    var backTrack = function(startId){
        if (startId >= n){
            result.push([...path]);
            return
        }

        for (let i = startId; i < n; i++){
            if(dp[startId][i] === true){
                let subS = s.slice(startId,i+1);
                path.push(subS);
                backTrack(i+1);
                path.pop();
            }
        }
    }

    backTrack(0);

    return result;

};