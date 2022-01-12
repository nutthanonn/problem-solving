#include <bits/stdc++.h>
using namespace std;

int main() {
    int ro;
    string number, str;
    cin >> ro;
    for(int i = 0; i < ro; i ++ ){
        cin  >> number;
        str += number;
    }
    reverse(str.begin(), str.end());
    for(int i = 0; i < str.size(); i ++) {
        cout << str[i] << endl;
    }
    
}