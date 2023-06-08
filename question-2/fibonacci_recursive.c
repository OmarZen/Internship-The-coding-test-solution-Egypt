#include <stdio.h>

int fibonacci_recursive(int n) {
    if (n == 0)
        return 0;
    else if (n == 1)
        return 1;
    else if (n == 2)
        return 2;
    else
        return fibonacci_recursive(n - 3) + fibonacci_recursive(n - 2);
}

int main() {
    int n = 10;
    int result = fibonacci_recursive(n);
    printf("F(%d) = %d\n", n, result);
    
    return 0;
}
