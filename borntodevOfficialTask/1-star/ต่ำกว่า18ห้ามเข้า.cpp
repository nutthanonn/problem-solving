#include <iostream>
using namespace std;

int main() {
    string s;
    int a;
    getline(cin, s);
    cin >> a;
    if(s == "Sommai MeeMakMai"){
        cout << "Welcome Sommai MeeMakMai to NongGipsy Pub";
    } else{
        if(a < 2005){
            cout << "Welcome " << s << " to NongGipsy Pub";
        } else {
            cout << "You shall not pass!";
        }
    }
    return 0;
}