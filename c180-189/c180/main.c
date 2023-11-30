#include <stdio.h>
#include <stdbool.h>
#include <string.h>

short max_dig(bool *digs) {
    short p = 9;
    while (!digs[p]) { p--; }
    return p;
}

short min_dig (bool *digs) {
    short q = 0;
    while (!digs[q]) { q++; }
    return q;
}

short max_diff(bool *digs) {
    return max_dig(digs) - min_dig(digs);
}

void digits(int a, bool *digs) {
    memset(digs, false, 10 * sizeof(bool));
    while (a > 0) {
        digs[a % 10] = true;
        a /= 10;
    }
}

int diff(int a, int b) {
    return a > b ? a - b : b - a;
}

bool digs_a[10], digs_b[10];

void brute_force(int a, int b) {
    short m = max_diff(digs_a), t;
    int m_i = a;
    for (int i = a + 1; i < b + 1 && m < 9; i++) {
        digits(i, digs_a);
        t = max_diff(digs_a);
        m_i = m < t ? i : m_i;
        m = m < t ? t : m;
    }
    printf("%d\n", m_i);
}

int main() {
    int T, a, b;
    short p, q, r, s;
    scanf("%d", &T);

    while (T--) {
        scanf("%d %d", &a, &b);
        
        digits(a, digs_a);
        digits(b, digs_b);

        if (b - a > 9 && a > 9) {
            if ((a / 100) == (b / 100)) {
                if (a / 10 == b / 10) {
                    printf("%d\n", max_diff(digs_a) >= max_diff(digs_b) ? a : b);
                } else if (a / 10 == b / 10 - 1){
                    p = max_diff(digs_a), q = max_diff(digs_b);
                    digits(10 * (b / 10), digs_a);
                    digits(10 * (b / 10) - 1, digs_b);
                    r = max_diff(digs_a), s = max_diff(digs_b);

                    if (p >= q && p >= r && p >= s)
                        printf("%d\n", a);
                    else if (s >= r && s >= q)
                        printf("%d\n", 10 * (b / 10) - 1);
                    else if (r >= q)
                        printf("%d\n", 10 * (b / 10));
                    else
                        printf("%d\n", b);
                } else {
                    brute_force(a, b);
                }
            } else {
                digits(a / 100, digs_a);
                digits(1 + a / 100, digs_b);

                if (digs_a[9] || digs_a[0]) {
                    digits(a, digs_a);
                    brute_force(a, b);
                } else {
                    if (a % 100 < 10)
                        printf("%d\n", 9 + 100 * (a / 100));
                    else if (a % 100 < 91)
                        printf("%d\n", 90 + 100 * (a / 100));
                    else if (b % 100 > 8 && !digs_b[9] && !digs_b[0])
                        printf("%d\n", 109 + 100 * (a / 100));
                    else if (b % 100 > 89 && !digs_b[9] && !digs_b[0])
                        printf("%d\n", 190 + 100 * (b / 100));
                    else {
                        digits(a, digs_a);
                        brute_force(a, b);
                    }
                }
            }
        } else {
            brute_force(a, b);
        }
    }
}
