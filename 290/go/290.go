func wordPattern(pattern string, s string) bool {
    n := len(pattern)
    // turn s into slice
    sliceS := strings.Split(s, " ")
    m := len(sliceS)

    if n != m {
        return false
    }

    // make double map
    mapP := make(map[byte]string)
    mapS := make(map[string]byte)

    for i:= 0; i < n; i++ {
        str := sliceS[i]
        byt := pattern[i]

        valS, findS := mapS[str]
        if findS {
            if valS != byt {
                return false
            }
        }else{
            mapS[str] = byt
        }

        valP, findP := mapP[byt]
        if findP {
            if valP != str {
                return false
            }
        }else{
            mapP[byt] = str
        }

    }

    return true
}