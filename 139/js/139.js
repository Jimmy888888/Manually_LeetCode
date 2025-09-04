/**
 * @param {string} s
 * @param {string[]} wordDict
 * @return {boolean}
 */
// 


var wordBreak = function(s, wordDict) {
    // make wordDict become a map
    let mapD = new Map();
    for (let i = 0; i < wordDict.length; i++){
        mapD.set(wordDict[i], 1);
    }

    let sl = s.length;

    let dp = new Array(sl+1).fill(false);
    dp[0] = true;

    for (let i = 1; i < sl+1; i++){
       for (let j = 0; j < i; j++){

            if (dp[j] && mapD.has(s.substr(j, i-j))){
                dp[i] = true;
                break;
            }
       }
    }


    return dp[sl];
    
};