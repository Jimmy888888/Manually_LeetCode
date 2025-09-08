# 202

## check the happy number:

define happy number is a number that will end up with 1

and if not happy number, it will end up with cycle

the cycle means it will repeat the same number


```js
/**
 * @param {number} n
 * @return {boolean}
 */
var isHappy = function(n) {
    let sum = 0;
    let num = 0;
    let mapCheck = new Map();
    mapCheck.set(n, 0);

    while(true){
        num = n % 10;
        n = Math.floor(n / 10);
    
        sum += num ** 2;
        if (n === 0){
            if (sum === 1){
                return true;
            }else if (mapCheck.has(sum)){
                return false;
            }else{
                mapCheck.set(sum, 0);
                n = sum;
                sum = 0;
            }
        }
    }

    return false;
};
```