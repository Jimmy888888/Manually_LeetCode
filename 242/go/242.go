func isAnagram(s string, t string) bool {
    // check 
    if len(s) != len(t){
        return false
    }

    // make a map for t, if find t dump it
    mapT := make(map[rune]int, len(t))
    for _, chT := range t {
        mapT[chT] += 1
    }

    // check anagram
    for _, chS := range s{
        mapT[chS] -= 1
        if mapT[chS] < 0 {
            return false
        }
    }

    return true
}

// rune --> range str
// byte --> str[i]