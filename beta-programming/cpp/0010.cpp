#include <iostream>
using namespace std;

int main(){
    string w;
    cin >> w;
    int ball_position[3] = {1, 2, 3};
    //  find difference between lenght and sizeof
    for(int i=0;i<w.length();i++){
        if(w[i] == 'A') swap(ball_position[0], ball_position[1]);
        else if(w[i] == 'B') swap(ball_position[1], ball_position[2]);
        else if(w[i] == 'C') swap(ball_position[0], ball_position[2]);
    }
    for(int j=0;j<3;j++){
        if(ball_position[j] == 1) cout << j+1;
    }
}