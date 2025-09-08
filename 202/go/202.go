func isHappy(n int) bool {
    num := 0
    sum := 0
    mapC := make(map[int]int)
    mapC[n] = 0

    for true {
        num = n % 10
        n = n / 10

        sum += num * num

        if n == 0 {
            if sum == 1 {
                return true
            }else if _, find := mapC[sum]; find{
                return false
            }else{
                n = sum
                sum = 0
                mapC[n] = 0
            }
        }
    }
    return false
}