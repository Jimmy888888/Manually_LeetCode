/**
 * @param {string} s
 * @param {string} t
 * @return {boolean}
 */
var isAnagram = function(s, t) {
    // check
    let sl = s.length;
    let tl = t.length;
    if (sl !== tl){
        return false;
    }

    // make map
    let mapT = new Map();
    for (const chT of t){
        const count = mapT.get(chT) || 0;
        mapT.set(chT, 1+count);
    }

    // check anagram
    for (const chS of s){
        const count = mapT.get(chS) || 0;
        if (count <= 0){
            return false;
        }else{
            mapT.set(chS, count-1);
        }
    }

    return true;
};