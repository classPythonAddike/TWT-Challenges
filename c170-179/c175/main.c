#include <stdio.h>

static inline int find_set_bits(long long int n, int* arr) {
    int num_set_bits = __builtin_popcountll(n);
    for (int num = 0; num < num_set_bits; num++) {
        *(arr + num + 1) = __builtin_ffsll(n >> *(arr + num)) + *(arr + num);
    }
    return num_set_bits;
}

int main() {
    setvbuf(stdout, NULL, _IOFBF, 1024 * 1024);
    int T, set_bits[21] = {0}, num_set_bits;
    scanf("%d", &T);

    long long int m, n, reconstruct, num_multiples;
    while (T--) {
        scanf("%lld %lld", &m, &n);

        num_set_bits = find_set_bits(m, set_bits);
        num_multiples = 0;
        for (int i = 0; i < 1 << num_set_bits; i++) {
            reconstruct = 0;
            for (int j = 0; j < num_set_bits; j++) {
                reconstruct += ((((1 & i >> j) == 1) * 1LL << set_bits[j + 1]) >> 1);
            }
            num_multiples += reconstruct % n == 0;
        }

        printf("%lld\n", num_multiples);
    }
}
