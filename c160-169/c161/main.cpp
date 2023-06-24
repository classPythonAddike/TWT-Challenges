#include <iostream>

using namespace std;

int main(int argc, char const *argv[])
{
	int T, a, b, i;
	cin >> T;

	for (i = 0; i < T; i++) {
		cin >> a;
		cin >> b;

		if (a > b) {
			cout << 2 * b - 1 << endl;
		} else {
			cout << 2 * a - 1 << endl;
		}
	}

	return 0;
}