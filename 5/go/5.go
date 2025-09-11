func longestPalindrome(s string) string {
    // check edge
    n := len(s)
    if n == 1 {
        return s
    }
    
    // use center expend strategy
    maxS := ""

    for i := 0; i < n; i++ {
        // find max
        curSOdd := ""
        curSEven := ""
        
        curSOdd = CenterExpend(i, i, s)

        if (i < n-1){
            curSEven = CenterExpend(i, i+1, s)
        }

        curMax := ""
        if len(curSOdd) > len(curSEven){
            curMax = curSOdd
        }else{
            curMax = curSEven
        }

        if len(curMax) > len(maxS) {
            maxS = curMax
        }
    }

    return maxS
}

func CenterExpend(rId int, lId int, s string) string {
    // check
    n := len(s)
    subS := ""

    for rId < n && lId >= 0 {
        if s[rId] == s[lId]{
            subS = s[lId:rId+1]
        }else{
            break
        }
        rId++
        lId--
    }

    return subS
}