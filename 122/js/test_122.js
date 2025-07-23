const assert = require('assert');
const groupAnagrams = require('./49.js');

function testGroupAnagrams() {
    // Test case 1
    const strs1 = ["eat", "tea", "tan", "ate", "nat", "bat"];
    const expected1 = [["eat", "tea", "ate"], ["tan", "nat"], ["bat"]];
    const result1 = groupAnagrams(strs1);
    assert.deepStrictEqual(result1.map(arr => arr.sort()).sort(), expected1.map(arr => arr.sort()).sort());

    // Test case 2
    const strs2 = [""];
    const expected2 = [[""]];
    const result2 = groupAnagrams(strs2);
    assert.deepStrictEqual(result2.map(arr => arr.sort()).sort(), expected2.map(arr => arr.sort()).sort());

    // Test case 3
    const strs3 = ["a"];
    const expected3 = [["a"]];
    const result3 = groupAnagrams(strs3);
    assert.deepStrictEqual(result3.map(arr => arr.sort()).sort(), expected3.map(arr => arr.sort()).sort());

    console.log("All test cases passed!");
}

testGroupAnagrams();