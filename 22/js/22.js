/**
 * @param {number} n
 * @return {string[]}
 */
var generateParenthesis = function(n) {
    let result = [];

    var backTracking = function(curS, left_num, right_num) {
        if (curS.length == 2 * n) {
            result.push(curS);
            return;
        }

        if (left_num < n){
            backTracking(curS + "(", left_num + 1, right_num);
        }

        if (right_num < left_num){
            backTracking(curS + ")", left_num, right_num + 1);
        }
    }

    backTracking("", 0, 0);

    return result;
};