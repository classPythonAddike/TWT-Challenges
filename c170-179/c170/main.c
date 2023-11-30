#include <stdio.h>

int main() {
    setvbuf(stdout, NULL, _IOFBF, 1024 * 32);

    int T;
    long long A, B, sum;
    scanf("%d", &T);

    while (T--) {
        scanf("%lld %lld", &A, &B);
        sum = B - A - 1;
        if (sum * sum >= 4 * A && sum > 0) {
            printf("OK\n");
        } else {
            printf("LIE\n");
        }
    }
}
