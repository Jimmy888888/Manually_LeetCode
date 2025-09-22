/**
 * @param {number} n
 * @return {number}
 */
var fib = function(n) {
    if (n === 0){
        return 0;
    }

    if (n === 1 || n === 2){
        return 1;
    }

    let pre1 = 1, pre2 = 0, res = 0;

    for (let i = 2; i <= n; i++){
        res = pre1 + pre2;
        pre2 = pre1;
        pre1 = res;
    }

    return res;
};