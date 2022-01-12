#include <bits/stdc++.h>
using namespace std;

int main () {
    int round;
    string name1;
    cin >> round;
    cin.ignore();
    string name[round], gender[round];
    int age[round];
    for(int i = 0; i < round; i++){
        getline(cin, name[i]);
        cin >> age[i];
        cin >> gender[i];
        cin.ignore();
    }

    cout << "--Customers Information--" << endl;

    for(int i = 0; i < round; i++){
        cout << name[i] << " (age : " << 2017 - age[i] << ")" << endl;
    }

}