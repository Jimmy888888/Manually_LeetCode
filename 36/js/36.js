// Your solution code here
/**
 * @param {character[][]} board
 * @return {boolean}
 */

// can't repeat, check every column, row and 3 by 3 sub box

var isValidSudoku = function(board) {

    let n = board.length;
    // check every row
    for (let i = 0; i < n; i++){
        let checkMap = new Map();
        for (let j = 0; j < n; j++){
            const chr = board[i][j];
            if (chr === "."){
                continue;
            }
            if (checkMap.has(chr)){
                return false;
            }
            checkMap.set(chr, 1);
        }
    }

    // check every column
    for (let i = 0; i < n; i++){
        let checkMap = new Map();
        for (let j = 0; j < n; j++){
            const chr = board[j][i];
            if (chr === "."){
                continue;
            }
            if (checkMap.has(chr)){
                return false;
            }
            checkMap.set(chr, 1);
        }
    }

    // check every 3 by 3 sub box
    for (let x = 0; x < n; x += 3){
        for (let y = 0; y < n; y += 3){
            let checkMap = new Map();
            for (let i = x; i < x+3; i ++){
                for (let j = y; j < y+3; j++){
                    const chr = board[i][j];
                    if (chr === "."){
                        continue;
                    }
                    if (checkMap.has(chr)){
                        return false;
                    }
                    checkMap.set(chr, 1);
                }
            }
        }
    }
    
    return true;
};


// [
//     ["5","3",".",".","7",".",".",".","."],
//     ["6",".",".","1","9","5",".",".","."],
//     [".","9","8",".",".",".",".","6","."],
//     ["8",".",".",".","6",".",".",".","3"],
//     ["4",".",".","8",".","3",".",".","1"],
//     ["7",".",".",".","2",".",".",".","6"],
//     [".","6",".",".",".",".","2","8","."],
//     [".",".",".","4","1","9",".",".","5"],
//     [".",".",".",".","8",".",".","7","9"]
// ]