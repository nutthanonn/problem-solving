#include <bits/stdc++.h>
using namespace std;

int main() {
    int a;
    cin >> a;
    if(a >= 80  &&  a<= 100){
        printf("A");
    } else if(a >= 70 && a <= 79){
        printf("B");
    } else if(a >= 60 && a <= 69){
        printf("C");
    }else if(a >= 50 && a<= 59){
        printf("D");
    }else if( a < 50){
        printf("F");
    }
    return 0;
}