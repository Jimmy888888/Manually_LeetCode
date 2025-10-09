/**
 * @param {number[]} nums
 * @return {number}
 */
var arraySign = function(nums) {
    let res = 1;
    for (const n of nums) {
        if (n === 0){
            return 0;
        }else if (n < 0){
            res *= -1;
        }
    }
    return res;
};