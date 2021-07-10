#include <iostream>
using namespace std;

int main() {
    int a, b,c;
    cin >> a >> b;
    c = a;
for(int i=1;i<b;i++){
        a *= c;
    }
    cout << a;
    return 0;
}