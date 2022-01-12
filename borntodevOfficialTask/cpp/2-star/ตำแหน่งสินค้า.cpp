#include <bits/stdc++.h>
using namespace std;

int main()  {
    int all;
    cin >> all;
    int number[all];
    int findNumber;
    for(int i = 0; i < all; i++){
        cin >> number[i];
    }
    cin >> findNumber;

    string output = "";
    for(int i = 0; i < all; i++){
        if(number[i] == findNumber){
            output += to_string(i+1);
        }
    }
    cout << "Position: ";
    for(int i = 0; i < output.length(); i++){
        if(i == output.length() - 1){
            cout << output[i] << endl;
        }else{
            cout << output[i] << ",";
        }
    }

    return 0;
}