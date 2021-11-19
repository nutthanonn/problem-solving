#include <iostream>
using namespace std;


void inc(int* y){
    (*y)++;
}


int main(){
    int x = 5;
    inc(&x);
    cout << x << endl;
    return 0;
}