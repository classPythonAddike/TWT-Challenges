#include <stdio.h>
#include <string.h>
#include <stdbool.h>

unsigned short int find_gcd(unsigned short int n, unsigned short int m) {
    unsigned short int r;
    while (n > 0) {
        r = m % n;
        m = n;
        n = r;
    }

    return m;
}

unsigned short int all(bool* L, unsigned short int N) {
    bool a = true;
    for (unsigned short int i = 0; i < N && a; i++) {
        a = a && L[i];
    }
    return a;
}

int main() {
    int T;
    scanf("%d", &T);

    unsigned short int n, m, n_happy, m_happy, gcd, dummy;
    bool step[100];

    while (T--) {
        scanf("%hu %hu", &n, &m);
        
        gcd = find_gcd(n, m);
        memset(step, false, gcd * sizeof(bool));
        
        scanf("%hu", &n_happy);
        while (n_happy--) {
            scanf("%hu", &dummy);
            step[dummy % gcd] = true;
        }

        scanf("%hu", &m_happy);
        while (m_happy--) {
            scanf("%hu", &dummy);
            step[dummy % gcd] = true;
        }

        if (all(step, gcd)) {
            puts("YES");
        } else {
            puts("NO");
        }
    }
}
