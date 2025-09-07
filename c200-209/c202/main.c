#include <stdio.h>

int main() {
    int T;
    scanf("%d", &T);

    while (T--) {
        long long n;
        scanf("%lld", &n);

        // Index of Least Significant Set Bit
        int lindex = __builtin_ffsll(n) - 1;
        int rindex;
        long long n_sib;

        if (lindex) {
            rindex = __builtin_ctzll(~(n >> lindex)) - 1 + lindex;
            n_sib = (n | (1LL << (rindex + 1))) & (~(1LL << rindex));

            /* XXXXX11111100
             *     r     l
             * n = 2, i = r - n, j = 0
             * XXXXX11000000
             *     r l
             * n = 2, i = r - n, j = 0
             */
            int l = lindex > rindex - lindex ? rindex - lindex : lindex; // number of bits to swap
            int i = rindex - l; // positions of bits to swap
            long long x = ((n_sib >> i) ^ n_sib) & ((1LL << l) - 1);
            n_sib ^= (x << i) | x;
        } else {
            rindex = __builtin_ctzll(~n);
            n_sib = (n | (1LL << rindex)) & (~(1LL << (rindex - 1)));
        }

        n_sib = n ? n_sib : -1;

        printf("%lld\n", n_sib);
    }
}
