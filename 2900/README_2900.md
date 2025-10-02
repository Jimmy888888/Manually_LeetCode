# 2900

## force solution:

```c
/**
 * Note: The returned array must be malloced, assume caller calls free().
 */
void initIntList(int* intList, int len){
    for (int i = 0; i < len; i++){
        intList[i] = 0;
    }
}

void copyIntList(int* targetList, int* curList, int len){
    initIntList(targetList, len);
    for (int i = 0; i < len; i++){
        targetList[i] = curList[i];
    }
}

char** getLongestSubsequence(char** words, int wordsSize, int* groups, int groupsSize, int* returnSize) {

    int cur = 0, curL = 0, maxL = 0;
    int* curList = (int*)malloc(groupsSize * sizeof(int));
    int* maxList = (int*)malloc(groupsSize * sizeof(int));
    initIntList(maxList, groupsSize);

    for (int i = 0; i < groupsSize; i++){
        initIntList(curList, groupsSize);
        curL = 1;
        curList[i] = 1;
        cur = groups[i];

        for (int j = i; j < groupsSize; j++){
            if (cur != groups[j]){
                cur = groups[j];
                curL++;
                curList[j] = 1;
            }
        }

        if (maxL < curL){
            copyIntList(maxList, curList, groupsSize);
            maxL = curL;
        }
    }

    char** result = (char**)malloc(maxL * sizeof(char*));
    int j = 0;

    for (int i = 0; i < groupsSize; i ++){
        if (maxList[i] != 0){
            result[j] = words[i];
            j++;
        }
    }

    *returnSize = j;
    free(curList);
    free(maxList);

    return result;
}
```



## DP

### assume
- cannot be consecutive 0 or 1

- dp[i] is the longest subsequence, and word[i] is the end

- if groups[j] != groups[i], i > j >= 0,   dp[i] = max(dp[i], dp[j]+1)

- use array preWord[i] to recorder the subsequence, 

- i means the end of subsequence is words[i], 
 
- and words[preWord[i]] is the previous word of words[i]

```c
/**
 * Note: The returned array must be malloced, assume caller calls free().
 */

char** getLongestSubsequence(char** words, int wordsSize, int* groups, int groupsSize, int* returnSize) {
    // make dp, preWord
    int* dp = (int*)malloc(wordsSize * sizeof(int));
    int* preWord = (int*)malloc(wordsSize * sizeof(int));
    // init
    for (int i = 0; i < wordsSize; i++){
        dp[i] = 1;
        preWord[i] = -1;
    }

    // state transition equation
    for (int i = 1; i < wordsSize; i++){
        for (int j = 0; j < i; j++){
            if (groups[i] != groups[j]){
                if (dp[i] < dp[j] + 1){
                    dp[i] = dp[j] + 1;
                    preWord[i] = j;
                }
            }
        }
    }

    // find the longest
    int max = 0, lastId = -1;
    for (int i = 0; i < wordsSize; i++){
        if (max < dp[i]){
            max = dp[i];
            lastId = i;
        }
    }
    // form the result from preWord
    char** result = (char**)malloc(max * sizeof(char*));
    for (int i = max-1; i >= 0; i--){
        result[i] = words[lastId];
        lastId = preWord[lastId];
    }

    free(dp);
    free(preWord);

    *returnSize = max;

    return result;
}
```