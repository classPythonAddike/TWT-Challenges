#include <stdio.h>

struct game {
    unsigned short o : 2;
    unsigned short e : 1;
};

int read() {
    int n = 0;
    char new = getchar_unlocked();
    while (new >= '0' && new <= '9') {
        n = (n << 3) + (n << 1) + new - '0';
        new = getchar_unlocked();

        if (new < '0' || new > '9')
            return n;

        n = (n << 3) + (n << 1) + new - '0';
        new = getchar_unlocked();
    }
    return n;
}

int main() {
    int T, n, i, reached;
    struct game g;
    scanf("%d\n", &T);

    char line[10000], prev;

    while (T--) {
        n = read();
        reached = n;
        g.o = 0;

        while (n) {
            fgets(line, 10000, stdin);
            for (i = 0; line[i] != 0; i++) {
                if (line[i] < 44) {
                    g.o += (i ? line[i - 1] : prev) & 1;
                    n--;
                }
            }
            prev = line[i - 1];
        }

        g.e = reached - g.o;
        puts(185 >> (4 * g.e + g.o) & 1 ? "Yan" : "Tim");
    }
}
