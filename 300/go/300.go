// use an array, keep recording the smallest number in the range of nums[i]
// go through the nums, and replace or add new num

func lengthOfLIS(nums []int) int {
    // check
    n := len(nums)
    if n == 1 {
        return 1
    }

    // start go through nums
    Sub := make([]int, 1)
    Sub[0] = nums[0]

    for i := 0; i < n; i++{
        if (Sub[len(Sub)-1] < nums[i]){
            Sub = append(Sub, nums[i])
        }else{
            id := FindIdtoReplace(Sub, nums[i])
            Sub[id] = nums[i]
        }
    }

    return len(Sub)
}

func FindIdtoReplace(arr []int, target int) int{
    r := len(arr)-1
    l := 0

    for l <= r {
        mind := (l + r)/2

        if (arr[mind] == target){
            return mind
        }else if (arr[mind] > target){
            r = mind - 1
        }else{
            l = mind + 1
        }
    }

    return l
}


