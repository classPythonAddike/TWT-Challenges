#include <stdio.h>
#include <stdlib.h>

struct node {
    struct node *next;
    int x;
    int y;
};

typedef struct node node;

int mod(int x) {
    return x > 0 ? x : -x;
}

int distance(node *start, int goal_x, int goal_y) {
    return (start->x - goal_x) * (start->x - goal_x) + (start->y - goal_y) * (start->y - goal_y);
}

const int moves[8][2] = {
    {-2, -1},
    {-2, 1},
    {-1, -2},
    {-1, 2},
    {1, -2},
    {1, 2},
    {2, -1},
    {2, 1}
};

int make_move(node **start, node **seen_tail, int n, int goal_x, int goal_y, int *m_dist, int *visted) {
    int n_x, n_y, r_x, r_y, d_x, d_y;
    node *move;
    
    for (int i = 0; i < 8; i++) {
        n_x = moves[i][0] + (*start)->x;
        n_y = moves[i][1] + (*start)->y;

        if (n_x < 0 || n_y < 0 || n_x >= n || n_y >= n) {
            continue;
        }

        r_x = mod(n_x - goal_x);
        r_y = mod(n_y - goal_y);
        
        if (r_x * r_x + r_y * r_y > 200) {
            int n_dist = r_x * r_x + r_y * r_y;
            if (n_dist > *m_dist) {
                continue;
            } else {
                *m_dist = n_dist;
            }
        }
        
        if (visted[n_x * n + n_y]) {
            continue;
        }
        
        visted[n_x * n + n_y] = visted[(*start)->x * n + (*start)->y] + 1;
        
        move = (node*)calloc(1, sizeof(node));
        move->x = n_x;
        move->y = n_y;
        
        (*seen_tail)->next = move;
        *seen_tail = move;
        
        if (goal_x == n_x && goal_y == n_y) {
            return 1;
        }
    }
    
    *start = (*start)->next;
    
    return 0;
}

int solve(int n, int x0, int y0, int x1, int y1) {
    if (x0 == x1 && y0 == y1) {
        return 0;
    }

    int *visited = (int*)calloc(n * n, sizeof(int));

    node *head = (node*)calloc(1, sizeof(node)), *dummy, *pointer = head, *tail = head;
    head->x = x0;
    head->y = y0;

    int m_dist = distance(head, x1, y1);
    
    while (!make_move(&pointer, &tail, n, x1, y1, &m_dist, visited)) {        
        dummy = head->next;
        free(head);
        head = dummy;
    }

    m_dist = visited[x1 * n + y1];
    
    while (head) {
        dummy = head->next;
        free(head);
        head = dummy;
    }
    
    free(visited);
    return m_dist;
}

int main() {
    int T;
    scanf("%d", &T);

    int n, x0, y0, x1, y1;
    while (T--) {
        scanf("%d %d %d %d %d", &n, &x0, &y0, &x1, &y1);
        printf("%d\n", solve(n, x0, y0, x1, y1));
    }
}
