# 66

## description
```
You are given a large integer represented as an integer array digits, where each digits[i] is the ith digit of the integer. The digits are ordered from most significant to least significant in left-to-right order. The large integer does not contain any leading 0's.

Increment the large integer by one and return the resulting array of digits.
```

## solution

```js
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
```