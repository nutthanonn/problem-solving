#include <iostream>
using namespace std;

int main() {

    int arr[9] = {};
    int all = -100;
    for(int i=0;i<9;i++){
        cin >> arr[i];
    }
    for(int i=0;i<9;i++){
        all += arr[i];
    }
    for(int j=0;j<9;j++){
        for(int k=0;k<9;k++){
            if(arr[j] + arr[k] == all){
                for(int m=0;m<9;m++){
                    if(m==j || m==k){
                        continue;
                    }
                    cout << arr[m] << endl;
                }
                return 0;
            }
        }
    }
    return 0;
}