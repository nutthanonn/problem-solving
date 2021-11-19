#include <iostream>
using namespace std;


int main(){
    int x = 10;
    int* p;  //Type pointer
    p = &x; // poiter to x
    cout << &x << endl;
    cout << *p << endl;
    cout << &p << endl;
}