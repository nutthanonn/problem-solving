#include <bits/stdc++.h>
using namespace std;


int main()  {
    string name;
    cin >> name;
    int space = name.length() * 2 - 2;
    int re = 0;
    for(int i = 0; i < name.length(); i++) {
        //space
        for(int j = 0; j < space; j++){
            cout << " ";
        }
        if(i == 0){
            cout << name[i];
        }else{
        //reversse
            for(int j = re; j > 0; j--){
                cout << name[j] << " ";
            }

            for(int j = 0 ; j < i+1; j ++){
                cout << name[j] << " ";
            }

        }

       
        re++;
        space -= 2;
        cout << endl;
    }
}