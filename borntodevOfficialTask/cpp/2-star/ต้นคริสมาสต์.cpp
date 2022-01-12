#include <bits/stdc++.h>
using namespace std;

int main() {
    int round = 2;
    int number;
    cin >> number;
    for(int i = 0; i < number; i++){
        int space = number;
        for(int j = 0; j < round; j++){
            for(int k = 0; k < space; k++){
                cout << " ";
            }
            space--;
            for(int l=0; l < j*2+1; l ++){
                cout << "*";
            }
            cout << endl;
        }
        round++;
    }

    for(int i = 0; i < number; i++){
        cout << " ";
    }
    cout << "|" << endl;

    for(int i = 0; i < number; i++){
        cout << "=";
    }
    cout << "V";

    for(int i = 0; i < number; i++){
        cout << "=";
    }

}