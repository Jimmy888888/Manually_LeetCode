# 1025

## Game Theory
 
- 1 -> A

- 2 -> A1 B

- 3 -> A1 B1 A

- 4 -> A2 B1 A / A1 B1 A1 B

- 5 -> A1 B1 A1 B1 A

### if get
  
- odd  -> must loss
  
- even -> must win

```c
bool divisorGame(int n) {

    return (n % 2 == 0);
}
```


## dp

### assume 

- dp[i] is the result of must win/loss at i, 1<= i <= n

 - so len(dp) = n+1

- find n % x == 0

- how to decide the x , the biggest? --> go throught all possibilitys

- find x :

- a loop from n-1 to 1

```c
bool divisorGame(int n) {
    // check
    if (n == 1){
        return false;
    }
    
    // make dp
    bool* dp = (bool*)malloc((n+1) * sizeof(bool));
    memset(dp, 0, (n+1) * sizeof(bool));
    // dp[1] = false

    // go through , and Alice go first
    for (int i = 2; i <= n; i++){
        for (int j = 1; j < i; j++){
            if (i % j == 0){
                if (dp[i-j] == false){
                    dp[i] = true;
                    break;
                }
            }
        }
    }

    bool res = dp[n];
    free(dp);

    return res;
}
```
