func findTheDifference(s string, t string) byte {
    sSum := 0
    tSum := 0

    for _, ch := range s {
        sSum += int(ch)
    }
    for _, ch := range t {
        tSum += int(ch)
    }

    return byte(tSum - sSum)
}