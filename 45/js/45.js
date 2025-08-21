
var jump = function(nums) {
    let min_step = 0;
    let i = 0;
    while ( i < nums.length - 1 ){
        let loc_max = 0;
        let jup_to = 0;

        // test 
        if ( i + nums[i] > nums.length - 1){
            min_step += 1;
        }

        // find
        for (let j = i+1; j <= i + nums[i]; j++){
            }else if(loc_max < nums[j]){
                loc_max = nums[j];
                jup_to = j;
            }
        }
        // handle edge case
        if (loc_max === 0){
            return 0;
        }

        // jump
        i = jup_to;
        console.log(jup_to);
        min_step += 1;
        
    }

    return min_step;
};