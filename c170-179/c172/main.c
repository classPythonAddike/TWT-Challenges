#include <stdio.h>
#include <stdlib.h>

#define MIN(x, y) (((x) < (y)) ? (x) : (y))

struct video {
    int time;
    short skillA;
    short skillB;
};

int main() {
    setvbuf(stdout, NULL, _IOFBF, 1024 * 32);

    int T, N;
    long int skill_A_time, skill_B_time, skill_A_B_time;
    struct video vid;
    scanf("%d", &T);

    while (T--) {
        scanf("%d", &N);
        skill_A_time = -1;
        skill_B_time = -1;
        skill_A_B_time = -1;

        while (N--) {
            scanf("%d %1hd%1hd", &(vid.time), &(vid.skillA), &(vid.skillB));
            
            if (vid.skillA && vid.skillB) {
                if (vid.time < skill_A_B_time || skill_A_B_time == -1) {
                    skill_A_B_time = vid.time;
                }
            } else if (vid.skillA) {
                if (vid.time < skill_A_time || skill_A_time == -1) {
                    skill_A_time = vid.time;
                }
            } else if (vid.skillB) {
                if (vid.time < skill_B_time || skill_B_time == -1) {
                    skill_B_time = vid.time;
                }
            }
        }

        if (skill_A_time > 0 && skill_B_time > 0 && skill_A_B_time > 0)
            printf("%ld\n", MIN(skill_A_time + skill_B_time, skill_A_B_time));
        else if (skill_A_B_time > 0)
            printf("%ld\n", skill_A_B_time);
        else if (skill_A_time > 0 && skill_B_time > 0)
            printf("%ld\n", skill_A_time + skill_B_time);
        else
            printf("-1\n");
    }
}
