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