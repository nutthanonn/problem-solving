#include <iostream>
using namespace std;

int main() {
    int n, x, a, b;
    scanf("%d", &n);
    for(int i=0;i<n;i++) {
        scanf("%d", &x);
        if(i==0) {
            a = x;
            b = x;
        }
        else{
            if(x>a) {
                a = x;
            }
            if(x<b) {
                b = x;
            }
        }

    }
    printf("%d", b);
    cout << endl;
    printf("%d", a);
}
