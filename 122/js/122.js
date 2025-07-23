/**
 * @param {string[]} strs
 * @return {string[][]}
 */
var groupAnagrams = function(strs) {
    // map sort 
    const groupMap = new Map();//[[string,string[]]]

    for (const str of strs){
        const charArray = str.split('');
        charArray.sort();
        const sortStr = charArray.join('');

        if (groupMap.has(sortStr)){
            groupMap.get(sortStr).push(str);
        }else{
            groupMap.set(sortStr, [str]);
        }
    }

    return Array.from(groupMap.values());
};

module.exports = groupAnagrams;