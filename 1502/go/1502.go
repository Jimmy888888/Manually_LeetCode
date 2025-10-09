func canMakeArithmeticProgression(arr []int) bool {
    sort.Ints(arr)
    commDiff := arr[0] - arr[1]
    n := len(arr)

    for i := 1; i < n-1; i++ {
        if (arr[i] - arr[i+1]) != commDiff {
            return false
        }
    }

    return true
}