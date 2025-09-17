/**
 * @param {number} numRows
 * @return {number[][]}
 */
var generate = function(numRows) {
    let train = [[1]]

    for (let i = 1; i < numRows; i++){
        let newR = []
        for (let j = 0; j < i+1; j++){
            if (j === 0 || j === i){
                newR[j] = 1
            }else {
                newR[j] = train[i-1][j] + train[i-1][j-1]
            }
        }
        train[i] = newR
    }


    return train
};