#include <bits/stdc++.h>
using namespace std;

int main()  {
    int number;
    int round;
    string ans;
    cin >> number;
    cin >> round;
    for(int i = 0; i < round; i ++){
        int min_renge;
        int max_range;
        string key;
        cin >>  min_renge >> max_range >> key;
        if(number >= min_renge && number <= max_range){
            ans = key;
        }
        
    }
    cout << ans << endl;
   return 0;
}