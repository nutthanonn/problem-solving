#include <iostream>
using namespace std;


// ผู้ที่ได้คะแนนตั้งแต่ 90 คะแนนขึ้นไปจะได้เกรด A
// ผู้ที่ได้คะแนนตั้งแต่ 85 - 89 จะได้เกรด B+
// ผู้ที่ได้คะแนนตั้งแต่ 80 - 84 จะได้เกรด B
// ผู้ที่ได้คะแนนตั้งแต่ 75 - 79 จะได้เกรด C+
// ผู้ที่ได้คะแนนตั้งแต่ 70 - 74 จะได้เกรด C
// ผู้ที่ได้คะแนนตั้งแต่ 65 - 69 จะได้เกรด D+
// ผู้ที่ได้คะแนนตั้งแต่ 60 - 64 จะได้เกรด D
// ผู้ที่ได้คะแนนต่ำกว่า 60 ลงไปจะได้เกรด F



int main(){
   double score;
   cin >> score;

   



   if(score >= 90 && score <= 100){
      cout << "A";
   }else if(score >= 85 && score <= 89){
      cout << "B+";
   }else if(score >= 80 && score <= 84){
      cout << "B";
   }else if(score >= 75 && score <= 79){
      cout << "C+";
   }else if(score >= 70 && score <= 74){
      cout << "C";
   }else if(score >= 65 && score <= 69){
      cout << "D+";
   }else if(score >= 60 && score <= 64){
      cout << "D";
   } else if(score <= 50 && score >= 0 ){
      cout << "F";
   }else if(score > 100){
      cout << "Error : Value must be less than or equal to 100.";
   }else{
      cout << "Error : Value must be greater than or equal to 0.";
   }


}