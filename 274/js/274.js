/**
 * @param {number[]} citations
 * @return {number}
 */
var hIndex = function(citations) {
    // sort first 
    citations.sort(function(a,b){return b - a})

    hId = 0
    for (const cited of citations) {
        if (hId+1 > cited) {
            return hId
        }
        hId++
    }

    hId = citations.length

    return hId
}