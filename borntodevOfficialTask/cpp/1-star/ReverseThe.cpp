#include <bits/stdc++.h>
using namespace std;
 
void split(string s) {
    string arr[100];
    int index = 0;
    stringstream ss(s);
    string word;

    while (ss >> word) {
        arr[index] = word;
        index++;
    }
    for(int i = index-1 ; i >= 0 ; i--) {
        cout << arr[i] << " ";
    }

}
 
int main() {
    string a;
    getline(cin, a);
    split(a);
    return 0;
}