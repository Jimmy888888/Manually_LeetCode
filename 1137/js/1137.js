/**
 * @param {number} n
 * @return {number}
 */
var tribonacci = function(n) {
    // check n
    switch (n){
        case 0:
            return 0;
        case 1:
            return 1;
        case 2:
            return 1;
        default:
            break;
    }

    // recorder
    let t1 = 0, t2 = 1, t3 = 1;
    let res = 0;

    // state trsnsition equation
    for (let i = 3; i <= n; i++){
        res = t1 + t2 + t3;
        t1 = t2;
        t2 = t3;
        t3 = res;
    }

    return res;
};