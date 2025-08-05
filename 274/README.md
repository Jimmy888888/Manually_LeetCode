## h-index:

an author has at least 'h' papers that have each been cited 'h' times

## greedy : 

values :

- citedTimes : the times of cited

- h papers : the papers that have each been cited 'h' times


## way 1 : sort the slices first

```go
import "sort"

func hIndex(citations []int) int {
    // sort citations in descending order
    sort.Slice( citations, func(i, j int) bool {
        return citations[i] > citations[j]
    })

    hId := 0
    for i, citedTimes := range citations {
        if i+1 >= citedTimes {
            hId = citedTimes
            break
        }
    }

    return hId
}
```

#### result : failed in TestCase [100] , because only consider the citedTimes , should consider "i+1" too 

## use the same way , but the point is paper amount

find the first paper index that paper amount > citedTimes

```go
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
```
