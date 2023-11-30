#include <stdio.h>
#include <stdbool.h>

const int N = 100000;
int palindromes[1099];

bool isPalindrome(int num)
{
    int rem = 0, rev = 0;
    
    for (int temp = num; temp > 0; temp = temp / 10)
    {
        rem = temp % 10;
        rev = 10 * rev + rem;
    }

    return rev == num;
}

int solve(int a, int b, int m)
{
    int count = 0, i = 0;

    for (int j = 0; j < 1099; j++)
    {
        i = palindromes[j];

        if (i < a)
            continue;
        else if (i > b)
            return count;
        else if (isPalindrome(i / m))
            count++;
    }

    return count;
}

void main()
{
    setvbuf(stdout, NULL, _IOFBF, 1024);
    
    int pointer = 0;
    for (int i = 0; i < N; i++)
    {
        if (isPalindrome(i))
        {
            palindromes[pointer] = i;
            pointer++;
        }
    }

    int T, a, b, m;
    
    scanf("%d", &T);

    for (int i = 0; i < T; i++)
    {
        scanf("%d", &a);
        scanf("%d", &b);
        scanf("%d", &m);

        printf("%d", solve(a, b, m));
        printf("\n");
    }
}
