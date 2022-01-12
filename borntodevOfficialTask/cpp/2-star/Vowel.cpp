#include <bits/stdc++.h>
using namespace std;


int main() {
    string name;
    string out;
    getline(cin, name);
    for(int i = 0; i < name.length(); i ++){
        if(name[i] == 'a' || name[i] == 'A' || name[i] == 'e' || name[i] == 'E' || name[i] == 'I' || name[i] == 'i' || name[i] == 'o' || name[i] == 'O' || name[i] == 'u' || name[i] == 'U'){
            continue;
        }else{
            out += name[i];
        }
    };
    
    cout << out;

}