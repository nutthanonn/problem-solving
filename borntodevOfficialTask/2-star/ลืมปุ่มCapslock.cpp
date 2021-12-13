#include <bits/stdc++.h>
using namespace std;

int main()  {
    string word;
    getline(cin, word);
    string output;
    for(int i = 0; i < word.length(); i++) {
        if(int(word[i]) <= 90 && int(word[i]) >= 65 ){
            output += tolower(word[i]);
        }else if(int(word[i]) <= 122 && int(word[i]) >= 97 ){
            output += toupper(word[i]);
        }else{
            output += word[i];
        }
    }
    cout << output << endl;
    return 0;
}