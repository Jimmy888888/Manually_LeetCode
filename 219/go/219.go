func containsNearbyDuplicate(nums []int, k int) bool {
    n := len(nums)
    check := make(map[int]int, n)

    for id, val := range nums {
        preId, find := check[val]

        if find && id - preId <= k {
            return true
        }
        check[val] = id
    }

    return false
}