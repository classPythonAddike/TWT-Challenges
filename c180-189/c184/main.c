#include <stdio.h>
#include <stdlib.h>

#define MOD 1000000007LL

long long *trib_s;

void nth_trib(unsigned int n) {
	trib_s[0] = 0;
	trib_s[1] = 1;
	trib_s[2] = 2;
	trib_s[3] = 10;
	
	long long a = 1, b = 1, c = 2, d = 2;

	for (int i = 4; i <= n; i++) {
		d = (a + b + c) % MOD;
		a = b;
		b = c;
		c = d;

		trib_s[i] = (trib_s[i - 1] + (((d * d) % MOD) * d) % MOD) % MOD;
	}
}

int main() {
	int T;
	unsigned int n;
	
	char *buf = NULL;
	size_t t;

	getline(&buf, &t, stdin);
	T = atoi(buf);


	trib_s = (long long*)malloc(1000001 * sizeof(long long));
	nth_trib(1000000);

	while (T--) {
		getline(&buf, &t, stdin);
		printf("%lld\n", trib_s[atoi(buf)]);
	}

	free(trib_s);
}
