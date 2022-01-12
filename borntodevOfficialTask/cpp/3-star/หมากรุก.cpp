#include <bits/stdc++.h>
using namespace std;


int main()  {
    string chess;
    map<char, int> score;
    score['p'] = 1;
    score['n'] = 3;
    score['b'] = 3;
    score['r'] = 5;
    score['q'] = 9;
    score['k'] = 0;
    int score_me = 0;
    int score_you = 0;
    for(int i = 0; i < 8; i ++){
        cin >> chess;
        for(int j = 0; j < chess.length(); j++){
            if(int(chess[j]) <= 90 && int(chess[j]) >= 65){
                score_you += score[tolower(chess[j])];    
            }else if(int(chess[j]) <= 122 && int(chess[j]) >= 97){
                score_me += score[chess[j]];
            }
        }
    }
    if(score_you - score_me == 0){
        cout << "equal";
    }else{
        cout << score_you - score_me << endl;
    }


    return 0;
}