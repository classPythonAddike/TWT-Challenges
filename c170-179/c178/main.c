#define MOD 1000000007

#include <stdio.h>
#include <string.h>

int letters[26];
long facts[501];
long mod_inv[501];

long MODInverse(long A, long M) {
    long y = 0, x = 1;
 
    while (A > 1) {
        long q = A / M;
        long t = M;
 
        M = A % M, A = t;
        t = y;
 
        y = x - q * y;
        x = t;
    }
    
    return x > 0 ? x : x + MOD;
}

int main() {
    setvbuf(stdout, NULL, _IOFBF, 1024 * 128);
    
    int T = 0;

    facts[0] = 1;
    mod_inv[0] = 1;
    for (int i = 1; i <= 500; i++) {
        facts[i] = (facts[i - 1] * i) % MOD;
        mod_inv[i] = MODInverse(facts[i], MOD) % MOD;
    }
    
    scanf("%d", &T);

    while (T--) {
        char word[1000];
        scanf("%s", word);
        memset(letters, 0, 26 * 4);

        for (int i = 0; i < strlen(word) / 2; i++) {
            letters[word[i] - 'a']++;
        }

        long perm = 1;
        int tot = 0;

        for (int r = 0; r < 26; r++) {
            perm = (perm * mod_inv[tot]) % MOD;
            tot += letters[r];
            perm = (perm * ((facts[tot] * mod_inv[letters[r]]) % MOD)) % MOD;
        }

        printf("%ld\n", perm);
    }
}
