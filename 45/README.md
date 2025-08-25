

# Greedy

take the max num in "every" step, untill reach the nums.length-1

need to find the max in "every" step

should more clear

defind : 1. find 2. jump

```js
/**
 * @param {number[]} nums
 * @return {number}
 */

var jump = function(nums) {
    let min_steps = 0;
    let cur_bound = 0;
    let cur_max = 0;

    let n = nums.length - 1

    if (n <= 0){
        return 0;
    }

    for (let i = 0; i < n; i++){
        cur_max = Math.max(cur_max, i + nums[i]);

        if(i === cur_bound){
            cur_bound = cur_max;
            min_steps += 1;
        }

        if (cur_bound >= n) {
            break;
        }

    }

    return min_steps;
};
```

### result: pass

# DP

f(i) = minS = min(nums[i] + nums[j] + muns[k]...)

### recursive case : 

assume at point n, min steps = Step(n) = 

func_find_pre(find one num that can reach n, and has the 

smallest id at the same time) + Step(n-1) 

### func_find_pre:

loop whole nums? in worst case : O(n*n) -> any good idea?



```js
var jump = function(nums) {
    
    let min_steps = 0;
    let i = nums.length - 1;

    // when find the pre_point, i = pre_point
    // i = 0 means find the min_steps
    while (i > 0){
        // find pre point
        let pre_point = i;
        for (let j = i; j >= 0; j--){
            if (j + nums[j] < i){
                continue;
            }

            if (j < pre_point){
                pre_point = j;
            }
        }
        // jump to pre_point
        i = pre_point;
        min_steps += 1;
    }

    return min_steps;
};
```

### result: pass
