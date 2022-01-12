#include <bits/stdc++.h>
using namespace std;

int main(){
    int number;
    cin >> number;
    int space = number;
    for(int i = 0; i < number; i++){
        for(int j = 0; j < space-1; j ++){
            cout << " ";
        }
        for(int k = 0; k < i*2+1; k ++){
            cout << "*";
        }

        space--;
        cout << endl;
    }
    space = 1;
    for(int i = number - 1; i > 0; i--){
        for(int j = 0; j < space; j++){
            cout << " ";
        }
        for(int k = 0; k < i*2-1; k++) {
            cout << "*";
        }
        space++;
        cout << endl;
    }



    return 0;
}


// พีระมิด(ระดับ3)