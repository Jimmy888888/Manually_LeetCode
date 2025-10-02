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