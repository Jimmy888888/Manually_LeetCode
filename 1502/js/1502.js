/**
 * @param {number[]} arr
 * @return {boolean}
 */
var canMakeArithmeticProgression = function(arr) {
    arr.sort((a, b) => a - b);
    let n = arr.length
    let commDiff = arr[0] - arr[1]

    for (let i = 1; i < n-1; i++) {
        if (commDiff != (arr[i] - arr[i+1])) {
            return false;
        }
    }

    return true;
};