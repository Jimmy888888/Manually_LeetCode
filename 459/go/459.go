// if s is a substring of ( s without first char + s without last char )
// than s must can be formed by a substring

func repeatedSubstringPattern(s string) bool {
    n := len(s)
    sCheck := s[1:] + s[0:n-1]

    return strings.Contains(sCheck, s)
}