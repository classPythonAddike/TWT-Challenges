#include <stdio.h>
#include <stdlib.h>

typedef long long int llint ;

struct demon {
    llint defense;
    llint attack;
};

void read_heroes(int, llint*);
void read_demons(int, struct demon*);
int compare(const void*, const void*);
void binary_search(llint, int, llint*, llint*, int);
llint sum(int, llint*);
llint min(llint, llint);

llint *less_hero, *equal_hero, *more_hero;

int main() {
    int T;
    scanf("%d", &T);

    int n_heroes, n_demons;

    while (T--) {
        scanf("%d %d", &n_heroes, &n_demons);
        
        llint heroes[n_heroes];
        struct demon demons[n_demons];

        read_heroes(n_heroes, heroes);
        read_demons(n_demons, demons);

        llint coins = 0, hero_strength = sum(n_heroes, heroes);
        
        for (llint i = 0; i < n_demons; i++) {
            struct demon current_demon = demons[i];
            binary_search(demons[i].defense, n_heroes, heroes, heroes + n_heroes - 1, 1);

            if (equal_hero) {
                if (hero_strength - *equal_hero < current_demon.attack)
                    coins += current_demon.attack - hero_strength + *equal_hero;
            } else {
                llint l_coins = -1, m_coins = -1;
                
                if (less_hero) {
                    if (hero_strength - *less_hero < current_demon.attack)
                        l_coins = current_demon.attack - hero_strength + current_demon.defense;
                    else
                        l_coins = current_demon.defense - *less_hero;
                }

                if (more_hero) {
                    if (hero_strength - *more_hero < current_demon.attack)
                        m_coins = current_demon.attack - hero_strength + *more_hero;
                    else
                        m_coins = 0;
                }
                
                if (l_coins >= 0 && m_coins >= 0) 
                    coins += min(l_coins, m_coins);
                else if (l_coins >= 0)
                    coins += l_coins;
                else
                    coins += m_coins;
            }

        }

        printf("%lld\n", coins);
    }
}

void read_heroes(int n_heroes, llint* heroes) {
    for (int i = 0; i < n_heroes; i++)
        scanf("%lld", heroes + i);

    qsort((void*)heroes, n_heroes, sizeof(llint), compare);
}

void read_demons(int n_demons, struct demon* demons) {
    for (int i = 0; i < n_demons; i++)
        scanf("%lld %lld", &(demons[i].defense), &(demons[i].attack));
}

int compare(const void* a, const void* b) {
    return *(llint*)a - *(llint*)b;
}

void binary_search(llint key, int n_heroes, llint* heroes_start, llint* heroes_end, int first_call) {
    if (key > *heroes_end) {
        less_hero = heroes_end;
        equal_hero = NULL;

        if (first_call)
            more_hero = NULL;
        else
            more_hero = heroes_end + 1;
        
        return;
    }

    if (key < *heroes_start) {
        if (first_call)
            less_hero = NULL;
        else
            less_hero = heroes_start - 1;

        equal_hero = NULL;
        more_hero = heroes_start;
        return;
    }

    llint mid_ind = (n_heroes - 1) / 2;
    llint mid = *(heroes_start + mid_ind);

    if (mid > key)
        binary_search(key, mid_ind, heroes_start, heroes_start + mid_ind - 1, 0);
    else if (mid < key)
        binary_search(key, mid_ind, heroes_start + mid_ind + 1, heroes_end, 0);
    else {
        less_hero = heroes_start + mid_ind - 1;
        equal_hero = heroes_start + mid_ind;
        more_hero = heroes_start + mid_ind + 1;
    }
}

llint sum(int n_heroes, llint* heroes) {
    llint s = 0;
    for (llint i = 0; i < n_heroes; i++)
        s += heroes[i];
    return s;
}

llint min(llint a, llint b) {
    if (a > b)
        return b;
    return a;
}
