import "sort"

func hIndex(citations []int) int {
    // sort citations in descending order
    sort.Slice( citations, func(i, j int) bool {
        return citations[i] > citations[j]
    })

    hId := 0
    for i, citedTimes := range citations {
        if i+1 > citedTimes {
            hId = i
            return hId
        }
    }

    if hId == 0 {
        hId = len(citations)
    }

    return hId
}