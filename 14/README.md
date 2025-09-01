# 14

need to find the longest common prefix string

## solution 1

```js
/**
 * @param {string[]} strs
 * @return {string}
 */

var longestCommonPrefix = function(strs) {
    //find the longest common prefix, index and char in str must the same
    //go through every str, compare two str at once

    let comPreStrs = new Array();
    let comPreChr = "";
    let strId = 0;
    let lastId = 0;
    let n = strs.length;
    let finding = true;

    // keep finding untill strId > {the shortest str}.length
    comPreStrs.push("");
    while(finding) {
        finding = true;
        if (strId > strs[0].length - 1){
            finding = false;
            continue;
        }
        comPreChr = strs[0][strId];

        for (let i = 0; i < n; i++){
            if (strId > strs[i].length - 1){
                comPreChr = "";
                finding = false;
                break;
            }
            
            if (comPreChr !== strs[i][strId]){
                comPreChr = "";
                break;
            }
        }

        strId += 1;

        if (comPreChr !== "") {
            if (strId === lastId + 1){
                comPreStrs[comPreStrs.length - 1] += comPreChr;
                lastId = strId;
            }else{
                comPreStrs.push("");
                comPreStrs[comPreStrs.length - 1] += comPreChr;
                lastId = strId;
            }
        }
    }

    let theLongest = 0;
    for (let i = 0; i < comPreStrs.length; i++){
        if (comPreStrs[i].length > comPreStrs[theLongest].length){
            theLongest = i;
        }
    }
    
    return comPreStrs[theLongest];
};
```

#### result : fail

this solution can't fit LeetCode 14, because 14 only need the "first" common prefix string.

this solution can find the real longest common string.

failed at test case 122 / 126

Input:
```
strs = ["flower","fkow"]
```

Output:
```
"ow"
```

Expected:
```
"f"
```


## solution 2

just simplily find the "first" common prefix string

```go
func longestCommonPrefix(strs []string) string {
    
    if len(strs) == 0 {
        return ""
    }

    // just find the first common prefix string
    comPreStr := ""
    comFinding := true

    for i := 0; i < len(strs[0]); i++ {
        comFinding = true
        for j := 0; j < len(strs); j++ {

            if i > len(strs[j]) - 1 {
                comFinding = false
                break
            }

            if strs[0][i] != strs[j][i] {
                comFinding = false
                break
            }
        }

        if !comFinding {
            break
        }
        comPreStr += string(strs[0][i])
    }

    return comPreStr
}
```