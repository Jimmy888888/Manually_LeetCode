/**
 * @param {number} rowIndex
 * @return {number[]}
 */
var getRow = function(rowIndex) {
    let result = [[1]];

    for (let i = 1; i <= rowIndex; i++ ){
        let newR = [];
        for (let j = 0; j < i+1; j++){
            if (j === 0 || j === i){
                newR[j] = 1;
            }else{
                newR[j] = result[i-1][j] + result[i-1][j-1];
            }
        }
        result[i] = newR;
    }
    return result[rowIndex];
};