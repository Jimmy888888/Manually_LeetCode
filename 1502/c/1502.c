# include <stdlib.h>

int compare_int(const void* a, const void* b) {
    int a_val = *(const int*) a;
    int b_val = *(const int*) b;

    return a_val - b_val;
}

bool canMakeArithmeticProgression(int* arr, int arrSize) {
    qsort(arr, arrSize, sizeof(int), compare_int);

    int commDiff = arr[0] - arr[1];
    for (int i = 1; i < arrSize-1; i++) {
        if (arr[i] - arr[i+1] != commDiff){
            return false;
        }
    }

    return true;
}