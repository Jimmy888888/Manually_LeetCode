

# Greedy

take the max num in every step, untill reach the nums.length-1

need to find the max in every step

should more clear

defind : 1. find 2. jump

```js
/**
 * @param {number[]} nums
 * @return {number}
 */

var jump = function(nums) {
    let min_step = 0;
    let i = 0;
    let jup_to = 0;
    let loc_max_jup = 0;
    while ( i < nums.length - 1 ){

        // check whether can reach the end
        if ( i + nums[i] >= nums.length - 1){
            min_step += 1;
            console.log(i + nums[i]);
            break;
        }

        // find, the biggest and the farthest step
        jup_to = i + nums[i];
        loc_max_jup = jup_to;
        for (let j = i+1; j < i + nums[i]; j++){
            if(loc_max_jup < j + nums[j]){
                jup_to = j;
                loc_max_jup = j + nums[j];
            }
        }

        // handle edge case
        if (nums[jup_to] === 0){
            return 0;
        }

        // jump
        i = jup_to;
        console.log(i);
        min_step += 1;
        
    }

    return min_step;
};
```

### result: not pass 

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
