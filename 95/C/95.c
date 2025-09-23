/**
 * Definition for a binary tree node.
 * struct TreeNode {
 *     int val;
 *     struct TreeNode *left;
 *     struct TreeNode *right;
 * };
 */
/**
 * Note: The returned array must be malloced, assume caller calls free().
 */

 struct TreeNode** _generateTrees(int start, int end, int *returnSize) {
    if (start > end){
        struct TreeNode** result = (struct TreeNode**)malloc(sizeof(struct TreeNode*));
        result[0] = NULL;
        *returnSize = 1;
        return result;
    }

    struct TreeNode** allTrees = NULL;
    int capacity = 10;
    int count = 0;
    allTrees = (struct TreeNode**)malloc(sizeof(struct TreeNode*) * capacity);

    for (int i = start; i <= end; i++){
        int leftTreeSize = 0;
        struct TreeNode** leftTrees = _generateTrees(start, i-1, &leftTreeSize);

        int rightTreeSize = 0;
        struct TreeNode** rightTrees = _generateTrees(i+1, end, &rightTreeSize);

        for (int l = 0; l < leftTreeSize; l++){
            for (int r = 0; r < rightTreeSize; r++){
                struct TreeNode* root = (struct TreeNode*)malloc(sizeof(struct TreeNode));
                root->val = i;
                root->left = leftTrees[l];
                root->right = rightTrees[r];

                if (count >= capacity){
                    capacity *= 2;
                    allTrees = (struct TreeNode**)realloc(allTrees, sizeof(struct TreeNode*) * capacity);
                }

                allTrees[count++] = root;
            }
        }

        free(leftTrees);
        free(rightTrees);
    }

    *returnSize = count;
    return allTrees;
}

struct TreeNode** generateTrees(int n, int* returnSize) {
    if ( n == 0){
        *returnSize = 0;
        return NULL;
    }

    return _generateTrees(1, n, returnSize);
}