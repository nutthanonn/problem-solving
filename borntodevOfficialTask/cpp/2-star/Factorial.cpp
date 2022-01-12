#include <bits/stdc++.h>
using namespace std;


int main () {
    int number;
    long factor = 1;
    cin >> number;
    for(int i = 1; i < number + 1 ; i ++) {
        factor *= i;
    }
    cout << factor;
}