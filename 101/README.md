# LeetCode 101
### Symmetric Tree
<img width="355" alt="image" src="https://github.com/user-attachments/assets/8ccbf624-327e-4381-93e6-520278110d81" /> </br>
</br>
### Symmetric: make sure two sub trees are the same.
1.</br>
Q: How to compare two trees?</br>
A: Flatten two trees into two lists, and then compare them.</br>
2.</br>
Q: How to flatten a tree?</br>
A: recursive, Ask GPT how to flatten a tree by using recursive: there are three travesal types for flattening a tree</br>

### Three traversal type:
#### DFS, Depth-first Search
##### Preorder:
root->left->right</br>
##### Inorder:
left->root->right</br>
##### Postorder:
left->right->root</br>
</br>
### Another traversal type(not recursive)
#### BFS, Breadth-First Search
##### level-order:
layer1->layer2->layeer3

### C++ code:

```cpp
class Solution {
public:
    bool isSymmetric(TreeNode* root) {
        if(!root){
            return true;
        }
        
        vector<int> l_result, r_result;
        
        perOrder(root->left, l_result);
        perOrder(root->right, r_result);
        
        return l_result == r_result;
    }

    void perOrder(TreeNode *root, vector<int> &result){
        if(!root){
            return;
        }
        
        result.push_back(root->val);
        perOrder(root->left, result);
        perOrder(root->right, result);
    }

};
```

### result:
#### fail , because "Symmetric" dosen't mean "the same", it's more like "mirror"

### Next, define "mirror"
#### left node value == right node value
#### Note: -100 <= Node.val <= 100
```cpp
class Solution {
public:
    bool isSymmetric(TreeNode* root) {
        if(!root){
            return true;
        }
        
        vector<int> l_result, r_result;
        
        perOrder(root->left, l_result);
        perOrder(root->right, r_result);
        
        return l_result == r_result;
    }

    void perOrder(TreeNode *root, vector<int> &result){
        if(!root){
            result.push_back(101);
            return;
        }
        
        result.push_back(root->val);
        perOrder(root->left, result);
        perOrder(root->right, result);
    }

    void perOrderRev(TreeNode *root, vector<int> &result){
        if(!root){
            result.push_back(101);
            return;
        }
        
        result.push_back(root->val);
        perOrderRev(root->right, result);
        perOrderRev(root->left, result);
    }
};
```

### result:
#### pass, time complexity O(N)

### other solution:
#### use thought of "Mirror"
```cpp
class Solution {
public:
    bool isSymmetric(TreeNode* root) {
        if(!root){return true;}
        return isMirror(root->left, root->right);
    }

    bool isMirror(TreeNode* t1, TreeNode* t2){
        if(!t1 && !t2){return true;}
        if(!t1 || !t2){return false;}

        return (t1->val == t2->val) &&
        isMirror(t1->left, t2->right) &&
        isMirror(t1->right, t2->left);
    }
};
```
### result:
#### pass, time complexity O(N)

