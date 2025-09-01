func longestCommonPrefix(strs []string) string {
    
    if len(strs) == 0 {
        return ""
    }

    // just find the first common prefix string
    comPreStr := ""
    comFinding := true

    for i := 0; i < len(strs[0]); i++ {
        comFinding = true
        for j := 0; j < len(strs); j++ {

            if i > len(strs[j]) - 1 {
                comFinding = false
                break
            }

            if strs[0][i] != strs[j][i] {
                comFinding = false
                break
            }
        }

        if !comFinding {
            break
        }
        comPreStr += string(strs[0][i])
    }

    return comPreStr
}