func canConstruct(ransomNote string, magazine string) bool {
    checkMap := make(map[rune]int)

    for _, str := range magazine {
        checkMap[str] = checkMap[str] + 1
    }

    for _, str := range ransomNote {
        if checkMap[str] <= 0 {
            return false
        }

        checkMap[str] = checkMap[str]  - 1
    }

    return true
}