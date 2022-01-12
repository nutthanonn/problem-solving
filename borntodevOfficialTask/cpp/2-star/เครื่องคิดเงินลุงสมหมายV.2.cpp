#include <bits/stdc++.h>
using namespace std;


int main() {
    int round, money, allMoney;
    cin >> round;
    for(int i = 0; i < round; i++){
        cin >> money;
        allMoney += money;

    }
    cout << allMoney << " THB" << endl;
}