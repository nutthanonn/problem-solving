#include <bits/stdc++.h>
using namespace std;

int main (){
    int a = 4;
   for(int i=1;i<6;i++){
       for(int j=0;j<a;j++){
           cout << " ";
       }
       a--;
       for(int k=0;k<i;k++){
           cout << "$ ";
       }
       cout << endl;
   }
    return 0;
}