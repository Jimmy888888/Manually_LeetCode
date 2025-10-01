/**
 * @param {string} sequence
 * @param {string} word
 * @return {number}
 */
var maxRepeating = function(sequence, word) {
    let sl = sequence.length;
    let wl = word.length;
    // make dp
    let dp = new Array(sl+1).fill(0);

    // record
    let max = 0;

    // state transition equation
    for (let i = 1; i < sl+1; i++){
        if (i >= wl){
            if (sequence.slice(i-wl, i) === word){
                dp[i] = dp[i-wl] + 1;
                if (max < dp[i]){
                    max = dp[i];
                }
            }
        }
    }

    return max;
};

//// another solution:

/**
 * @param {string} sequence
 * @param {string} word
 * @return {number}
 */
var maxRepeating = function(sequence, word) {
    let count=0;
    let repeat=word;
    while(sequence.includes(repeat)){
        count++;
        repeat+=word;

    }
    return count;
    
};