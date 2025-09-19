/**
 * @param {string} s
 * @return {number}
 */
var numDecodings = function(s) {
    // check
    if (s[0] === "0"){
        return 0;
    }
    let n = s.length;
    if (n === 1){
        return 1;
    }

    // assume preTwoBity start with empty ""
    // assume preOneBity start with s[0]
    let preOneBity = 1;
    let preTwoBity = 1;
    let res = 0;

    for (let i = 2; i <= n; i++){
        let num1 = s[i-1] - "0";
        let num2 = (s[i-2] - "0")*10 + num1;
        let find1 = (num1 > 0 && num1 < 10);
        let find2 = (num2 > 9 && num2 < 27);

        res = 0;

        if ( find1){
            res += preOneBity;
        }
        if (find2){
            res += preTwoBity;
        }

        if (res === 0){
            return 0;
        }

        preTwoBity = preOneBity;
        preOneBity = res;
    }
    return res;
};