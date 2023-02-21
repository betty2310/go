#include <iostream>

int fibonacci(int x) {
    if (x == 0 || x == 1) return x;
    return fibonacci(x - 1) + fibonacci(x - 2);
}

int main() { std::cout << fibonacci(10); }
