/**
 * @param {string} s
 * @param {string} t
 * @return {character}
 */
var findTheDifference = function(s, t) {
    let result = 0;

    for (const c of s){
        result = result ^ c.charCodeAt(0);
    }
    for (const c of t){
        result = result ^ c.charCodeAt(0);
    }

    return String.fromCharCode(result);
};