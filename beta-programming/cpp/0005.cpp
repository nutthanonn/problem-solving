#include <iostream>
using namespace std;

int main() {
    int lo=0, to=0;
    string a;
    cin >> a;
    for(int i=0;i<a.size();i++){
        if(int(a[i]) >= int('A') && int(a[i]) <= int('Z')){
            to++;
        }
        else if(int(a[i]) >= int('a') && int(a[i]) <= int('z')) {
            lo++;
        }
    }
    if(lo==0 && to!=0){
        cout << "All Capital Letter";
    }
    else if(lo!=0 && to==0){
        cout << "All Small Letter";
    }
    else{
        cout << "Mix";
    }
    return 0;
}