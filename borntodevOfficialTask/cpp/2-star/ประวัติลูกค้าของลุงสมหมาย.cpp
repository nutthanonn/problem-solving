#include <iostream>
using namespace std;

void personData(){
   string fname, lname, address, gender, tel;
   getline(cin, fname);
   getline(cin, lname);
   getline(cin, address);
   getline(cin, gender);
   getline(cin, tel);
   
   cout << "--- Customer Detail ---" << endl;
   cout << "Name : " << fname << " " << lname << endl;
   cout << "Address : " << address << endl;
   cout << "Gender : " << gender << endl;
   cout << "Tel : " << tel << endl;
};

int main(){
   personData();   
}
