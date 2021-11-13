#include <iostream>
#include <cmath>
using namespace std;

int main() {
    double a, b, c;
    cin >> a >> b;
    c = hypot(a,b);
    printf("%f", c);
    return 0;
}