#include <bits/stdc++.h>
using namespace std;

int main () {
    int round;
    cin  >> round;
    int number[round];
    for(int i = 0; i < round; i++) {
        cin >> number[i]; 
    }
    int size = round;

    sort(number, number + size, greater<int>());
    for(int i = 0; i < round; i++) {
        cout << number[i] << endl;
    }
}