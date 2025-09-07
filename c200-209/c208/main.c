#include <stdio.h>
#include <stdlib.h>

long int read_long_int() {
    long int c = getchar_unlocked();
    long int x = 0;

    while (c < 48 || c > 57)
        c = getchar_unlocked();

    for(; c > 47 && c < 58 ; c = getchar_unlocked())
        x = (x << 1) + (x << 3) + c - 48;

    return x;
}

int main() {   
    long int T, n, m;
    long int *arr;
    
    scanf("%ld", &T);
    arr = (long int*)malloc(100001 * sizeof(long int));
    arr[0] = 0;

    while (T--) {
        n = read_long_int();
        m = read_long_int();

        for (int i = 0; i < n; i++) {
            arr[i + 1] = arr[i] ^ read_long_int();
        }

        for (int i = 1; i < m; i++) {
            printf("%ld ", arr[read_long_int() - 1] ^ arr[read_long_int()]);
        }
    
        printf("%ld\n", arr[read_long_int() - 1] ^ arr[read_long_int()]);
    }

    free(arr);
}
