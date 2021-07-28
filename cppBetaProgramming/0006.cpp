#include <iostream>
#include <cmath>
using namespace std;

int main() {
    float a, b, c;
    cin >> a >> b;
    c = hypot(a, b);
    cout << c;
    return 0;
}