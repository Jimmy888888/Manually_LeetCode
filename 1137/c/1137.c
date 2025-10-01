int tribonacci(int n) {

    switch (n){
        case 0:
            return 0;
        case 1:
            return 1;
        case 2:
            return 1;
        default:
            break;
    }

    int t1 = 0, t2 = 1, t3 = 1;
    int curT = 0;

    for (int i = 3; i <= n; i++){
        curT = t1 + t2 + t3;
        t1 = t2;
        t2 = t3;
        t3 = curT;
    }

    return curT;
}