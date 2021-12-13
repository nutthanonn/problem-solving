#include <iostream>
using namespace std;

int main(){
    int a, b, c;
    cin >> a;
    b = a-1;
    for(int i=0;i<a;i++){
        for(int j=0;j<b;j++){
            cout << " ";
        }
        b--;
        c = i*2+1;
        for(int k=1;k<c+1;k++){
            cout << "*";
        }
        cout << endl;
    }
    return 0;
}