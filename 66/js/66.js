/**
 * @param {number[]} digits
 * @return {number[]}
 */
var plusOne = function(digits) {
    let carry = 0;
    let n = digits.length;

    if (digits[n-1] + 1 < 10){
        digits[n-1] = digits[n-1] + 1
        return digits;
    }else{
        carry = 1;
        digits[n-1] = (digits[n-1] + 1) % 10;
    }

    for (let i = n-2; i >= 0; i--){
        if (digits[i] + 1 < 10){
            digits[i] = digits[i] + 1;
            carry = 0;
            break;
        }else{
            carry = 1;
            digits[i] = (digits[i] + 1) % 10;
        }
    }

    if (carry === 1){
        digits.unshift(1);
    }

    return digits;
};