/**
 * @param {string[]} words
 * @param {number[]} groups
 * @return {string[]}
 */
var getLongestSubsequence = function(words, groups) {
    let n = words.length;

    // make dp, preList
    let dp = new Array(n).fill(1);
    let preL = new Array(n).fill(-1);

    // state transition
    for (let i = 0; i < n; i++){
        for (let j = 0; j < i; j++){
            if (groups[i] != groups[j]){
                if (dp[i] < dp[j]+1){
                    dp[i] = dp[j] + 1;
                    preL[i] = j;
                }
            }
        }
    }

    // find max
    let maxL = 0, endId = 0;
    for (let i = 0; i < n; i++){
        if (maxL < dp[i]){
            maxL = dp[i];
            endId = i;
        }
    }

    let result = new Array("").fill(maxL);
    for (let i = maxL-1; i >= 0; i--){
        result[i] = words[endId];
        endId = preL[endId];
    }

    return result;
};