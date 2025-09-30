# 1668

find the longest sequence of "word"

for loop go through and count the "k"

recorde the max "k"

## DP

assume dp[i] is the current "k" at sequence[i-1]

j = (i-1) - len(word) + 1

if sequence[j~(i-1)] is a "word", dp[i] = 1 + dp[j]

if sequence[j~(i-1)] is not a "word", dp[i] = 0

return dp[n-1]

```c
bool isSubSame(char* s, int s_start, char* sub, int sub_len){
    int j = 0;
    for (int i = 0; i < sub_len; i++){
        j = i + s_start;
        if (s[j] != sub[i]){
            return false;
        }
    }
    return true;
}

int maxRepeating(char* sequence, char* word) {
    int n = strlen(sequence);
    int m = strlen(word);
    int max = 0;

    // make dp
    int* dp = (int*)malloc((n+1) * sizeof(int));
    for (int i = 0; i < n+1; i++){
        dp[i] = 0;
    }

    // // base case
    // if (m == 1 && word[0] == sequence[0]){
    //     dp[0] = 1;
    // }

    //  0 1 2 3
    // -1 0 1 2

    for (int i = 1; i < n+1; i++){
        if (i >= m){
            if (isSubSame(sequence, (i-1)-m+1, word, m)){
                dp[i] = 1 + dp[i-m];
                if (max < dp[i]){
                    max = dp[i];
                }
            }
        }
    }

    free(dp);

    return max;
}
```