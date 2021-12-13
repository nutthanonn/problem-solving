#include <bits/stdc++.h>
using namespace std;

int main() {
    string number;
    cin >> number;
    int result = 0;
    while (number.length() != 1) {
        for(int i = 0; i < number.length(); i++){
            int a = number[i];
            result += a - 48;
        }
        number = to_string(result);
        result = 0;
    }
    cout << number << endl;
}