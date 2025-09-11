# 28

## description
```
Given two strings needle and haystack, return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.
```

## solution
```go
func strStr(haystack string, needle string) int {
    return strings.Index(haystack, needle)
}
```
