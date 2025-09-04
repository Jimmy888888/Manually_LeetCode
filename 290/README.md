# 290

use double maps to check the pattern and ths s

```js
/**
 * @param {string} pattern
 * @param {string} s
 * @return {boolean}
 */
var wordPattern = function(pattern, s) {
    let pl = pattern.length;
    
    let arrS = s.split(' ');
    if (pl !== arrS.length){
        return false;
    }

    // make maps
    let mapS = new Map();
    let mapP = new Map();

    for (let i = 0; i < pl; i++){
        let st = arrS[i];
        let p = pattern[i];

        if (mapS.has(st)) {
            if (mapS.get(st) !== p) {
                return false;
            }
        }else{
            mapS.set(st, p);
        }

        if (mapP.has(p)) {
            if (mapP.get(p) !== st) {
                return false;
            }
        }else{
            mapP.set(p, st);
        }
    }

    return true
};
```