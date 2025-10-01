func tribonacci(n int) int {
    // check n
    switch n {
        case 0:
            return 0
        case 1:
            return 1
        case 2:
            return 1
    }

    // recorder
    t1, t2, t3 := 0, 1, 1
    res := 0

    for i := 3; i <= n; i++{
        res = t1 + t2 + t3
        t1 = t2
        t2 = t3
        t3 = res
    }

    return res
}