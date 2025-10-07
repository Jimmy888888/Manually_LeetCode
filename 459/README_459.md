# 459

find a substring

fource solution

using the first char of "s"

if lenght of substring longer than 1/2 "s", stop

### outer loop :
 find start and end char
### inner loop :
 check the substring 

```c
bool checkSubStr(int start, int end, char* s, int sl){
    int subl = end - start + 1;
    int subI = 0;
    for (int i = end+1; i < sl; i++){
        subI = i % subl;
        if (s[i] != s[subI]){
            return false;
        }
        if (i == sl - 1){
            return true;
        }
    }
    return false;
}

bool repeatedSubstringPattern(char* s) {
    int n = strlen(s);
    for (int i = 0; i <= n/2; i++){
        if (n % (i+1) != 0){
            continue;
        }
        if (checkSubStr(0, i, s, n)){
            return true;
        }
    }
    return false;
}
```
