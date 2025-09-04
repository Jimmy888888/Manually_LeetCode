# hash map

make wordDict become a map

use a string buffer, it come from a part of s

the buffer will keep collecting words from s, untill it fit by wordDict, or fit nothing

```js
var wordBreak = function(s, wordDict) {
    // 1 <= wordDict[i].length <= 20
    let maxWl = 20;
    // make wordDict become a map
    let mapD = new Map();
    for (let i = 0; i < wordDict.length; i++){
        mapD.set(wordDict[i], 1);
    }

    let buf = "";
    let checkId = -1;
    let sl = s.length;

    for (let i = 0; i < sl; i++){
        buf += s[i];

        console.log(buf);

        if (mapD.has(buf)){
            buf = "";
            checkId = i;
        }else{
            if (buf.length > 20){
                return false;
            }
        }

    }

    console.log(checkId);

    return (checkId === sl - 1);
    
};
```


### how to decide which dict word the buffer can fit?  

 --> "aaaaaaa" ["aaaa","aaa"]

a

aa

aaa

a

aa

aaa

a

5



# DP

assume dp[i] = if s[i] can be formed by wordDict

if dp[i] == true --> s.substring(j, i) in wordDict && dp[j]=true

```js
/**
 * @param {string} s
 * @param {string[]} wordDict
 * @return {boolean}
 */
// 


var wordBreak = function(s, wordDict) {
    // make wordDict become a map
    let mapD = new Map();
    for (let i = 0; i < wordDict.length; i++){
        mapD.set(wordDict[i], 1);
    }

    let sl = s.length;

    let dp = new Array(sl+1).fill(false);
    dp[0] = true;

    for (let i = 1; i < sl+1; i++){
       for (let j = 0; j < i; j++){

            if (dp[j] && mapD.has(s.substr(j, i-j))){
                dp[i] = true;
                break;
            }
       }
    }


    return dp[sl];
    
};
```