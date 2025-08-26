# 383

check whether words in magazine can build words in ransomNote

## way 1: O(n*m) 

check every word in ransomNote through magazine, and change the word to "-" if a word had found


```js
var canConstruct = function(ransomNote, magazine) {
    let m = ransomNote.length;
    let n = magazine.length;
    let magazineArray = magazine.split('');
    for (let i = 0; i < m; i ++){
        let find = false;
        for (let j = 0; j < n; j++){
            if (ransomNote[i] === magazineArray[j]){
                magazineArray[j] = "-";
                find = true;
                break;
            }
        }
        if (find == false){
            return false;
        }
    }

    return true;
};
```


## way 2: O(m)
make a map for magazine, and go through ransomNote

```js
var canConstruct = function(ransomNote, magazine) {
    let magazineMap = new Map();

    for (let chr of magazine){
        let count = 1;
        if (magazineMap.has(chr)){
            count = magazineMap.get(chr) + 1;
        }
        magazineMap.set(chr, count);
    }


    for (let chr of ransomNote){
        if(magazineMap.has(chr)){
            if(magazineMap.get(chr) > 0){
                magazineMap.set(chr, magazineMap.get(chr) - 1);
            }else{
                return false;
            }
            
        }else{
            return false;
        }
    }
    return true;
};
```


# Intresting result

because  1 <= ransomNote.length, magazine.length <= 105

way 1 is faster than way 2

way 1 : under 10ms

way 2 : 20~30 ms
