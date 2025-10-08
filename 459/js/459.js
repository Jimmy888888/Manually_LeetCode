/**
 * @param {string} s
 * @return {boolean}
 */
var repeatedSubstringPattern = function(s) {
    let sCheck = s.slice(1) + s.slice(0,-1);
    return sCheck.includes(s);
};