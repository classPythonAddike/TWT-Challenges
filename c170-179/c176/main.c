#include <math.h>
#include <stdio.h>

double a, b, c;

double f(double x) {
    return (a * x * x + b * x + c) / sin(x);
}

double f_dash(double x) {
    double sin_x = sin(x);
    return (sin_x * (2 * a * x + b) - cos(x) * (a * x * x + b * x + c)) / pow(sin_x, 2);
}

int product_is_positive(double a, double b) {
    int a_pos = a > 0, b_pos = b > 0;
    return (a_pos && b_pos) || (!a_pos && !b_pos);
}

double approx_bisection(double x_lower, double x_upper, double f_x_lower, double f_x_upper, double tolerance) {
    double x_guess = (x_lower + x_upper) / 2;

    if (x_upper - x_guess < tolerance) {
        return x_guess;
    }

    double f_x_guess = f_dash(x_guess);
    
    if (f_x_guess == 0) {
        return x_guess;
    } else if (product_is_positive(f_x_upper, f_x_guess)) {
        return approx_bisection(x_lower, x_guess, f_x_lower, f_x_guess, tolerance);
    } else {
        return approx_bisection(x_guess, x_upper, f_x_guess, f_x_upper, tolerance);
    }
}

int main() {
    setvbuf(stdout, NULL, _IOFBF, 1024 * 256);
    int T;

    scanf("%d", &T);
    double x_lower = 0.01, x_upper = 1.55;

    while (T--) {
        scanf("%lf %lf %lf", &a, &b, &c);
        printf("%.5lf\n", f(approx_bisection(x_lower, x_upper, f_dash(x_lower), f_dash(x_upper), 0.000001)));
    }
}
