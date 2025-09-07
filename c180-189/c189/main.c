#include <stdio.h>
#include <stdlib.h>

struct Group {
    int id;
    int parent;
};

int main() {
    int T, M, N, parent, child;
    scanf("%d", &T);

    struct Group* groups;

    for (int j = 0; j < T; j++) {
        scanf("%d %d", &N, &M);
        groups = calloc(N, sizeof(struct Group));

        for (int i = 0; i < N; i++) {
            groups[i].id = i;
            groups[i].parent = -1;
        }

        for (int i = 0; i < M; i++) {
            scanf("%d %d\n", &parent, &child);
            groups[child - 1].parent = parent - 1;
            printf("- %d %d\n", parent, child);
        }

        printf("Scanning\n");

        scanf("? %d %d", &parent, &child);

        for (int i = 0; i < N; i++) {
            printf("ID: %d Parent: %d\n", groups[i].id + 1, groups[i].parent + 1);
        }
        
        printf("Now: %d %d\n", parent, child);
        
        child -= 1;
        parent -= 1;
        
        while (child != parent && groups[child].parent != -1) {
            child = groups[child].parent;
            printf("Now: %d %d\n", parent + 1, child + 1);
        }

        printf(child == parent ? "Yes\n" : "No\n");
        free(groups);
    }
}
