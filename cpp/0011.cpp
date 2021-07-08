#include <iostream>
#include <unordered_set>
using namespace std;

int main() {
    int num[10] = {};
    int len;
    for(int i=0;i<10;i++){
        cin >> num[i];
    }
    for(int j=0;j<10;j++){
        num[j] = num[j]%42;
    }
    const size_t len = sizeof(num) / sizeof(num[0]);
    unordered_set<int> s(num, num + len);
    cout << s.size() << endl;
    return EXIT_SUCCESS;

}