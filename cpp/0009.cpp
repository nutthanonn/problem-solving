#include <iostream>
using namespace std;
int main() {
    int num[3] = {};
    string text;
    for(int i=0;i<3;i++) {
        cin >> num[i];
    }
    cin >> text;
    sort(num, num + 3);
    for(int j=0;j<3;j++){
        if(text[j] == 'A') cout << num[0] << " ";
        else if(text[j] == 'C')cout << num[2] << " ";
        else if(text[j] == 'B') cout << num[1] << " ";
    }
    return 0;
}