func isIsomorphic(s string, t string) bool {
    mapS := make(map[byte]byte)
    mapT := make(map[byte]byte)

    n := len(s)

    for i := 0; i < n; i++ {
        chrS := s[i]
        chrT := t[i]
        valS, findS := mapS[chrS]

        if findS {
            if valS != chrT {
                return false
            }
        }else{
            mapS[chrS] = chrT
        }

        valT, findT := mapT[chrT]

        if findT {
            if valT != chrS {
                return false
            }
        }else{
            mapT[chrT] = chrS
        }
    }

    return true
}