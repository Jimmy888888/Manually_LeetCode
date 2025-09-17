# 53

## DP
 
### assume:

- contiguous

- inheritance or not , at every nums[i], need to consider:

- nums[i] is larger or nums[i] + nums[i-1] ~ nums[j] (j <= i-1)larger?

```js
/**
 * @param {number[]} nums
 * @return {number}
 */
var maxSubArray = function(nums) {
    let n = nums.length;
    if (n === 1) {
        return nums[0];
    }
    
    let curS = nums[0];
    let golS = nums[0];

    for (let i = 1; i < n; i++){
        let num = nums[i];
        if (curS + num < num) {
            curS = num;
        }else{
            curS += num;
        }

        if (curS > golS) {
            golS = curS;
        }
    }

    return golS;
};
```
