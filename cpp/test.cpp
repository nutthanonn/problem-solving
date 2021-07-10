#include <iostream>
using namespace std;

int main(){
    string a;
    getline(cin, a);
    for(int i=0;i<a.length();i++){
        if(a[i] == 'A' || a[i] == 'a' ||  a[i] == 'E' || a[i] == 'e' || a[i] == 'I' || a[i] == 'i' || a[i] == 'O' || a[i] == 'o' || a[i] == 'U' || a[i] == 'u'){
            a[i] = '\n';
        }
    }
    cout << a;
    return 0;
}