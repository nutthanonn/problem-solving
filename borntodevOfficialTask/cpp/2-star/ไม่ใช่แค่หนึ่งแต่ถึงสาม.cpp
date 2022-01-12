#include <iostream>
using namespace std;

int main(){
    int a = 0, b = 0;
    for(int i =0;i<3;i++){
        cin >> a;
        if(b<a){
            b = a;
        }
    }
    cout << "MAX : " << b;
    return 0;
}