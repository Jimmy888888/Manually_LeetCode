# 389

```c
char findTheDifference(char* s, char* t) {
    int sASCII = 0;
    int tASCII = 0;

    char* p = s;

    while ( *p != '\0'){
        sASCII += *p;
        p++;
    }

    p = t;

    while ( *p != '\0'){
        tASCII += *p;
        p++;
    }

    char result = tASCII - sASCII;

    return result;
}
```


## solution 2 :

using XOR :

 A ^ A = 0

 A ^ 0 = A

 A ^ A ^ B = B
```c
char findTheDifference(char* s, char* t) {
    char result = 0;
    while (*s) {result ^= *s; s++;}
    while (*t) {result ^= *t; t++;}
    return result;
}
```