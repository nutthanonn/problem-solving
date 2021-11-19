#include <iostream>
using namespace std;





int main(){
    string name = "nutthanon";
    string* p = &name;
    
    cout << *p << endl;
}