#include <bits/stdc++.h>
using namespace std;


int main() {
    int a[5];
    for(int i = 0; i < 5; i++){
        cin >> a[i];
    }        
    int size = sizeof(a) / sizeof(a[0]);
    sort(a, a + size, greater<int>());
    for(int i = 0; i < size; i ++ ){
        cout << a[i] << endl;
    }
}