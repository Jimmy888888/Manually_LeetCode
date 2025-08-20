import (
    "unicode"
    "fmt"
    )

// modify func
func modify(r rune) rune {
    if unicode.IsLetter(r) {
        r = unicode.ToLower(r)
    }else if unicode.IsDigit(r){
        r = r
    }else{
        // reset the rune
        r = rune(0)
    }
    return r
}

func isPalindrome(s string) bool {
    // two point
    // first do modify
    var sModified []rune
    for _, ru := range s {
        rM := modify(rune(ru))
        if rM != 0 {
            sModified = append(sModified, rM)
        }
    }

    // fmt.Printf("%s", string(sModified))

    // two point
    sLen := len(sModified)
    i := 0
    j := sLen - 1
    for i < sLen {
        if sModified[i] != sModified[j]{
            return false
        }
        j--
        i++
    }

    return true
}