# 67

## add binary

Given two binary strings a and b, return their sum as a binary string.

## solution

```js
/**
 * @param {string} a
 * @param {string} b
 * @return {string}
 */
var addBinary = function(a, b) {
    const m = a.length;
    const n = b.length;

    let ia = m-1;
    let ib = n-1;
    let carry1 = 0;
    let carry2 = 0;
    let res = "";
    let bita = "0";
    let bitb = "0";
    let curbit = "0";

    while (ia >= 0 || ib >= 0){
        bita = "0";
        bitb = "0";
        curbit = "0";
        carry2 = 0;

        if(ia >= 0){
            bita = a[ia];
        }

        if(ib >= 0){
            bitb = b[ib];
        }

        if(bita === "1" && bitb === "1"){
            curbit = "0";
            carry2 = 1;
        }else if(bita === "1" && bitb === "0"){
            curbit = "1";
            carry2 = 0;
        }else if(bita === "0" && bitb === "1"){
            curbit = "1";
            carry2 = 0;
        }else if(bita === "0" && bitb === "0"){
            curbit = "0";
            carry2 = 0;
        }


        if(carry1 === 1){
            if(curbit === "0"){
                curbit = "1";
            }else{
                curbit = "0";
                carry2 = 1;
            }
        }

        res = curbit + res;

        carry1 = carry2;
        ia -= 1;
        ib -= 1;
    }

    if (carry1 === 1){
        res = "1" + res;
    }

    return res;
};
```

### another solution

use the big int function

```py
class Solution:
    def addBinary(self, a: str, b: str) -> str:
        anum = int(a, 2)
        bnum = int(b, 2)

        result = anum + bnum

        return bin(result)[2:]#for cuting "0b"
```