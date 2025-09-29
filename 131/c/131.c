/**
 * Return an array of arrays of size *returnSize.
 * The sizes of the arrays are returned as *returnColumnSizes array.
 * Note: Both returned array and *columnSizes array must be malloced, assume caller calls free().
 */
 #include <string.h>
 #include <stdio.h>
 #include <stdlib.h>
 #include <string.h>
 #include <stdbool.h>
 
 char*** g_result;
 int g_result_size;
 int g_result_capacity;
 
 char** g_path;
 int g_path_size;
 int g_path_capacity;
 
 int* g_col_sizes;
 
 
 // DSF & backtrack
 void backtrack(const char* s, int start_index, int n, bool** dp){
     // base condiction
     if (start_index == n){
         // check g_result caoacity
         if (g_result_size >= g_result_capacity){
             g_result_capacity *= 2;
             g_result = (char***)realloc(g_result, g_result_capacity * sizeof(char**));
             g_col_sizes = (int*)realloc(g_col_sizes, g_result_capacity * sizeof(int));
         }
 
         // set mem
         g_result[g_result_size] = (char**)malloc(g_path_size * sizeof(char*));
 
         // copy
         for (int i = 0; i < g_path_size; i++){
             g_result[g_result_size][i] = (char*)malloc((strlen(g_path[i]) + 1) * sizeof(char));
             strcpy(g_result[g_result_size][i], g_path[i]);
         }
 
         // record the length
         g_col_sizes[g_result_size] = g_path_size;
 
         // add sum
         g_result_size++;
         return;
     }
 
     // recursive
     for (int i = start_index; i < n; i++){
         // check by using dp
         if (dp[start_index][i]){
             // get the palindrome string
             int sub_len = i - start_index + 1;
             char* sub_str = (char*)malloc((sub_len + 1) * sizeof(char));
             strncpy(sub_str, s + start_index, sub_len);
             sub_str[sub_len] = '\0';
 
             // add sub_str into g_path
             // check g_path capacity
             if (g_path_size >= g_path_capacity){
                 g_path_capacity *= 2;
                 g_path = (char**)realloc(g_path, g_path_capacity * sizeof(char*));
             }
 
             g_path[g_path_size] = sub_str;
             g_path_size++;
 
             // call next
             backtrack(s, i+1, n, dp);
 
             // delete the choose
             g_path_size--;
             free(g_path[g_path_size]);
         }
     }
 }
 
 char*** partition(char* s, int* returnSize, int** returnColumnSizes){
     int n = strlen(s);
     if (n == 0){
         *returnSize = 0;
         *returnColumnSizes = NULL;
         return NULL;
     }
 
     // make the dp
     bool** dp = (bool**)malloc(n * sizeof(bool*));
     for (int i = 0; i < n; i++){
         dp[i] = (bool*)malloc(n * sizeof(bool));
         memset(dp[i], 0, n * sizeof(bool));
     }
 
     // fill the dp
     for (int i = n-1; i > -1; i--){
         for (int j = i; j < n; j++){
             if((s[i] == s[j]) && (j-i < 2 || dp[i+1][j-1])){
                 dp[i][j] = true;
             }
         }
     }
 
     // init the gol vals
     g_result_capacity = 8;
     g_result = (char***)malloc(g_result_capacity * sizeof(char**));
     g_result_size = 0;
 
     g_col_sizes = (int*)malloc(g_result_capacity * sizeof(int));
 
     g_path_capacity = 8;
     g_path = (char**)malloc(g_path_capacity * sizeof(char*));
     g_path_size = 0;
 
 
     backtrack(s, 0, n, dp);
 
     // clean memory
     for(int i = 0; i < n; i++){
         free(dp[i]);
     }
     free(dp);
 
     free(g_path);
 
     // set LeetCode returns
     *returnSize = g_result_size;
     *returnColumnSizes = g_col_sizes;
 
     return g_result;
 }



 /////////////////////////
 //////runtime: 120ms/////
 /////////////////////////




 //////////////another version//////////////
 ////////////reduce malloc & free///////////
 ////////////runtime: 67ms//////////////////

 /**
 * Return an array of arrays of size *returnSize.
 * The sizes of the arrays are returned as *returnColumnSizes array.
 * Note: Both returned array and *columnSizes array must be malloced, assume caller calls free().
 */
#include <string.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <stdbool.h>

char*** g_result;
int g_result_size;
int g_result_capacity;

char** g_path;
int g_path_size;
int g_path_capacity;

int* g_col_sizes;


// DSF & backtrack
void backtrack(const char* s, int start_index, int n, char*** palin_cach){
    // base condiction
    if (start_index == n){
        // check g_result caoacity
        if (g_result_size >= g_result_capacity){
            g_result_capacity *= 2;
            g_result = (char***)realloc(g_result, g_result_capacity * sizeof(char**));
            g_col_sizes = (int*)realloc(g_col_sizes, g_result_capacity * sizeof(int));
        }

        // set mem
        g_result[g_result_size] = (char**)malloc(g_path_size * sizeof(char*));

        // copy
        for (int i = 0; i < g_path_size; i++){
            g_result[g_result_size][i] = strdup(g_path[i]);
        }

        // record the length
        g_col_sizes[g_result_size] = g_path_size;

        // add sum
        g_result_size++;
        return;
    }

    // recursive
    for (int i = start_index; i < n; i++){
        // check by using dp
        if (palin_cach[start_index][i] != NULL){
            // check g_path capacity
            if (g_path_size >= g_path_capacity){
                g_path_capacity *= 2;
                g_path = (char**)realloc(g_path, g_path_capacity * sizeof(char*));
            }

            g_path[g_path_size] = palin_cach[start_index][i];
            g_path_size++;

            // call next
            backtrack(s, i+1, n, palin_cach);

            // delete the choose
            g_path_size--;
        }
    }
}

char*** partition(char* s, int* returnSize, int** returnColumnSizes){
    int n = strlen(s);
    if (n == 0){
        *returnSize = 0;
        *returnColumnSizes = NULL;
        return NULL;
    }

    // make the palin_cach
    char*** palin_cach = (char***)malloc(n * sizeof(char**));
    for (int i = 0; i < n; i++){
        // set NULL
        palin_cach[i] = (char**)calloc(n, sizeof(char*));
    }

    // fill the palin_cach
    for (int i = n-1; i > -1; i--){
        for (int j = i; j < n; j++){
            if((s[i] == s[j]) && (j-i < 2 || palin_cach[i+1][j-1] != NULL)){
                int len = j - i + 1;
                palin_cach[i][j] = (char*)malloc((len+1) * sizeof(char));
                strncpy(palin_cach[i][j], s+i, len+1);
                palin_cach[i][j][len] = '\0';
            }
        }
    }

    // init the gol vals
    g_result_capacity = 8;
    g_result = (char***)malloc(g_result_capacity * sizeof(char**));
    g_result_size = 0;

    g_col_sizes = (int*)malloc(g_result_capacity * sizeof(int));

    g_path_capacity = 8;
    g_path = (char**)malloc(g_path_capacity * sizeof(char*));
    g_path_size = 0;


    backtrack(s, 0, n, palin_cach);

    // clean memory
    for(int i = 0; i < n; i++){
        for(int j = 0; j < n; j++){
            if (palin_cach[i][j] != NULL){
                free(palin_cach[i][j]);
            }
        }
        free(palin_cach[i]);
    }
    free(palin_cach);

    free(g_path);

    // set LeetCode returns
    *returnSize = g_result_size;
    *returnColumnSizes = g_col_sizes;

    return g_result;
}


