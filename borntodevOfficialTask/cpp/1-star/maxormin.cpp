#include <bits/stdc++.h>
using namespace std;

int main (){
    int a,b;
    cin >> a >> b;
    if(a>b){
        cout << "A";
    } else if(a==b){
        cout << "AB";
    }else{
        cout << "B";
    }
    return 0;
}