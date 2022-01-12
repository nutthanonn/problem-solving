#include <iostream>
using namespace std;

int main() {
    int a, b, c;
    cin >> a;
    c = a - 1;
    for(int i=0;i<a;i++){
        for(int j=0;j<c;j++){
            cout << " ";
        }
        b = i*2+1;
        for(int k=0;k<b;k++){
            cout << "*";
        }
        c--;
        cout << endl;
    }

    return 0;
}