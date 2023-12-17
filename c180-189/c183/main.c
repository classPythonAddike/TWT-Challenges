#include <stdio.h>
#include <stdlib.h>

struct letter {
    char var;
    int depth;
    struct letter *parent;
};

struct letter* get_root(struct letter* node) {
    while (node->parent != node) {
        node = node->parent->parent;
    }

    return node;
}

int main() {
    setvbuf(stdout, NULL, _IOFBF, 1024 * 256);
    
    char *t = NULL, *a, *b;
    char *sub_a;
    char new_id, old_id;
    size_t len_a = 0, len_b = 0, len_t = 0;
    int s, *swaps;
    struct letter *a_let, *b_let;

    getline(&t, &len_t, stdin);
    int T = atoi(t);

    while (T--) {
        a = NULL;
        b = NULL;
        
        len_a = getline(&a, &len_t, stdin);
        len_b = getline(&b, &len_t, stdin);
        struct letter *letters = (struct letter*)calloc(6, sizeof(struct letter));

        for (int i = 0; i < 6; i++) {
            letters[i].var = i + 'a';
        }
        
        for (int i = 0; i <= len_a - len_b; i++) {
            sub_a = a + i;
            s = 0;

            for (int j = 0; j < 6; j++) {
                letters[j].parent = letters + j;
                letters[j].depth = 1;
            }
            
            for (int j = 0; j < len_b - 1; j++) {
                a_let = get_root(letters + (*(sub_a + j) - 'a'));
                b_let = get_root(letters + (*(b + j) - 'a'));
                if (a_let != b_let) {
                    s++;

                    if (s == 5) {
                        break;
                    }

                    if (a_let->depth > b_let->depth) {
                        b_let->parent = a_let;
                        a_let->depth += b_let->depth;
                    } else {
                        a_let->parent = b_let;
                        b_let->depth += a_let->depth;
                    }
                }
            }
            
            if (i != len_a - len_b) {
                printf("%d ", s);
            } else {
                printf("%d\n", s);
            }
        }
        
        free(a);
        free(b);
        free(letters);
    }
}
