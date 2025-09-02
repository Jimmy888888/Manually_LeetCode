/**
 * @param {string} s
 * @param {string} t
 * @return {boolean}
 */

// a map for recoeding
// t.length == s.length
// a loop go through two strings

var isIsomorphic = function(s, t) {
    // make a map
    let checkS = new Map();
    let checkT = new Map();
    let n = s.length;

    // s is key , t is val
    for (let i = 0; i < n; i++){
        let chrS = s[i];
        let chrT = t[i];

        if (checkS.has(chrS)){
            if (checkS.get(chrS) !== chrT){
                return false;
            }
        }else{
            checkS.set(chrS,chrT);
        }

        if (checkT.has(chrT)){
            if (checkT.get(chrT) !== chrS){
                return false;
            }
        }else{
            checkT.set(chrT, chrS);
        }
    }

    return true;
};