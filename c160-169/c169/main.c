#include <stdio.h>

int main() {
    setvbuf(stdout, NULL, _IOFBF, 1024 * 128);

    int T;
    long long int n;

    scanf("%d", &T);

    while (T--) {
        scanf("%lld", &n);
        
        if (n & 1 && n > 1) {
            printf("-1\n");
        } else {
            n += 2;
            if (n % 3) {
                printf("-1\n");
            } else {
                n /= 3;
                if ((n & (n - 1))) {
                    printf("-1\n");
                } else {
                    printf("%d\n", __builtin_ffsll(n));
                }
            }
        }
    }
}
