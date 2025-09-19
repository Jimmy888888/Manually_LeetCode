// shift right
// i & 1 , check the last bit, odd or even

/**
 * @param {number} n
 * @return {number[]}
 */
var countBits = function(n) {
    let res = new Array(n+1).fill(0);

    for (let i = 0; i < n+1; i++){
        res[i] = res[i >> 1] + (i & 1)
    }
    return res
};