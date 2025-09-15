/**
 * @param {number[]} nums
 * @param {number} k
 * @return {boolean}
 */
var containsNearbyDuplicate = function(nums, k) {
    let n = nums.length;
    let check = new Map();

    for (let i = 0; i < n; i++){
        num = nums[i];
        if (check.has(num)) {
            let preId = check.get(num)

            if ((i - preId) <= k) {
                return true;
            }
        }

        check.set(num, i);
    }

    return false;
};