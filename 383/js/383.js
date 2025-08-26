var canConstruct = function(ransomNote, magazine) {
    let m = ransomNote.length;
    let n = magazine.length;
    let magazineArray = magazine.split('');
    for (let i = 0; i < m; i ++){
        let find = false;
        for (let j = 0; j < n; j++){
            if (ransomNote[i] === magazineArray[j]){
                magazineArray[j] = "-";
                find = true;
                break;
            }
        }
        if (find == false){
            return false;
        }
    }

    return true;
};