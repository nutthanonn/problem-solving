#include <iostream>
using namespace std;

int main() {
   int a, b;
   cin >> a >> b;
   if(a<b){
       cout << "MAX : " << b << endl; 
       cout << "MIN : " << a << endl; 
   }else if(b<a){
       cout << "MAX : " << a << endl; 
       cout << "MIN : " << b << endl; 
   }
    return 0;
}