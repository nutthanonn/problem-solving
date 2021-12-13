#include <iostream>
using namespace std;

int main() {
    int a, b;
    cin >> a;
    if(a == 1){
        cout << "#";
    } else if(a == 2){
        cout << "##" << endl << "##";
    }else{
        for(int i=0;i<a;i++){
            cout << "#";
        }
        cout << endl;
        b = a - 2;
        for(int j=0;j<b;j++){
            cout << "#";
            for(int k=0;k<b;k++){
                cout << " ";
            }
            cout << "#" << endl;
        }
        for(int i=0;i<a;i++){
            cout << "#";
        }
    }
    return 0;
}
