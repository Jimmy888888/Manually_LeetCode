# 1822

if there is an odd number of negative numbers, return -1

if there is an even number of negative numbers, return 1

if there has any 0, return 0


```c
int arraySign(int* nums, int numsSize) {

    int result = 1;
    for (int i = 0; i < numsSize; i++){
        if (nums[i] == 0){
            return 0;
        }else if (nums[i] > 0){
            result = result;
        }else if (nums[i] < 0){
            result = result * -1;
        }
    }

    return result;
}
```