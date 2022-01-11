#include <bits/stdc++.h>
using namespace std;

int main() {
    int arr[1000];
    int number = 0, index = 0;
    while(true) {
        cin >> number;
        if(number == 0){
            break;
        }
        arr[index] = number;
        index++;
    }
    int sortArr[index];

    string word;
    cin >> word;
    transform(word.begin(), word.end(), word.begin(), ::tolower);

    for(int i = 0; i < index; i++) {
        if(arr[i] == 0){
            break;
        }else{
            sortArr[i] = arr[i];
        }
    }

    if(word == "min"){
        sort(sortArr, sortArr + index);
    }else{
        sort(sortArr, sortArr + index, greater<int>());
    }


    for(int i = 0; i < index; i++) {
        cout << sortArr[i] << " ";
    }
    
    return 0;
}